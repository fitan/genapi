package gen_apiV2

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"strings"
)

var jsonInterface string = `
package interfacepkg


type Json interface {
	MarshalJSON() ([]byte, error)
}
`

func defineJsonInterface() *types.Interface {
	fileName := "jsonInterface.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, fileName, jsonInterface, parser.AllErrors)
	if err != nil {
		panic(err)
	}
	info := &types.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue, 0),
		Defs:       make(map[*ast.Ident]types.Object, 0),
		Uses:       make(map[*ast.Ident]types.Object, 0),
		Implicits:  make(map[ast.Node]types.Object, 0),
		Selections: make(map[*ast.SelectorExpr]*types.Selection, 0),
		Scopes:     make(map[ast.Node]*types.Scope, 0),
		InitOrder:  make([]*types.Initializer, 0, 0),
	}

	_, err = new(types.Config).Check(fileName, fset, []*ast.File{f}, info)
	if err != nil {
		log.Panic(err)
	}
	var jsonInterface *types.Interface

	ast.Inspect(f, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if i, ok := typeSpec.Type.(*ast.InterfaceType); ok {
				jsonInterface = info.TypeOf(i).(*types.Interface)
				return false
			}
		}
		return true
	})

	return jsonInterface
}

func CheckJsonImplements(ctx *DepthContext, node ast.Node) *ast.Ident {
	typeOf := ctx.Pkg.TypesInfo.TypeOf(node.(ast.Expr))
	if types.Implements(typeOf, JsonInterface) {
		return ast.NewIdent("string")
	}
	return nil
}

type DepthContext struct {
	Pkg  *packages.Package
	File *ast.File
	Node ast.Node
}

var JsonInterface *types.Interface

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	JsonInterface = defineJsonInterface()
}


func SelectorCover(expr *ast.SelectorExpr) *ast.Ident {
	if expr.X.(*ast.Ident).Name == "time" && expr.Sel.Name == "Time" {
		return ast.NewIdent("string")
	}
	return nil
}

func ToTs(pkg *packages.Package, file *ast.File, node ast.Node, format func(string) string) []string {
	tss := make([]string,0,0)
	ctx := NewDepthContext(pkg, file, node)
	ts := Convert(format(Node2String(pkg.Fset,DepthType(ctx))))
	tss = append(tss,ts)
	return tss
}

func NewDepthContext(pkg *packages.Package, file *ast.File, node ast.Node) *DepthContext {
	return &DepthContext{Pkg: pkg, File: file, Node: node}
}

func DepthType(ctx *DepthContext) ast.Node {
	defer func() {
		err := recover()
		if err != nil {
			//log.Printf("err: %v, pkgName: %v, fileName: %v, nodeString %v", err, ctx.Pkg.Name, GetFileNameByPos(ctx.Pkg.Fset, ctx.File.Pos()), Node2String(ctx.Pkg.Fset, ctx.Node))
			log.Printf("err: %v, pkgName: %v, fileName: %v", err, ctx.Pkg.Name, GetFileNameByPos(ctx.Pkg.Fset, ctx.File.Pos()))
		}
	}()
	newNode := astutil.Apply(ctx.Node, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		// 是struct 的field类型
		case *ast.Field:
			//log.Printf("field name: %v, obj: %v", t.Names[0].Name, t.Names[0].Obj.Kind.String())
			// 匿名字段 取出字段放入父struct
			if t.Names == nil {
				switch tt := t.Type.(type) {
				// 匿名selector
				case *ast.SelectorExpr:
					if s := SelectorCover(tt); s != nil {
						tmp := ast.Field{
							Doc:     t.Doc,
							Names:   []*ast.Ident{ast.NewIdent(tt.Sel.Name)},
							Type:    s,
							Tag:     t.Tag,
							Comment: t.Comment,
						}

						c.Replace(&tmp)
						return false
					}

					findPkg := FindPkgBySelector(ctx.Pkg, ctx.File, tt)
					findFile, findTs := FindTypeByName(findPkg, tt.Sel.Name)

					nextCtx := NewDepthContext(findPkg, findFile, findTs.Type)
					nextNode := DepthType(nextCtx)

					structType, ok := nextNode.(*ast.StructType)
					if ok {
						for _, f := range structType.Fields.List {
							c.InsertBefore(f)
						}
						c.Delete()
					} else {
						tmp := ast.Field{
							Doc:     t.Doc,
							Names:   []*ast.Ident{ast.NewIdent(tt.Sel.Name)},
							Type:    nextNode.(ast.Expr),
							Tag:     t.Tag,
							Comment: t.Comment,
						}
						c.Replace(&tmp)
					}
					return false
				case *ast.Ident:
					if JudgeBuiltInType(tt.Name) {
						return true
					}
					findFile, findTs := FindTypeByName(ctx.Pkg, tt.Name)

					nextCtx := NewDepthContext(ctx.Pkg, findFile, findTs.Type)
					nextNode := DepthType(nextCtx)

					structType, ok := nextNode.(*ast.StructType)
					if ok {
						for _, f := range structType.Fields.List {
							c.InsertBefore(f)
						}
						c.Delete()
					} else {
						tmp := ast.Field{
							Doc:     t.Doc,
							Names:   []*ast.Ident{ast.NewIdent(tt.Name)},
							Type:    nextNode.(ast.Expr),
							Tag:     t.Tag,
							Comment: t.Comment,
						}
						c.Replace(&tmp)
						//log.Panicf("匿名引用字段类型不为struct: %v, nextNode: %v", Node2String(ctx.Pkg.Fset, findTs), Node2String(ctx.Pkg.Fset, nextNode))
					}
					return false
				}
			}

			// 不为匿名字段

			//没有 tag 跳过
			if t.Tag == nil {
				return false
			}

			// 没有json 跳过
			if tags, ok := reflect.StructTag(t.Tag.Value[1 : len(t.Tag.Value)-1]).Lookup("json"); ok {
				for _, tag := range strings.Split(tags, ",") {
					if tag == "-" {
						return false
					}
				}
			}



		case *ast.StarExpr:
			jsonIdent := CheckJsonImplements(ctx, t)
			if jsonIdent != nil {
				c.Replace(jsonIdent)
				return false
			}

			nextCtx := NewDepthContext(ctx.Pkg,ctx.File, t.X)
			nextType := DepthType(nextCtx)
			c.Replace(nextType)
			return false

		//非匿名selector
		case *ast.SelectorExpr:
			jsonIdent := CheckJsonImplements(ctx, t)
			if jsonIdent != nil {
				c.Replace(jsonIdent)
				return false
			}

			//if n := SelectorCover(t); n != nil {
			//	c.Replace(n)
			//	return false
			//}
			findPkg := FindPkgBySelector(ctx.Pkg, ctx.File, t)
			findFile, findTs := FindTypeByName(findPkg, t.Sel.Name)
			nextCtx := NewDepthContext(findPkg, findFile, findTs.Type)
			nextType := DepthType(nextCtx)

			c.Replace(nextType)
			return false
		case *ast.Ident:
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {
					//if JudgeBuiltInType(t.Name) {
					//	return false
					//}

					jsonIdent := CheckJsonImplements(ctx, t)
					if jsonIdent != nil {
						c.Replace(jsonIdent)
						return false
					}

					log.Printf("kind: type,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj.Decl)
					findFile, findTs := FindTypeByName(ctx.Pkg, t.Name)

					nextCtx := NewDepthContext(ctx.Pkg, findFile, findTs.Type)
					nextType := DepthType(nextCtx)
					c.Replace(nextType)
					return false
				}
				if t.Obj.Kind.String() == "var" {
					//if !JudgeBuiltInType(t.Name) {
					//	//jsonIdent := CheckJsonImplements(ctx, t)
					//	//if jsonIdent != nil {
					//	//	c.Replace(jsonIdent)
					//	//	return false
					//	//}
					//
					//	findFile, findTs := FindTypeByName(ctx.Pkg, t.Name)
					//
					//	nextCtx := NewDepthContext(ctx.Pkg, findFile, findTs.Type)
					//	nextType := DepthType(nextCtx)
					//	c.Replace(nextType)
					//	return false
					//
					//}
					log.Printf("kind: var,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj)
					return false
					//jsonIdent := CheckJsonImplements(ctx, t)
					//if jsonIdent != nil {
					//	c.Replace(jsonIdent)
					//}
				}

				log.Printf("kind: 未知,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj.Decl)

			} else {
				log.Printf("kind: Nil未知,local pkg: %v, name: %v, tstring: %v", ctx.Pkg.PkgPath, t.Name, t.String())
				if !JudgeBuiltInType(t.Name) {
					jsonIdent := CheckJsonImplements(ctx, t)
					if jsonIdent != nil {
						c.Replace(jsonIdent)
						return false
					}
					findFile, findTs := FindTypeByName(ctx.Pkg, t.Name)
					nextCtx := NewDepthContext(ctx.Pkg, findFile, findTs.Type)
					nextType := DepthType(nextCtx)
					c.Replace(nextType)
					return false
				}
			}

		}

		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
	return newNode

}
