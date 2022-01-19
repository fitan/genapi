package gen_apiV2

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"io/ioutil"
	"testing"
	"time"
)

type Test2Json struct {
	Link
	time.Time
}

type Link string

type LinkStruct struct {
	I  string
}

func TestNewExtractStruct2Ts(t *testing.T) {

	//l := Test2Json{
	//	Link: "Fdsa",
	//	Time: time.Now(),
	//}
	//b, _ := json.Marshal(l)
	//fmt.Println(string(b))
	//return

	type args struct {
		pkg  *packages.Package
		file *ast.File
		node ast.Node
	}
	_, pkg, _ := LoadPackages("./TestData")
	//err := ast.Print(pkg.Fset, pkg.Syntax)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//return
	//f, findTs := FindTypeByName(pkg, "UserResult")
	workFile, workTs := FindTypeByName(pkg, "Worker")
	ast.Inspect(workTs.Type, func(node ast.Node) bool {
		switch t := node.(type) {
		case *ast.SelectorExpr:
			fmt.Println("selector")

			if t.Sel.Obj != nil {
				fmt.Println(t.Sel.Obj)
			} else {
				fmt.Println("selector is nill")
			}

		case *ast.Ident:
			if t.Name == "Next" {
				fmt.Printf("Next %+v", t.Obj)
			}
			fmt.Printf("ident name: %s\n", t.Name)
			if t.Obj != nil {
				fmt.Println("ident obj not nil")
				if _, ok := t.Obj.Decl.(*ast.TypeSpec); ok {
					fmt.Printf("decl is type: %s\n", t.Name)
					pos := t.Obj.Pos()
					if pos.IsValid() {
						position := pkg.Fset.Position(pos)
						fmt.Printf("position: %s\n", position.String())
						for _, f := range pkg.Syntax {
							fileName := pkg.Fset.Position(f.Pos()).Filename
							if fileName == position.Filename {
								fmt.Println("文件名字相等")
								for _, d := range f.Decls {
									if gd, ok := d.(*ast.GenDecl); ok {
										for _, s :=range gd.Specs {
											if s.Pos()  == t.Obj.Pos() {
												fmt.Println("pos 相等")
											}
										}
									}
								}
							}
						}
					}
				}
			}


		}
		return true
	})
	return
	//k8sFile, k8sTs := FindTypeByName(pkg, "K8sDeploy")
	tests := []struct {
		name string
		args args
		want *ExtractStruct2Ts
	}{
		//{
		//	name: "UserResult",
		//	args: args{
		//		pkg:  pkg,
		//		file: f,
		//		node: findTs.Type,
		//	},
		//	want: nil,
		//},
		//{
		//	name: "k8s",
		//	args: args{
		//		pkg:  pkg,
		//		file: k8sFile,
		//		node: k8sTs.Type,
		//	},
		//	want: nil,
		//},
		{
			name: "k8s",
			args: args{
				pkg:  pkg,
				file: workFile,
				node: workTs.Type,
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewDepthContext(tt.args.pkg, tt.args.file, tt.args.node)
			n := DepthType(ctx)
			//fmt.Println(n)
			ts := Convert("type MyInterface " + Node2String(tt.args.pkg.Fset, n))
			ioutil.WriteFile("gostruct.ts", []byte(ts), 0777)
			//got := NewExtractStruct2Ts(tt.args.pkg, tt.args.file, tt.args.node, make(map[string]struct{}, 0))
			//got.Parse()
			//for index, v  := range got.ToTs(func(s string) string {
			//	fmt.Println(s)
			//	return s
			//}) {
			//	fmt.Println(index, v)
			//}
		})
	}
}
