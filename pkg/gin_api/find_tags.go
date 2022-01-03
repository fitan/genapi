package gen_apiV2

import (
	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"path"
	"reflect"
	"strings"
)

type TagMsg struct {
	TagValue string
	Comment  string
}

func FindTagAndCommentByStruct(pkg *packages.Package, file *ast.File, structType *ast.StructType, tagName string) []TagMsg {
	tagMsgs := make([]TagMsg, 0, 0)
	ast.Inspect(structType.Fields, func(node ast.Node) bool {
		fd, ok := node.(*ast.Field)
		if ok {
			if fd.Tag != nil {
				tagTool := reflect.StructTag(fd.Tag.Value[1 : len(fd.Tag.Value)-1])
				value, ok := tagTool.Lookup(tagName)
				if ok {
					msg := TagMsg{
						TagValue: value,
						Comment:  strings.ReplaceAll(fd.Doc.Text(), "\n", "\\n"),
					}
					tagMsgs = append(tagMsgs, msg)
				}
			}
		}

		if _, ok := node.(*ast.BasicLit); ok {
			return false
		}
		switch nodeType := node.(type) {
		case *ast.Field:
			tagMsgs = append(tagMsgs, FindTagByType(pkg, file, nodeType.Type, tagName)...)
		}
		return true
	})
	return tagMsgs
}
func TrimImport(s string) string {
	s = strings.TrimSuffix(s, `"`)
	s = strings.TrimPrefix(s, `"`)
	return s
}
func FindImportPath(importSpecs []*ast.ImportSpec, target string) string {
	for _, importSpec := range importSpecs {
		if importSpec.Name != nil {
			//log.Println("import name: ", importSpec.Name.Name, "import path: ", importSpec.Path.Value)
			if importSpec.Name.Name == target {
				return TrimImport(importSpec.Path.Value)
			}
		} else {
			//log.Println("import name: ", nil, "import path: ", importSpec.Path.Value)
			if target == path.Base(TrimImport(importSpec.Path.Value)) {
				return TrimImport(importSpec.Path.Value)
			}
		}
	}
	spew.Dump(importSpecs)
	log.Fatalln("not find import path: ", target)
	return ""
}

func FindTagByType(pkg *packages.Package, file *ast.File, ty ast.Node, tagName string) []TagMsg {
	tagMsgs := make([]TagMsg, 0, 0)
	ast.Inspect(ty, func(node ast.Node) bool {
		switch t := node.(type) {
		case *ast.StructType:
			return false
		default:
			e, ok := node.(ast.Expr)
			if ok {
				_, ok := pkg.TypesInfo.TypeOf(e).Underlying().(*types.Struct)
				if ok {
					switch structType := t.(type) {
					// remote pkg
					case *ast.SelectorExpr:
						log.Printf("find import path. path: %v, pkgName: %v, file: %v, typeName: %v", pkg.PkgPath, pkg.Name, GetFileNameByPos(pkg.Fset, file.Pos()), Node2String(pkg.Fset, t))
						importPath := FindImportPath(file.Imports, structType.X.(*ast.Ident).Name)
						remotePkg := pkg.Imports[importPath]
						remoteFile, _, st := FindStructTypeByName(remotePkg, structType.Sel.Name)
						tagMsgs = append(tagMsgs, FindTagAndCommentByStruct(remotePkg, remoteFile, st, tagName)...)
						return false
					// local pkg
					case *ast.Ident:
						localFile, _, st := FindStructTypeByName(pkg, structType.Name)
						tagMsgs = append(tagMsgs, FindTagAndCommentByStruct(pkg, localFile, st, tagName)...)
					}
				}
			}
		}
		return true
	})
	return tagMsgs
}

func FindStructTypeByName(pkg *packages.Package, structName string) (*ast.File, *ast.TypeSpec, *ast.StructType) {
	f, t := FindTypeByName(pkg, structName)
	st, ok := t.Type.(*ast.StructType)
	if ok {
		return f, t, st
	}
	log.Fatal("node found " + structName)
	return nil, nil, nil
	//var f *ast.File
	//var t *ast.TypeSpec
	//var st *ast.StructType
	//for _, file := range pkg.Syntax {
	//	has := false
	//	ast.Inspect(file, func(node ast.Node) bool {
	//		ts, ok := node.(*ast.TypeSpec)
	//		if ok {
	//			if ts.Name.Name == structName {
	//				has = true
	//				f = file
	//				t = ts
	//				st = ts.Type.(*ast.StructType)
	//				return false
	//			}
	//		}
	//		if has {
	//			return false
	//		}
	//		return true
	//	})
	//}
	//return f, t,st
}

func FindTypeByName(pkg *packages.Package, TypeName string) (*ast.File, *ast.TypeSpec) {
	log.Printf("findtypeByName pkgPath: %v, typeName: %v", pkg.PkgPath, TypeName)
	var f *ast.File
	var t *ast.TypeSpec
	has := false
	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(node ast.Node) bool {
			ts, ok := node.(*ast.TypeSpec)
			if ok {
				//if TypeName == "Time" {
				//	log.Printf("Typespec Name: %v", ts.Name.Name)
				//}
				if ts.Name.Name == TypeName {
					has = true
					f = file
					t = ts
					return false
				}
			}
			//if has {
			//	return false
			//}
			return true
		})
		if has {
			return f, t
		}
	}
	if f == nil || t == nil {
		log.Panicln("not find type by name. pkg: " + pkg.Name + " typeName: " + TypeName)
	}
	return f, t
}
