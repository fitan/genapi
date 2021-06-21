package plugins

import (
	_ "embed"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const (
	CasbinKeyserName     = "CasbinKeyser"
	CasbinListKeyserName = "CasbinListKeyser"
)

//go:embed interfacepkg/casbin.go
var casbinFile string
var casbinAstFile *ast.File
var casbinInfo *types.Info

var FindPluginInterfaceMap = make(map[string]*types.Interface)

func init() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "casbin.go", casbinFile, parser.AllErrors)
	if err != nil {
		log.Panic(err)
	}
	casbinAstFile = f
	casbinInfo = &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue, 0),
		Defs:       make(map[*ast.Ident]types.Object, 0),
		Uses:       make(map[*ast.Ident]types.Object, 0),
		Implicits:  make(map[ast.Node]types.Object, 0),
		Selections: make(map[*ast.SelectorExpr]*types.Selection, 0),
		Scopes:     make(map[ast.Node]*types.Scope, 0),
		InitOrder:  make([]*types.Initializer, 0, 0),
	}
	_, err = new(types.Config).Check("casbin.go", fset, []*ast.File{casbinAstFile}, casbinInfo)
	if err != nil {
		log.Panic(err)
	}

	log.Println(FindPluginInterfaceMap)
	FindInterface()
	log.Println("find plugin interface map", FindPluginInterfaceMap)
}

func FindInterface() {
	ast.Inspect(casbinAstFile, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
				//t,ok := casbinInfo.TypeOf(interfaceType).(*types.Interface)
				//if ok {
				//	log.Println(t)
				//} else {
				//	log.Println("false")
				//}
				FindPluginInterfaceMap[typeSpec.Name.Name] = casbinInfo.TypeOf(interfaceType).(*types.Interface)
				return false
			}
		}
		return true
	})
}
