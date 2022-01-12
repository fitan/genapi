package gen_apiV2

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"strings"
)

type DepthContext struct {
	Pkg  *packages.Package
	File *ast.File
	Node ast.Node
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func SelectorCover(expr *ast.SelectorExpr) *ast.Ident {
	if expr.X.(*ast.Ident).Name == "time" && expr.Sel.Name == "Time" {
		return ast.NewIdent("string")
	}
	return nil
}

func NewDepthContext(pkg *packages.Package, file *ast.File, node ast.Node) *DepthContext {
	return &DepthContext{Pkg: pkg, File: file, Node: node}
}

func DepthType(ctx *DepthContext) ast.Node {
	fmt.Println(ctx)
	fmt.Println("start depth", Node2String(ctx.Pkg.Fset, ctx.Node))
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("err: %v, pkgName: %v, fileName: %v, nodeString %v", err, ctx.Pkg.Name, GetFileNameByPos(ctx.Pkg.Fset, ctx.File.Pos()), Node2String(ctx.Pkg.Fset, ctx.Node))
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
							Names:   t.Names,
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
							Names:   t.Names,
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

		//非匿名selector
		case *ast.SelectorExpr:
			if n := SelectorCover(t); n != nil {
				c.Replace(n)
				return false
			}
			findPkg := FindPkgBySelector(ctx.Pkg, ctx.File, t)
			findFile, findTs := FindTypeByName(findPkg, t.Sel.Name)
			nextCtx := NewDepthContext(findPkg, findFile, findTs.Type)
			nextType := DepthType(nextCtx)

			c.Replace(nextType)
			return false
		case *ast.Ident:
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {
					if JudgeBuiltInType(t.Name) {
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
					log.Printf("kind: var,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj.Decl)
					return false
				}
				if !JudgeBuiltInType(t.Name) {
					log.Printf("kind: !judge,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj.Decl)
					findFile, findTs := FindTypeByName(ctx.Pkg, t.Name)
					nextCtx := NewDepthContext(ctx.Pkg, findFile, findTs.Type)
					nextType := DepthType(nextCtx)
					c.Replace(nextType)
					return false
				}
				log.Printf("kind: 未知,local pkg: %v, name: %v, decl: %v", ctx.Pkg.PkgPath, t.Name, t.Obj.Decl)

			}

		}

		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
	return newNode

}
