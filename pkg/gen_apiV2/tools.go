package gen_apiV2

import (
	"bytes"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"io/fs"
	"log"
	"path"
	"regexp"
	"strings"
)

func DirByImport(it string) string {
	p, err := build.Import(strings.ReplaceAll(it, `"`, ""), "/", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
}

const mode packages.LoadMode = packages.NeedName |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedImports |
	packages.NeedModule |
	packages.NeedExportsFile |
	packages.NeedTypesSizes |
	packages.NeedDeps

func LoadPkgV2(dir string) (string, *packages.Package, *token.FileSet) {
	pkgName := path.Base(dir)
	var fset = token.NewFileSet()
	cfg := &packages.Config{Fset: fset, Mode: mode}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		log.Panicln(err)
	}
	for _, pkg := range pkgs {
		if path.Base(pkg.String()) == pkgName {
			return pkgName, pkg, fset
		}
	}
	log.Panicln("not found pkg")
	return "", nil, nil
}

func LoadPkg(dir string) (string, *ast.Package, *token.FileSet) {
	pkgName := path.Base(dir)
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(info fs.FileInfo) bool {
		return true
	}, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	pkg, ok := pkgs[pkgName]
	if !ok {
		log.Fatalln("not found pkg name: " + pkgName)
	}
	return pkgName, pkg, fset
}

func FindStructByPkg(files []*ast.File, structName string) (*ast.File, *ast.StructType, bool) {
	var structType *ast.StructType
	var f *ast.File
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			if typeSpec, ok := node.(*ast.TypeSpec); ok {
				if typeSpec.Name.Name == structName {
					if st, ok := typeSpec.Type.(*ast.StructType); ok {
						structType = st
						f = file
						return false
					}
				}
			}
			if structType != nil {
				return false
			}
			return true
		})
	}
	if structType == nil {
		return f, structType, false
	}
	return f, structType, true
}

func FindStructByDir(dir string, structName string) (*packages.Package, *token.FileSet, *ast.File, *ast.StructType, bool) {
	_, pkg, fset := LoadPkgV2(dir)
	f, structType, has := FindStructByPkg(pkg.Syntax, structName)
	return pkg, fset, f, structType, has
}

func FindStructByImport(it string, structName string) (*packages.Package, *token.FileSet, *ast.File, *ast.StructType, bool) {
	dir := DirByImport(it)
	return FindStructByDir(dir, structName)
}

func IsSelector(field *ast.Field) (string, string, bool) {
	var has bool
	var x string
	var sel string
	ast.Inspect(field.Type, func(node ast.Node) bool {
		if selectorExpr, ok := node.(*ast.SelectorExpr); ok {
			has = true
			x = selectorExpr.X.(*ast.Ident).Name
			sel = selectorExpr.Sel.Name
			return false
		}
		return true
	})
	if !has {
		ast.Inspect(field.Type, func(node ast.Node) bool {
			ident, ok := node.(*ast.Ident)
			if ok {
				sel = ident.Name
				return false
			}
			return true
		})
	}
	return x, sel, has
}

func Node2String(fset *token.FileSet, node interface{}) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, node)
	if err != nil {
		spew.Dump(node)
		log.Fatalln(err.Error())
	}
	return buf.String()
}

//type Req struct {
//转换部分 xxx.Client
//Name Client
//}
func Node2SwagType(node ast.Node, selectName string) ast.Node {
	t := node2SwagType2(node, selectName)
	t = node2SwagType1(t)
	return t
}

// 去掉指针
func node2SwagType1(node ast.Node) ast.Node {
	return astutil.Apply(node, func(c *astutil.Cursor) bool {
		switch c.Node().(type) {
		case *ast.StarExpr:
			tmp := c.Node().(*ast.StarExpr).X
			c.Replace(tmp)
		}
		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
}

func node2SwagType2(node ast.Node, selectName string) ast.Node {
	return astutil.Apply(node, func(c *astutil.Cursor) bool {
		switch c.Node().(type) {
		case *ast.SelectorExpr:
			return false

		case *ast.Ident:
			spew.Dump(c.Node())
			if ok := JudgeBuiltInType(c.Node().(*ast.Ident).Name); !ok {
				tmp := ast.SelectorExpr{X: ast.NewIdent(selectName), Sel: ast.NewIdent(c.Node().(*ast.Ident).Name)}
				c.Replace(&tmp)
			}
		}
		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
}

func JudgeBuiltInType(t string) bool {
	m := map[string]int{
		"uint8":      0,
		"uint16":     0,
		"uint32":     0,
		"uint64":     0,
		"int8":       0,
		"int16":      0,
		"int32":      0,
		"int64":      0,
		"float32":    0,
		"float64":    0,
		"complex64":  0,
		"complex128": 0,
		"byte":       0,
		"rune":       0,
		"uint":       0,
		"int":        0,
		"uintptr":    0,
		"string":     0,
		"bool":       0,
		"error":      0,
	}
	_, ok := m[t]
	return ok
}

func MatchPathParam(s string) []string {
	reg := regexp.MustCompile(`{.*?}`)
	if reg == nil {
		panic("regexp err")
	}
	res := reg.FindAllString(s, -1)
	for i, _ := range res {
		res[i] = res[i][1 : len(res[i])-1]
	}
	return res
}

//func FindTagAndComment(structType *ast.StructType, tagName string) []TagMsg {
//	TagMsgs := make([]TagMsg,0,0)
//	ast.Inspect(structType, func(node ast.Node) bool {
//		field,ok := node.(*ast.Field)
//		if ok {
//			if field.Tag == nil {
//				return false
//			}
//			tagTool := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
//			value, ok := tagTool.Lookup(tagName)
//			if ok {
//				msg := TagMsg{
//					TagValue: value,
//					Comment:  field.Doc.Text(),
//				}
//				TagMsgs = append(TagMsgs, msg)
//			}
//			return false
//		}
//		return true
//	})
//	return TagMsgs
//}

//func FindTagAndCommentByPkg(pkg *packages.Package, structType *ast.StructType, tagName string) []TagMsg {
//	TagMsgs := make([]TagMsg,0,0)
//	ast.Inspect(structType, func(node ast.Node) bool {
//		switch ty := node.(type) {
//		case *ast.Ident:
//			st, ok := pkg.TypesInfo.TypeOf(ty).Underlying().(*types.Struct)
//			if ok {
//				st.Field(0).Pkg()
//				if st.NumFields() > 0 {
//
//				}
//			}
//		case *ast.Field:
//			tagTool := reflect.StructTag(ty.Tag.Value[1 : len(ty.Tag.Value)-1])
//			value, ok := tagTool.Lookup(tagName)
//			if ok {
//				msg := TagMsg{
//					TagValue: value,
//					Comment:  ty.Doc.Text(),
//				}
//				TagMsgs = append(TagMsgs, msg)
//			}
//			return true
//		}
//		return true
//	}
//	return TagMsgs
//}

type ImportMsg struct {
	Dir       string
	AliseName string
	PkgName   string
}

func ParseImport(file *ast.File) map[string]ImportMsg {
	m := make(map[string]ImportMsg)
	ast.Inspect(file, func(node ast.Node) bool {
		if importSpec, ok := node.(*ast.ImportSpec); ok {
			v := importSpec.Path.Value
			p, err := build.Import(strings.ReplaceAll(v, `"`, ""), "/", build.FindOnly)
			if err != nil {
				log.Fatalln(err)
			}
			name := p.Name
			if name == "" {
				name = path.Base(strings.ReplaceAll(v, `"`, ""))
			}
			msg := ImportMsg{
				Dir:       p.Dir,
				AliseName: "",
				PkgName:   name,
			}
			if importSpec.Name == nil {
				msg.AliseName = name
				m[name] = msg
				return true
			}
			if importSpec.Name.Name == "." || importSpec.Name.Name == "_" {
				return true
			}
			msg.AliseName = importSpec.Name.Name
			m[msg.AliseName] = msg
		}
		return true
	})
	return m
}

type LoopStruct struct {
	importMap map[string]ImportMsg
	LocalPkg  *ast.Package
	LocalFset *token.FileSet
	LocalFile *ast.File
}

type Quote struct {
	X         string
	Sel       string
	ImportMsg ImportMsg
}

func (q *Quote) IsSelect() bool {
	if q.X == "" {
		return false
	}
	return true
}

func NewLoopStruct(localPkg *ast.Package, localFset *token.FileSet, localFile *ast.File) *LoopStruct {
	l := &LoopStruct{LocalPkg: localPkg, LocalFset: localFset, LocalFile: localFile}
	l.importMap = ParseImport(l.LocalFile)
	return l
}

func (l *LoopStruct) FindQuote(t ast.Node) []Quote {
	fieldTypeNodes := make([]ast.Node, 0, 0)
	structTypeNodes := make([]ast.Node, 0, 0)
	quotes := make([]Quote, 0, 0)
	ast.Inspect(t, func(node ast.Node) bool {
		field, ok := node.(*ast.Field)
		if ok {
			structType, ok := field.Type.(*ast.StructType)
			if ok {
				structTypeNodes = append(structTypeNodes, structType)
				return false
			}
			fieldTypeNodes = append(fieldTypeNodes, field.Type)
			return false
		}
		return true
	})

	for _, fieldTypeNode := range fieldTypeNodes {
		ast.Inspect(fieldTypeNode, func(node ast.Node) bool {
			structType, ok := node.(*ast.StructType)
			if ok {
				structTypeNodes = append(structTypeNodes, structType)
				return false
			}

			se, ok := node.(*ast.SelectorExpr)
			if ok {
				x := se.X.(*ast.Ident).Name
				sel := se.Sel.Name
				quotes = append(quotes, Quote{x, sel, l.importMap[x]})
				return false
			}

			ident, ok := node.(*ast.Ident)
			if ok {
				hasBaseType := JudgeBuiltInType(ident.Name)
				if !hasBaseType {
					sel := ident.Name
					quotes = append(quotes, Quote{"", sel, ImportMsg{}})
				}
				return false
			}
			return true
		})
	}

	for _, structTypeNode := range structTypeNodes {
		quotes = append(quotes, l.FindQuote(structTypeNode)...)
	}

	return quotes
}

type Seter interface {
	GetKey() string
}

type Set struct {
	hash       map[string]struct{}
	containers []interface{}
}

func NewSet() *Set {
	return &Set{hash: map[string]struct{}{}, containers: []interface{}{}}
}

func (s *Set) Add(sets ...Seter) {
	for _, set := range sets {
		key := set.GetKey()

		_, ok := s.hash[key]
		if !ok {
			s.hash[key] = struct{}{}
			s.containers = append(s.containers, set)
		}
	}
}

func (s *Set) Get() []interface{} {
	return s.containers
}

func FindTagAndCommentByField(pkg *packages.Package, file *ast.File, field *ast.Field, TagName string) []TagMsg {
	_, ok := pkg.TypesInfo.TypeOf(field.Type).Underlying().(*types.Struct)
	if !ok {
		return []TagMsg{}
	}

	var findFile *ast.File
	var findPkg *packages.Package
	var findStruct *ast.StructType

	//switch t := field.Type.(type) {
	//// local pkg struct
	//case *ast.Ident:
	//	findPkg = pkg
	//	findFile, findStruct =  FindStructTypeByName(pkg, t.Name)
	//// selector pkg
	//case *ast.SelectorExpr:
	//	path := FindImportPath(file.Imports, t.X.(*ast.Ident).Name)
	//	findPkg = pkg.Imports[path]
	//	findFile, findStruct = FindStructTypeByName(findPkg, t.Sel.Name)
	//// struct
	//case *ast.StructType:
	//	findFile = file
	//	findStruct = t
	//	findPkg = pkg
	//}
	findPkg, findFile, _,findStruct = FindStructByExpr(pkg, file, field.Type)
	return FindTagAndComment(findPkg, findFile, findStruct, TagName)
}

func FindStructByExpr(pkg *packages.Package, file *ast.File, expr ast.Expr) (*packages.Package, *ast.File, *ast.TypeSpec,*ast.StructType) {
	_, ok := pkg.TypesInfo.TypeOf(expr).Underlying().(*types.Struct)
	if !ok {
		return nil, nil, nil,nil
	}
	switch t := expr.(type) {
	// local pkg struct
	case *ast.Ident:
		findFile, findType, findStruct := FindStructTypeByName(pkg, t.Name)
		return pkg, findFile,findType, findStruct
	// selector pkg
	case *ast.SelectorExpr:
		path := FindImportPath(file.Imports, t.X.(*ast.Ident).Name)
		findPkg := pkg.Imports[path]
		findFile, findType, findStruct := FindStructTypeByName(findPkg, t.Sel.Name)
		return findPkg, findFile,findType, findStruct
	// struct
	case *ast.StructType:
		return pkg, file, nil,t
	}
	return nil, nil, nil,nil
}

func GetFileNameByPos(fset *token.FileSet, pos token.Pos) string {
	filePath := fset.Position(pos).Filename
	_, fileName := path.Split(filePath)
	return fileName
}
