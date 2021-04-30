package pkg

import (
	"fmt"
	"go/ast"
	"go/token"
	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/imports"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type GoMod struct {
	Module Module
}

type Module struct {
	Path    string
	Version string
}

func GetModuleName() string {
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		log.Fatalln(err.Error())
	}

	modName := modfile.ModulePath(goModBytes)
	fmt.Fprintf(os.Stdout, "modName=%+v\n", modName)

	return modName
}

type (
	file struct {
		path    string
		content []byte
	}
	assets struct {
		dirs  []string
		files []file
	}
)

func (a assets) write() error {
	for _, dir := range a.dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}
	for _, file := range a.files {
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			return fmt.Errorf("write file %q: %w", file.path, err)
		}
	}
	return nil
}

// formatGo runs "goimports" on all assets.
func (a assets) formatGo() error {
	for _, file := range a.files {
		path := file.path
		src, err := imports.Process(path, file.content, nil)
		if err != nil {
			return fmt.Errorf("formatGo file %s: %v", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %v", path, err)
		}
	}
	return nil
}


func Res2SwagModel(fset *token.FileSet, node ast.Node, selectName string) ast.Node {
	t := res2SwagModel2(node, selectName)
	t = res2SwagModel1(t)
	return t
}

func res2SwagModel1(node ast.Node) ast.Node {
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

func res2SwagModel2(node ast.Node, selectName string) ast.Node {
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
