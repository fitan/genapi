package gen_apiV2

import (
	"go/ast"
	"go/token"
)
const GenMark string = "@GenApi"

type ApiContext struct {
	PkgName string
	Fset    *token.FileSet
	Pkg     *ast.Package
	Files   map[string]*FileContext
}

func NewApiContext() *ApiContext {
	return &ApiContext{}
}


func (c *ApiContext)Load(dir string)  {
	pkgName, pkg, fset := LoadPkg(dir)
	c.PkgName = pkgName
	c.Fset = fset
	c.Pkg = pkg
}

func (c *ApiContext)Parse() {
	files := make(map[string]*FileContext, 0)
	for _, f := range c.Pkg.Files {
		fc := NewFileContext(c.PkgName,c.Pkg,c.Fset,f)
		fc.Parse()
		files[f.Name.Name] = fc
	}
	c.Files = files
}
