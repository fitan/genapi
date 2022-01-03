package gen_apiV2

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"log"
	"path"
	"reflect"
	"regexp"
	"strings"
)

const mode packages.LoadMode = packages.NeedName |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedImports |
	packages.NeedModule |
	packages.NeedExportsFile |
	packages.NeedTypesSizes |
	packages.NeedDeps

func LoadPackages(dir string) (string, *packages.Package, *token.FileSet) {
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
	log.Fatalln("not found pkg")
	return "", nil, nil
}

// 判断是否是selecter类型
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
		log.Panicln(err.Error())
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
		switch t := c.Node().(type) {
		case *ast.SelectorExpr:
			return false

		case *ast.Ident:
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {
					tmp := ast.SelectorExpr{X: ast.NewIdent(selectName), Sel: ast.NewIdent(c.Node().(*ast.Ident).Name)}
					c.Replace(&tmp)
				}
			} else {
				spew.Dump(c.Node())
				if ok := JudgeBuiltInType(c.Node().(*ast.Ident).Name); !ok {
					tmp := ast.SelectorExpr{X: ast.NewIdent(selectName), Sel: ast.NewIdent(c.Node().(*ast.Ident).Name)}
					c.Replace(&tmp)
				}
			}
		}
		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
}

// go 的内置基础类型
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

type ImportMsg struct {
	Dir       string
	AliseName string
	PkgName   string
}

//type LoopStruct struct {
//	importMap map[string]ImportMsg
//	LocalPkg  *ast.Package
//	LocalFset *token.FileSet
//	LocalFile *ast.File
//}

//type Quote struct {
//	X         string
//	Sel       string
//	ImportMsg ImportMsg
//}
//
//func (q *Quote) IsSelect() bool {
//	if q.X == "" {
//		return false
//	}
//	return true
//}

//func (l *LoopStruct) FindQuote(t ast.Node) []Quote {
//	fieldTypeNodes := make([]ast.Node, 0, 0)
//	structTypeNodes := make([]ast.Node, 0, 0)
//	quotes := make([]Quote, 0, 0)
//	ast.Inspect(t, func(node ast.Node) bool {
//		field, ok := node.(*ast.Field)
//		if ok {
//			structType, ok := field.Type.(*ast.StructType)
//			if ok {
//				structTypeNodes = append(structTypeNodes, structType)
//				return false
//			}
//			fieldTypeNodes = append(fieldTypeNodes, field.Type)
//			return false
//		}
//		return true
//	})
//
//	for _, fieldTypeNode := range fieldTypeNodes {
//		ast.Inspect(fieldTypeNode, func(node ast.Node) bool {
//			structType, ok := node.(*ast.StructType)
//			if ok {
//				structTypeNodes = append(structTypeNodes, structType)
//				return false
//			}
//
//			se, ok := node.(*ast.SelectorExpr)
//			if ok {
//				x := se.X.(*ast.Ident).Name
//				sel := se.Sel.Name
//				quotes = append(quotes, Quote{x, sel, l.importMap[x]})
//				return false
//			}
//
//			ident, ok := node.(*ast.Ident)
//			if ok {
//				hasBaseType := JudgeBuiltInType(ident.Name)
//				if !hasBaseType {
//					sel := ident.Name
//					quotes = append(quotes, Quote{"", sel, ImportMsg{}})
//				}
//				return false
//			}
//			return true
//		})
//	}
//
//	for _, structTypeNode := range structTypeNodes {
//		quotes = append(quotes, l.FindQuote(structTypeNode)...)
//	}
//
//	return quotes
//}

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
	findPkg, findFile, _, findStruct := FindStructByExpr(pkg, file, field.Type)
	return FindTagAndCommentByStruct(findPkg, findFile, findStruct, TagName)
	//var findFile *ast.File
	//var findPkg *packages.Package
	//var findStruct *ast.StructType
	//switch field.Type.(type) {
	//case *ast.Ident, *ast.SelectorExpr:
	//	findPkg, findFile, _, findStruct = FindStructByExpr(pkg, file, field.Type)
	//	return FindTagAndCommentByStruct(findPkg, findFile, findStruct, TagName)
	//case *ast.StructType:
	//	return FindTagAndCommentByStruct(findPkg, findFile, findStruct, TagName)
	//}
	//
	//panic("FindTagAndCommentByField: 未知类型")
	//log.Printf("findTagAndComment file: %s, field: %s, tagName: %s", file.Name.Name, field.Names[0].Name, TagName)
	//log.Println(field.Type)
	//log.Println(pkg.TypesInfo.TypeOf(field.Type))
	//fmt.Println(pkg.TypesInfo.TypeOf(field.Type).Underlying().String())
	//_, ok := pkg.TypesInfo.TypeOf(field.Type).Underlying().(*types.Struct)
	//if !ok {
	//	return []TagMsg{}
	//}
	//
	//
	//findPkg, findFile, _, findStruct = FindStructByExpr(pkg, file, field.Type)
	//return FindTagAndCommentByStruct(findPkg, findFile, findStruct, TagName)
}

func FindStructByExpr(pkg *packages.Package, file *ast.File, expr ast.Expr) (*packages.Package, *ast.File, *ast.TypeSpec, *ast.StructType) {
	//_, ok := pkg.TypesInfo.TypeOf(expr).Underlying().(*types.Struct)
	//if !ok {
	//	return nil, nil, nil, nil
	//}
	switch t := expr.(type) {
	// struct 在同一个pkg里面
	case *ast.Ident:
		findFile, findType, findStruct := FindStructTypeByName(pkg, t.Name)
		return pkg, findFile, findType, findStruct
	// struct 是selector类型， 在另外的pkg里面
	case *ast.SelectorExpr:
		log.Printf("find import path. path: %v, pkgName: %v, file: %v, typeName: %v", pkg.PkgPath, pkg.Name, GetFileNameByPos(pkg.Fset, file.Pos()), Node2String(pkg.Fset, t))
		path := FindImportPath(file.Imports, t.X.(*ast.Ident).Name)
		findPkg := pkg.Imports[path]
		findFile, findType, findStruct := FindStructTypeByName(findPkg, t.Sel.Name)
		return findPkg, findFile, findType, findStruct
	// 本身就是struct类型
	case *ast.StructType:
		return pkg, file, nil, t
	}
	// 未知的状态
	return nil, nil, nil, nil
}

func GetFileNameByPos(fset *token.FileSet, pos token.Pos) string {
	filePath := fset.Position(pos).Filename
	_, fileName := path.Split(filePath)
	return fileName
}

func FindPkgBySelector(pkg *packages.Package, file *ast.File, selector *ast.SelectorExpr) *packages.Package {
	log.Printf("find import path. path: %v, pkgName: %v, file: %v, typeName: %v", pkg.PkgPath, pkg.Name, GetFileNameByPos(pkg.Fset, file.Pos()), Node2String(pkg.Fset, selector))
	path := FindImportPath(file.Imports, selector.X.(*ast.Ident).Name)
	log.Println("find import path: ", path)
	return pkg.Imports[path]
}

func SpliceType(pkg *packages.Package, file *ast.File, node ast.Node) bool {
	replace := false
	astutil.Apply(node, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		case *ast.Field:
			fmt.Println("ent ast field")
			if t.Tag == nil {
				fmt.Println("out ast field")
				return false
			}

			if tags, ok := reflect.StructTag(t.Tag.Value[1 : len(t.Tag.Value)-1]).Lookup("json"); ok {
				for _, tag := range strings.Split(tags, ",") {
					if tag == "-" {
						fmt.Println("out ast field")
						return false
					}
				}
			}
			fmt.Println("out ast field")
			return true

		case *ast.Ident:
			defer fmt.Println("out ast ident")
			fmt.Println("ent ast ident")
			//if pkg.TypesInfo.TypeOf(t) != nil {
			//	switch InfoT := pkg.TypesInfo.TypeOf(t).Underlying().(type) {
			//	case *types.Struct:
			//		fmt.Println("struct: ", t.Name, InfoT.String())
			//	case *types.Basic:
			//		switch tt := c.Parent().(type) {
			//		case *ast.Field:
			//			fmt.Println("t: ", t.Name, "parent: ", tt.Names)
			//			fmt.Println("basic: ", t.Name, InfoT.String())
			//		default:
			//			fmt.Println("basic: ", t.Name, InfoT.String())
			//
			//		}
			//	case *types.Named:
			//		fmt.Println("named: ",  t.Name, InfoT.String())
			//
			//
			//
			//
			//
			//	}
			//}
			fmt.Println("t name: ", t.Name, "t obj: ", t.Obj)
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {
					//fmt.Println("ident name: ", t.Name)
					fmt.Println("find need name: ", t.Name)
					f, findTs := FindTypeByName(pkg, t.Name)
					file.Imports = append(file.Imports, f.Imports...)
					fmt.Println("find: ", findTs)
					fmt.Println("find ident type: ", findTs.Type, "name: ", t.Name)
					replace = true
					c.Replace(findTs.Type)
				}
			} else {
				if !JudgeBuiltInType(t.Name) {
					f, findTs := FindTypeByName(pkg, t.Name)
					file.Imports = append(file.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			}
			//if t.Name == "Code" {
			//	fmt.Println("code obj ", t.Obj)
			//}
			//if t.Name == "AliaseInt" {
			//	fmt.Println("AliaseInt obj", t.Obj)
			//}
			//fmt.Println("enter ident: ", t.Name)
			//if t.Obj == nil {
			//	fmt.Println("obj is nil: ", t.Name)
			//	if !JudgeBuiltInType(t.Name) {
			//		fmt.Println("ident name: ", t.Name)
			//		_, findTs := FindTypeByName(pkg, t.Name)
			//		//	fmt.Println("find ident type: ", findTs.Type)
			//		c.Replace(findTs.Type)
			//	}
			//}
			//if _,ok := c.Parent().(*ast.Field);ok {
			//	return true
			//}
			//fmt.Println("ident Name ", t.Name)
			//if !JudgeBuiltInType(t.Name) {
			//	_, findTs := FindTypeByName(pkg, t.Name)
			//	fmt.Println("find ident type: ", findTs.Type)
			//	c.Replace(findTs.Type)
			//}
		//
		//
		//
		case *ast.SelectorExpr:
			defer fmt.Println("ent ast selector expr")
			fmt.Println("ent ast selector expr")
			if t.X.(*ast.Ident).Name == "time" && t.Sel.Name == "Time" {
				return false
			}
			log.Printf("find import path. path: %v, pkgName: %v, file: %v, typeName: %v", pkg.PkgPath, pkg.Name, GetFileNameByPos(pkg.Fset, file.Pos()), Node2String(pkg.Fset, t))
			path := FindImportPath(file.Imports, t.X.(*ast.Ident).Name)
			findPkg := pkg.Imports[path]
			if findPkg.Imports != nil {
				for index, importPath := range findPkg.Imports {
					pkg.Imports[index] = importPath
				}
			}
			for _, synx := range findPkg.Syntax {
				pkg.Syntax = append(pkg.Syntax, synx)
			}
			fmt.Printf("FindTypeByName: %v", Node2String(pkg.Fset, t))
			f, findTs := FindTypeByName(findPkg, t.Sel.Name)
			file.Imports = append(file.Imports, f.Imports...)

			replace = true
			c.Replace(findTs.Type)
			return false
		}
		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})

	//fmt.Println("splice struct :   ", Node2String(pkg.Fset, node))
	return replace
}

func Struct2Ts(pkg *packages.Package, file *ast.File, node ast.Node, objName string) string {
	ok := SpliceType(pkg, file, node)
	for ok {
		fmt.Println("ent struct 2 ts")
		ok = SpliceType(pkg, file, node)
	}
	//if SpliceType(pkg, file, node) {
	//	SpliceType(pkg, file, node)
	//}
	s := fmt.Sprintf("type %s %s", objName, Node2String(pkg.Fset, node))
	fmt.Println("convert struct2ts")
	return Convert(s)
}

func WarpResult2Ts(pkg *packages.Package, file *ast.File, node ast.Node, objName string) string {
	ok := SpliceType(pkg, file, node)
	for ok {
		fmt.Println("ent warp 2 ts")
		ok = SpliceType(pkg, file, node)
	}
	ResutlStr := "type %s struct {\nCode int `json:\"code\"`\nData %s `json:\"data\"`\nErr  string `json:\"err\"`}"
	s := fmt.Sprintf(ResutlStr, objName, Node2String(pkg.Fset, node))
	fmt.Println("convert warpresult2ts")
	return Convert(s)
}
