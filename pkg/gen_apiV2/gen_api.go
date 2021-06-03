package gen_apiV2

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)
const GenMark string = "@GenApi"

type ApiContext struct {
	PkgName string
	Fset    *token.FileSet
	Pkg     *packages.Package
	Files   map[string]*FileContext
}

func NewApiContext() *ApiContext {
	return &ApiContext{}
}


func (c *ApiContext)Load(dir string)  {
	pkgName, pkg, fset := LoadPkgV2(dir)
	c.PkgName = pkgName
	c.Fset = fset
	c.Pkg = pkg
}

func (c *ApiContext)Parse() {
	files := make(map[string]*FileContext, 0)
	for _, f := range c.Pkg.Syntax {
		fc := NewFileContext(c.PkgName,c.Pkg,f)
		fc.Parse()
		if len(fc.Funcs) != 0 {
			files[GetFileNameByPos(c.Pkg.Fset, f.Pos())] = fc
		}
	}
	c.Files = files
}
