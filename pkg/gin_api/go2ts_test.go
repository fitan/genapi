package gen_apiV2

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestNewExtractStruct2Ts(t *testing.T) {
	type args struct {
		pkg  *packages.Package
		file *ast.File
		node ast.Node
	}
	_, pkg, _ := LoadPackages("./TestData")
	//f, findTs := FindTypeByName(pkg, "UserResult")
	k8sFile, k8sTs := FindTypeByName(pkg, "K8sDeploy")
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
		{
			name: "k8s",
			args: args{
				pkg:  pkg,
				file: k8sFile,
				node: k8sTs.Type,
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := NewDepthContext(tt.args.pkg, tt.args.file, tt.args.node)
			n := DepthType(ctx)
			ts := Convert("type MyInterface " + Node2String(tt.args.pkg.Fset, n))
			fmt.Println(ts)
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
