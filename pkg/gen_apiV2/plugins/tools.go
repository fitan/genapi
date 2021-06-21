package plugins

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"path"
)

const (
	CasbinKeyserName = "CasbinKeyser"
	CasbinListKeyserName = "CasbinListKeyser"
	mode packages.LoadMode = packages.NeedName |
		packages.NeedTypes |
		packages.NeedSyntax |
		packages.NeedTypesInfo |
		packages.NeedImports |
		packages.NeedModule |
		packages.NeedExportsFile |
		packages.NeedTypesSizes |
		packages.NeedDeps
)

var interfacePkg *packages.Package



var FindPluginInterfaceMap map[string]*types.Interface

func init() {
	var dir = "interfacepkg"
	pkgName := path.Base(dir)
	var fset = token.NewFileSet()
	cfg := &packages.Config{Fset: fset, Mode: mode}
	pkgs, err := packages.Load(cfg, dir)
	if err != nil {
		log.Panicln(err)
	}
	for _, pkg := range pkgs {
		if path.Base(pkg.String()) == pkgName {
			interfacePkg = pkg
			return
		}
	}
	FindInterface()
}

func FindInterface()  {
	for _, f := range interfacePkg.Syntax {
		ast.Inspect(f, func(node ast.Node) bool {
			if typeSpec, ok := node.(*ast.TypeSpec); ok {
				if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
					FindPluginInterfaceMap[typeSpec.Name.Name] = interfacePkg.TypesInfo.TypeOf(interfaceType).(*types.Interface)
					return false
				}
			}
			return true
		})
	}
}




