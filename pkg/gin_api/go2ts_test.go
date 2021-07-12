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
	f, findTs := FindTypeByName(pkg, "UserResult")
	tests := []struct {
		name string
		args args
		want *ExtractStruct2Ts
	}{
		{
			name: "UserResult",
			args: args{
				pkg:  pkg,
				file: f,
				node: findTs.Type,
			},
			want: nil,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewExtractStruct2Ts(tt.args.pkg, tt.args.file, tt.args.node)
			got.Parse()
			for index, v := range got.ToTs("MainIn") {
				fmt.Println(index, v)
			}
		})
	}
}
