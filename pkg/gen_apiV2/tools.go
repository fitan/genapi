package gen_apiV2

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"io/fs"
	"log"
	"path"
	"strings"
)

func LoadImport(it string) string {
	p, err := build.Import(strings.ReplaceAll(it, `"`, ""), "/", build.FindOnly)
	if err != nil {
		log.Fatalln(err)
	}
	return p.Dir
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

func FindStructByPkg(pkg *ast.Package, structName string) (*ast.File, *ast.StructType, bool) {
	var structType *ast.StructType
	var f *ast.File
	fmt.Println("pkg: ",pkg,"structName: ", structName)
	for _, file := range pkg.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if typeSpec, ok := node.(*ast.TypeSpec);ok {
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

func FindStructByDir(dir string, structName string) (*ast.Package, *token.FileSet, *ast.File, *ast.StructType, bool) {
	_, pkg, fset := LoadPkg(dir)
	f, structType, has := FindStructByPkg(pkg, structName)
	return pkg,fset,f,structType,has
}

func IsSelector(field *ast.Field) (string, string, bool) {
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
	spew.Dump(field)
	ident, ok := field.Type.(*ast.StarExpr).X.(*ast.Ident)
	if ok {
		sel = ident.Name
	} else {
		log.Fatalln("判断错误")
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