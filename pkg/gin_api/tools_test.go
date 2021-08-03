package gen_apiV2

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestSpliceStruct(t *testing.T) {
	type args struct {
		pkgs *packages.Package
		file *ast.File
		st   ast.Node
		objName string
	}
	_, pkgs, _ := LoadPackages("./TestData")
	f, _, structType := FindStructTypeByName(pkgs, "UserResult")

	//f1, t1 := FindTypeByName(pkgs, "Test1T")
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{
			pkgs: pkgs,
			file: f,
			st:   structType,
			objName: "FuIn",
		}},
		//{
		//	name: "test2",
		//	args: args{
		//		pkgs: pkgs,
		//		file: f1,
		//		st:   t1,
		//	},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Struct2Ts(tt.args.pkgs, tt.args.file, tt.args.st, tt.args.objName))
			//spew.Dump(tt.args.st)
		})
	}
}

func TestWarpResult2Ts(t *testing.T) {
	type args struct {
		pkg  *packages.Package
		file *ast.File
		node ast.Node
		objName string
	}
	_, pkg, _ := LoadPackages("./TestData")
	f, findType := FindTypeByName(pkg, "Test1T")
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "type",
			args: args{
				pkg:  pkg,
				file: f,
				node: findType.Type,
				objName: "FnOut",
			},
			want: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println(WarpResult2Ts(tt.args.pkg, tt.args.file, tt.args.node, tt.args.objName))
			//if got := WarpResult2Ts(tt.args.pkg, tt.args.file, tt.args.node); got != tt.want {
			//	t.Errorf("WarpResult2Ts() = %v, want %v", got, tt.want)
			//}
		})
	}
}