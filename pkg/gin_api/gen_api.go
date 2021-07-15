package gen_apiV2

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

const GenMark string = "@GenApi"
const CasbinMark string = "@Casbin"

type ApiContext struct {
	PkgName      string
	Fset         *token.FileSet
	Pkg          *packages.Package
	Files        map[string]*FileContext
	ReginsterMap map[string][]Func
}

func NewApiContext() *ApiContext {
	return &ApiContext{}
}

func (c *ApiContext) Load(dir string) {
	pkgName, pkg, fset := LoadPackages(dir)
	c.PkgName = pkgName
	c.Fset = fset
	c.Pkg = pkg
}

func (c *ApiContext) Parse(option ParseOption) {
	files := make(map[string]*FileContext, 0)
	for _, f := range c.Pkg.Syntax {
		fc := NewFileContext(c.PkgName, c.Pkg, f)
		fc.Parse(option)
		if len(fc.Funcs) != 0 {
			files[GetFileNameByPos(c.Pkg.Fset, f.Pos())] = fc
		}
	}
	c.Files = files
	c.FuncsToMapByRouterGroupKey()
}

func (c *ApiContext) FuncsToMapByRouterGroupKey() {
	reginsterM := make(map[string][]Func)
	for _, file := range c.Files {
		for index, _ := range file.Funcs {
			key := file.Funcs[index].Router.RouterGroupKey
			if _, ok := reginsterM[key]; ok {
				reginsterM[key] = append(reginsterM[key], file.Funcs[index])
			} else {
				reginsterM[key] = make([]Func, 0, 0)
				reginsterM[key] = append(reginsterM[key], file.Funcs[index])
			}
		}
	}
	c.ReginsterMap = reginsterM
}
