package plugins

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

//go:embed interfacepkg/casbin.go
var casbinFile string

type FileMsg struct {
	Name string
	Src  string
}

var interfaceFiles = []FileMsg{{
	Name: "casbin.go",
	Src:  casbinFile,
}}

var FindPluginInterfaceMap = make(map[string]*types.Interface)

func init() {
	for _, interfaceFile := range interfaceFiles {

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, interfaceFile.Name, interfaceFile.Src, parser.AllErrors)
		if err != nil {
			log.Panic(err)
		}
		info := &types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue, 0),
			Defs:       make(map[*ast.Ident]types.Object, 0),
			Uses:       make(map[*ast.Ident]types.Object, 0),
			Implicits:  make(map[ast.Node]types.Object, 0),
			Selections: make(map[*ast.SelectorExpr]*types.Selection, 0),
			Scopes:     make(map[ast.Node]*types.Scope, 0),
			InitOrder:  make([]*types.Initializer, 0, 0),
		}
		_, err = new(types.Config).Check(interfaceFile.Name, fset, []*ast.File{f}, info)
		if err != nil {
			log.Panic(err)
		}

		log.Println(FindPluginInterfaceMap)
		findInterface(f, info)
		log.Println("find plugin interface map", FindPluginInterfaceMap)
	}
}

func findInterface(f *ast.File, info *types.Info) {
	ast.Inspect(f, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
				FindPluginInterfaceMap[typeSpec.Name.Name] = info.TypeOf(interfaceType).(*types.Interface)
				return false
			}
		}
		return true
	})
}

func CheckHasInterface(t types.Type, interfaceName string) bool {
	if i, ok := FindPluginInterfaceMap[interfaceName]; ok {
		return types.Implements(t, i)
	}
	log.Fatalln("not found " + interfaceName)
	return false
}

type PluginTemplate struct {
	Has          bool
	Keys         map[string]string
	InBindBefor  HandlerTemplate
	InBindAfter  HandlerTemplate
	OutBindBefor HandlerTemplate
	OutBindAfter HandlerTemplate
}

type HandlerTemplate struct {
	ImportPath string
	Template   string
}
