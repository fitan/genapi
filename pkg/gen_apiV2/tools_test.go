package gen_apiV2

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestSpliceStruct(t *testing.T) {
	type args struct {
		pkgs *packages.Package
		file *ast.File
		st   *ast.StructType
	}
	_, pkgs, _ := LoadPackages("./TestData")
	f, _, structType := FindStructTypeByName(pkgs, "User")
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{
			pkgs: pkgs,
			file: f,
			st:   structType,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Log(Struct2Ts(tt.args.pkgs, tt.args.file, tt.args.st))
			//spew.Dump(tt.args.st)
		})
	}
}
