package pkg

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"html/template"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

//go:embed internal/templateV2/gen_api.tmpl
var gen_api_tmpl string

const GenMark string = "@GenApi"

func ParseFuncApi(modName string, src string, dest string) ParseContext {
	fset, pkgs := LoadPkgs(src)
	context := ParseContext{
		LocalModuleName: modName,
		ParsePath:       src,
		Fset:            fset,
		Pkgs:            pkgs,
	}
	context.Enter()
	GenApi(context.ApiMap, dest)
	return context
}

type InHas struct {
	HasBody   bool
	BodyMsg   FieldMsg
	HasUri    bool
	UriMsg    FieldMsg
	HasHeader bool
	HeaderMsg FieldMsg
	HasQuery  bool
	QueryMsg  FieldMsg
}

type FuncRouter struct {
	Path   string
	Method string
}

type ApiMsg struct {
	SrcPkgName string
	FuncRouter FuncRouter
	Doc        *ast.CommentGroup
	In         struct {
		ObjectMsg ObjectMsg
		InHas     InHas
	}
	Out ObjectMsg
}

type ObjectMsg struct {
	PkgName string
	RawName string
	// Selector 中的Sel
	SelectorSel string
	// Selector 中的X
	SelectorX string
	// true 则x sel都有，false只有 sel
	IsSelector bool
	// isselector为true 则寻找引用包路径
	ObjectImportMsg ObjectImportMsg
}
type ParseContext struct {
	LocalModuleName string
	ParsePath       string
	Fset            *token.FileSet
	// pakName fileName funcName
	ApiMap map[string]map[string]map[string]ApiMsg
	Pkgs   map[string]*ast.Package
}

func (c *ParseContext) Enter() {

	for pkgName, pkg := range c.Pkgs {
		apiMap := make(map[string]map[string]map[string]ApiMsg, 0)
		apiMap[pkgName] = make(map[string]map[string]ApiMsg)
		for filePath, file := range pkg.Files {
			apiMap[pkgName][filePath] = make(map[string]ApiMsg, 0)
			filterFuncs := c.FileEnter(file)
			parseImport := ParseImport(file, c.Fset)
			for _, f := range filterFuncs {
				inObjectMsg := c.FindObjectMsg(f.Type.Params.List[1], file, filePath, parseImport)
				outObjectMsg := c.FindObjectMsg(f.Type.Results.List[0], file, filePath, parseImport)
				_, funcRouter := c.FuncHasSwaggerRouter(f.Doc)
				apiMsg := ApiMsg{
					SrcPkgName: pkgName,
					FuncRouter: funcRouter,
					Doc:        f.Doc,
					In: struct {
						ObjectMsg ObjectMsg
						InHas     InHas
					}{
						ObjectMsg: inObjectMsg,
					},
					Out: outObjectMsg,
				}
				apiMsg.In.InHas = c.FindInHas(apiMsg)

				apiMap[pkgName][filePath][f.Name.Name] = apiMsg
			}
		}
		c.ApiMap = apiMap
	}
}

func (c *ParseContext) FileEnter(f *ast.File) []*ast.FuncDecl {
	filterFunc := make([]*ast.FuncDecl, 0)
	ast.Inspect(f, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			hasRouter, _ := c.FuncHasSwaggerRouter(funcDecl.Doc)
			hasConform := c.ConformFormat(funcDecl)
			log.Println(fmt.Sprintf("func %v: hasRouter %v, hasConform %v", funcDecl.Name.Name, hasRouter, hasConform))
			if hasRouter && hasConform {
				filterFunc = append(filterFunc, funcDecl)
			}
			return false
		}
		return true
	})
	return filterFunc
}

func (c *ParseContext) FuncHasSwaggerRouter(doc *ast.CommentGroup) (bool, FuncRouter) {
	if doc == nil {
		return false, FuncRouter{}
	}
	for _, comment := range doc.List {
		fields := strings.Fields(comment.Text)
		if len(fields) < 4 {
			continue
		}
		if fields[0] == "//" && fields[1] == GenMark {
			return true, FuncRouter{
				Path:   fields[2],
				Method: strings.ToUpper(fields[4]),
			}
		}
	}
	return false, FuncRouter{}
}

// 符合 func Name(c *gin.context, in object) (out object, err error)
func (c *ParseContext) ConformFormat(f *ast.FuncDecl) bool {
	if f.Type.Params.NumFields() != 2 || f.Type.Results.NumFields() != 2 {
		return false
	}
	paramGinContext := f.Type.Params.List[0]
	if selectorExpr, ok := paramGinContext.Type.(*ast.StarExpr).X.(*ast.SelectorExpr); ok {
		if selectorExpr.X.(*ast.Ident).Name == "gin" && selectorExpr.Sel.Name == "Context" {
		}
	} else {
		return false
	}

	paramIn := f.Type.Params.List[1]
	if _, ok := paramIn.Type.(*ast.StarExpr); !ok {
		log.Fatalln(fmt.Sprintf("Func %v %v type not ptr", f.Name.Name, paramIn.Names[0].Name))
		return false
	}

	paramErr := f.Type.Results.List[1]
	if ident, ok := paramErr.Type.(*ast.Ident); ok {
		if ident.Name == "error" {
			return true
		}
	}
	return false
}

func (c *ParseContext) FindObjectMsg(field *ast.Field, file *ast.File, filePath string, msgMap map[string]ObjectImportMsg) ObjectMsg {
	msg := ObjectMsg{
		PkgName:     file.Name.Name,
		RawName:     NodeString(c.Fset, field.Type),
		SelectorSel: "",
		SelectorX:   "",
		IsSelector:  false,
	}

	has, x, sel := FieldIsSelector(field)
	if has {
		msg.IsSelector = true
		msg.SelectorX = x
		msg.SelectorSel = sel
		if fileImportMsg, ok := msgMap[x]; !ok {
			log.Fatalf("没有找到x %v \n", x)
		} else {
			msg.ObjectImportMsg = fileImportMsg
		}
		return msg
	}
	msg.SelectorSel = field.Type.(*ast.StarExpr).X.(*ast.Ident).Name
	log.Println(msg.RawName + " filepath: " + filePath)
	msg.ObjectImportMsg = ObjectImportMsg{RawImport: path.Join(c.LocalModuleName, c.ParsePath)}
	return msg
}

func FieldIsSelector(field *ast.Field) (bool, string, string) {
	var has bool
	var x string
	var sel string
	ast.Inspect(field, func(node ast.Node) bool {
		if selectorExpr, ok := node.(*ast.SelectorExpr); ok {
			has = true
			x = selectorExpr.X.(*ast.Ident).Name
			sel = selectorExpr.Sel.Name
			return false
		}
		return true
	})
	return has, x, sel
}

func (c *ParseContext) LocalPackagePath(file *ast.File) {

}

type ObjectImportMsg struct {
	RawImport string
	PathDir   string
	AliseName string
	PkgName   string
}

func ParseImport(file *ast.File, fset *token.FileSet) map[string]ObjectImportMsg {
	fileImportMsgs := make(map[string]ObjectImportMsg, 0)
	ast.Inspect(file, func(node ast.Node) bool {
		if importSpec, ok := node.(*ast.ImportSpec); ok {
			dir, name := FindPkg(importSpec.Path.Value)
			msg := ObjectImportMsg{PathDir: dir, PkgName: name, RawImport: NodeString(fset, node)}
			if importSpec.Name == nil {
				msg.AliseName = name
				fileImportMsgs[name] = msg
				return true
			}
			if importSpec.Name.Name == "." || importSpec.Name.Name == "_" {
				return true
			}
			msg.AliseName = importSpec.Name.Name
			fileImportMsgs[importSpec.Name.Name] = msg
			return true
		}
		return true
	})
	return fileImportMsgs
}
func NodeString(fset *token.FileSet, node interface{}) string {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, node)
	if err != nil {
		spew.Dump(node)
		log.Fatalln(err.Error())
	}
	return buf.String()
}

func FindPkg(pkgPath string) (string, string) {
	p, err := build.Import(strings.ReplaceAll(pkgPath, `"`, ""), "/", build.FindOnly)
	if err != nil {
		log.Fatalln(err.Error())
	}
	if p.Name == "" {
		return p.Dir, path.Base(strings.ReplaceAll(pkgPath, `"`, ""))
	}
	return p.Dir, p.Name
}

func LoadPkgs(path string) (*token.FileSet, map[string]*ast.Package) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, func(info os.FileInfo) bool {
		return true
	}, parser.ParseComments)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return fset, pkgs
}

func (c *ParseContext) FindInHas(api ApiMsg) InHas {
	if api.In.ObjectMsg.IsSelector {
		fset, pkgs := LoadPkgs(api.In.ObjectMsg.ObjectImportMsg.PathDir)
		fileAst, structType, has := FindStructByPkgs(pkgs[api.In.ObjectMsg.SelectorX], api.In.ObjectMsg.SelectorSel)
		if !has {
			log.Fatalln(fmt.Sprintf("not find struct %v by pkg %v", api.In.ObjectMsg.SelectorSel, api.In.ObjectMsg.SelectorX))
		}
		parseImport := ParseImport(fileAst, fset)
		inHas := FindInByStructType(structType, fset, pkgs[api.In.ObjectMsg.SelectorX], parseImport, api.In.ObjectMsg.ObjectImportMsg.RawImport)
		return inHas
	} else {
		fileAst, structType, has := FindStructByPkgs(c.Pkgs[api.In.ObjectMsg.PkgName], api.In.ObjectMsg.SelectorSel)
		if !has {
			log.Fatalln(fmt.Sprintf("not find struct %v by pkg %v", api.In.ObjectMsg.SelectorSel, api.In.ObjectMsg.PkgName))
		}
		parseImport := ParseImport(fileAst, c.Fset)
		inHas := FindInByStructType(structType, c.Fset, c.Pkgs[c.LocalModuleName], parseImport, path.Join(c.LocalModuleName, path.Dir(c.ParsePath)))
		return inHas

	}

}

func FindStructByPkgs(pkg *ast.Package, structName string) (*ast.File, *ast.StructType, bool) {
	var structType *ast.StructType
	var fileAst *ast.File

	for _, file := range pkg.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if typeSpec, ok := node.(*ast.TypeSpec); ok {
				if typeSpec.Name.Name == structName {
					if st, ok := typeSpec.Type.(*ast.StructType); ok {
						structType = st
						fileAst = file
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
	//if structType == nil {
	//	log.Fatalln(fmt.Sprintf("pkg %v中, 没有找到%v struct", pkgName, structName))
	//}
	if structType == nil {
		return fileAst, structType, false
	}
	return fileAst, structType, true
}

func FindInByStructType(structType *ast.StructType, fset *token.FileSet, pkg *ast.Package, importMsgs map[string]ObjectImportMsg, byFindImportPath string) InHas {
	inHas := InHas{}
	for _, field := range structType.Fields.List {
		for _, ident := range field.Names {
			switch ident.Name {
			case "Query":
				inHas.HasQuery = true
				inHas.QueryMsg = FieldTypeIsStruct(field, fset, pkg, importMsgs, byFindImportPath)
			case "Body":
				inHas.HasBody = true
				inHas.BodyMsg = FieldTypeIsStruct(field, fset, pkg, importMsgs, byFindImportPath)
			case "Uri":
				inHas.HasUri = true
				inHas.UriMsg = FieldTypeIsStruct(field, fset, pkg, importMsgs, byFindImportPath)
			case "Header":
				inHas.HasHeader = true
				inHas.HeaderMsg = FieldTypeIsStruct(field, fset, pkg, importMsgs, byFindImportPath)
			}
		}

	}
	return inHas
}

type FieldMsg struct {
	Raw        string
	ImportPath string
}

func FieldTypeIsStruct(field *ast.Field, fset *token.FileSet, pkg *ast.Package, importMsg map[string]ObjectImportMsg, byFindImportPath string) FieldMsg {
	fieldMsg := FieldMsg{Raw: NodeString(fset, field.Type)}
	switch field.Type.(type) {
	case *ast.SelectorExpr:
		selectorExpr, _ := field.Type.(*ast.SelectorExpr)
		x := selectorExpr.X.(*ast.Ident).Name
		//sel := selectorExpr.Sel.Name
		fieldMsg.ImportPath = importMsg[x].RawImport
		return fieldMsg
	case *ast.Ident:
		ident, _ := field.Type.(*ast.Ident)
		name := ident.Name
		_, _, has := FindStructByPkgs(pkg, name)
		if !has {
			log.Fatalln(fmt.Sprintf("local pkg not found struct %v", ident.Name))
		}
		fieldMsg.ImportPath = byFindImportPath
		return fieldMsg
	}

	log.Fatalln("field not structType")
	return fieldMsg
}

func GenApi(apiMap map[string]map[string]map[string]ApiMsg, dest string) {
	tpl, err := template.New("gen_api").Parse(gen_api_tmpl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	assets := assets{
		dirs: []string{
			filepath.Join(dest),
		},
	}

	for _, fileMap := range apiMap {
		for fileName, funcMap := range fileMap {
			log.Println("filename" + fileName)
			b := bytes.NewBuffer(nil)
			err := tpl.Execute(b, funcMap)
			if err != nil {
				log.Fatalln(err.Error())
			}
			assets.files = append(assets.files, file{
				path:    filepath.Join(dest, path.Base(fileName)),
				content: b.Bytes(),
			})
		}
	}

	if err := assets.write(); err != nil {
		log.Fatalln(err.Error())
	}

	err = assets.formatGo()
	if err != nil {
		log.Fatalln(err.Error())
	}

}
