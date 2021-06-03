package gen_apiV2

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path"
	"reflect"
	"strings"
)

type TagMsg struct {
	TagValue string
	Comment  string
}

func FindTagAndComment(pkg *packages.Package, file *ast.File, structType *ast.StructType, tagName string) []TagMsg {
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
						Comment:  strings.ReplaceAll(fd.Doc.Text(), "\n", `\\n`),
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
			if importSpec.Name.Name == target {
				return TrimImport(importSpec.Path.Value)
			}
		} else {
			if target == path.Base(TrimImport(importSpec.Path.Value)) {
				return TrimImport(importSpec.Path.Value)
			}
		}
	}
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
				fmt.Println(e)
				_, ok := pkg.TypesInfo.TypeOf(e).Underlying().(*types.Struct)
				if ok {
					switch structType := t.(type) {
					// remote pkg
					case *ast.SelectorExpr:
						importPath := FindImportPath(file.Imports, structType.X.(*ast.Ident).Name)
						remotePkg := pkg.Imports[importPath]
						remoteFile, st := FindStructTypeByName(remotePkg, structType.Sel.Name)
						tagMsgs = append(tagMsgs, FindTagAndComment(remotePkg, remoteFile, st, tagName)...)
						return false
					// local pkg
					case *ast.Ident:
						localFile, st := FindStructTypeByName(pkg, structType.Name)
						tagMsgs = append(tagMsgs, FindTagAndComment(pkg, localFile, st, tagName)...)
					}
				}
			}
		}
		return true
	})
	return tagMsgs
}

func FindStructTypeByName(pkg *packages.Package, structName string) (*ast.File, *ast.StructType) {
	var f *ast.File
	var st *ast.StructType
	for _, file := range pkg.Syntax {
		has := false
		ast.Inspect(file, func(node ast.Node) bool {
			ts, ok := node.(*ast.TypeSpec)
			if ok {
				if ts.Name.Name == structName {
					has = true
					f = file
					st = ts.Type.(*ast.StructType)
					return false
				}
			}
			if has {
				return false
			}
			return true
		})
	}
	return f, st
}
