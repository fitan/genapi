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
	"path"
	"strings"

	"golang.org/x/tools/go/packages"
)

func MatchPath(pkg *packages.Package, dir string) *packages.Package {
	path, _ := path.Split(dir)
	for k, v := range pkg.Imports {
		if strings.HasSuffix(path, k) {
			return v
		}
	}
	return nil
}

type TagMsg struct {
	TagValue string
	Comment  string
}

func FindTags(pkg *packages.Package, structType *ast.StructType, tagName string) {
	//tagMsgs := make([]TagMsg, 0,0)
	ast.Inspect(structType.Fields, func(node ast.Node) bool {
		//fd,ok := node.(*ast.Field)
		//if ok {
		//	tagTool := reflect.StructTag(fd.Tag.Value[1 : len(fd.Tag.Value)-1])
		//	value, ok := tagTool.Lookup(tagName)
		//	if ok {
		//		msg := TagMsg{
		//			TagValue: value,
		//			Comment:  fd.Doc.Text(),
		//		}
		//		tagMsgs = append(tagMsgs, msg)
		//	}
		//}
		_, ok := node.(*ast.BasicLit)
		if ok {
			return false
		}
		switch node.(type) {
		case *ast.Field:
			FindTagByType(pkg, node.(*ast.Field).Type)
			//e,ok := node.(*ast.Field).Names[0].Obj.Type.(*ast.Expr)
			//if ok {
			//	fmt.Println("期待 obj type ", e)
			//}
		}

		//expr, ok := node.(ast.Expr)
		//if ok {
		//	switch tp := pkg.TypesInfo.TypeOf(expr).Underlying().(type) {
		//	case *types.Struct:
		//		fmt.Println("node: ", pkg.Fset.Position(expr.Pos()))
		//		fmt.Println("struct: ",tp.String())
		//		return false
		//		//switch node.(type) {
		//		//case *ast.StructType:
		//		//case *ast.SelectorExpr:
		//		//case *ast.Ident:
		//		//
		//		//
		//		//
		//		//
		//		//}
		//	}
		//}
		return true
	})

}
func FindTagByType(pkg *packages.Package, ty ast.Node) {
	ast.Inspect(ty, func(node ast.Node) bool {
		switch t := node.(type) {
		case *ast.StructType:
			return false
		default:
			e, ok := node.(ast.Expr)
			if ok {
				ts, ok := pkg.TypesInfo.TypeOf(e).Underlying().(*types.Struct)
				if ok {
					fmt.Println("struct", t, ts.String())
				}
				_, seOk := t.(*ast.SelectorExpr)
				if seOk {
					return false
				}

			}
		}

		return true
	})
}

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

		for _, file := range pkg.Syntax {
			ast.Inspect(file, func(node ast.Node) bool {
				ts, ok := node.(*ast.TypeSpec)
				if ok {
					if ts.Name.Name == "In" {
						FindTags(pkg, ts.Type.(*ast.StructType), "")
						return false
					}
				}
				return true
			})
		}

		//fmt.Println(pkg.Imports)
		//findPkg(pkg)
		//fmt.Println(pkg.Imports)
		//for k,_ := range pkg.Imports {
		//	fmt.Println(k)
		//}
		//continue
		//findInPackage(pkg, fset)
	}
}

func findPkg(pkg *packages.Package) {
	for _, fileAst := range pkg.Syntax {
		ast.Inspect(fileAst, func(node ast.Node) bool {
			t, ok := node.(*ast.TypeSpec)
			if ok {

				if t.Name.Name == "In" {
					//fmt.Println(pkg.TypesInfo.TypeOf(t.Type).String())
					ty := pkg.TypesInfo.TypeOf(t.Type).Underlying().(*types.Struct)
					FindTypesStruct(pkg, ty)
					return false
					findIdent(pkg, t.Type.(*ast.StructType))
					return false
				}
			}
			return true
		})
	}

}

func FindTypesStruct(pkg *packages.Package, p types.Type) {
	st, ok := p.Underlying().(*types.Struct)
	if ok {
		n := st.NumFields()
		for i := 0; i < n; i++ {
			pkg.Fset.Position(st.Field(i).Pos())
			fmt.Println(pkg.Fset.Position(st.Field(i).Pos()).String())
			FindTypesStruct(pkg, st.Field(i).Type())
			fmt.Println(st.Tag(i))
		}
	}
}

func findIdent(pkg *packages.Package, structType *ast.StructType) {
	ast.Inspect(structType, func(node ast.Node) bool {
		node.(*ast.SelectorExpr).Sel.IsExported()
		ident, ok := node.(*ast.Ident)
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
			fmt.Println(fset.Position(field.Pos()))
			f := fset.File(nty.Field(0).Pos())
			fmt.Println(f.Name())
			fmt.Println(field.Names, "struct")
			spew.Dump(nty.Field(0))
			fmt.Println("field 0", nty.Field(0).Pkg().Path())
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
