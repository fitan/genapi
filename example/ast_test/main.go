
package main

import (
	"bytes"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"log"

	"golang.org/x/tools/go/packages"
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

func main() {


	var fset = token.NewFileSet()
	cfg := &packages.Config{Fset: fset, Mode: mode}
	pkgs, err := packages.Load(cfg, "./types")
	if err != nil {
		log.Fatal(err)
	}

	for _, pkg := range pkgs {
		findPkg(pkg)
		//findInPackage(pkg, fset)
	}
}

func findPkg(pkg *packages.Package)  {
	for _, fileAst := range pkg.Syntax {
		ast.Inspect(fileAst, func(node ast.Node) bool {
			t, ok := node.(*ast.TypeSpec)
			if ok {
				if t.Name.Name == "In" {
					findIdent(pkg,t.Type.(*ast.StructType))
					return false
				}
			}
			return true
		})
	}

}

func findIdent(pkg *packages.Package, structType *ast.StructType)  {
	ast.Inspect(structType, func(node ast.Node) bool {
		ident,ok := node.(*ast.Ident)
		if ok {
			fmt.Println("ident Name: " + ident.Name + " ty: ")
			spew.Dump(ident.Obj)
			return false
		}
		return true
	})
}

// findInPackage finds embeddings in the package pkg.
func findInPackage(pkg *packages.Package, fset *token.FileSet) {
	for _, fileAst := range pkg.Syntax {
		ast.Inspect(fileAst, func(n ast.Node) bool {

			if structTy, ok := n.(*ast.StructType); ok {
				findInFields(structTy.Fields, n, pkg.TypesInfo, fset)
			} else if interfaceTy, ok := n.(*ast.InterfaceType); ok {
				findInFields(interfaceTy.Methods, n, pkg.TypesInfo, fset)
			}

			return true
		})
	}
}

// findInFields finds embeddings in the field list fl. The field list is taken
// from either the fields of a struct or the method list of an interface.
func findInFields(fl *ast.FieldList, n ast.Node, tinfo *types.Info, fset *token.FileSet) {
	type FieldReport struct {
		Name string
		Kind string
		Type types.Type
	}
	var reps []FieldReport

	for _, field := range fl.List {
		switch nty := tinfo.TypeOf(field.Type).Underlying().(type) {
		case *types.Named:
			fmt.Println(field.Names, "named")
		case *types.Basic:
			fmt.Println(field.Names, "basic")
		case *types.Struct:
			fmt.Println(field.Names, "struct")
			fmt.Println("field 0",nty.Field(0).Pkg().Name())
		case *types.Map:
			fmt.Println(field.Names, "Map")
			nty.Key()

		default:
			fmt.Println(field.Names, "未知")
		}
		continue

		if field.Names == nil {

			fmt.Println(field.Type)
			tv, ok := tinfo.Types[field.Type]
			if !ok {
				log.Fatal("not found", field.Type)
			}

			embName := fmt.Sprintf("%v", field.Type)

			_, hostIsStruct := n.(*ast.StructType)
			var kind string

			switch typ := tv.Type.Underlying().(type) {
			case *types.Struct:
				if hostIsStruct {
					kind = "struct (s@s)"
				} else {
					kind = "struct (s@i)"
				}
				reps = append(reps, FieldReport{embName, kind, typ})
			case *types.Interface:
				if hostIsStruct {
					kind = "interface (i@s)"
				} else {
					kind = "interface (i@i)"
				}
				reps = append(reps, FieldReport{embName, kind, typ})
			default:
			}
		}
	}

	if len(reps) > 0 {
		pos := fset.Position(n.Pos())
		fmt.Println(pos.Filename)
		fmt.Printf("Found at %v\n%v\n", fset.Position(n.Pos()), nodeString(n, fset))

		for _, report := range reps {
			fmt.Printf("--> field '%s' is embedded %s: %s\n", report.Name, report.Kind, report.Type)
		}
		fmt.Println("")
	}
}

// nodeString formats a syntax tree in the style of gofmt.
func nodeString(n ast.Node, fset *token.FileSet) string {
	var buf bytes.Buffer
	format.Node(&buf, fset, n)
	return buf.String()
}