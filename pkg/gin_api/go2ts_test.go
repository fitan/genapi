package gen_apiV2

import (
	"fmt"
	"go/ast"
	"strings"
	"testing"
)

func Test_writeType(t *testing.T) {
	type args struct {
		s     *strings.Builder
		t     ast.Expr
		depth int
	}
	_, pkg, _ := LoadPackages("./TestData")
	_, findType := FindTypeByName(pkg, "Test1T")
	builder := strings.Builder{}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1t",
			args: args{
				s:    &builder,
				t:     findType.Type,
				depth: 0,
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeType(tt.args.s,tt.args.t,tt.args.depth)
			fmt.Println(tt.args.s.String())
		})
	}
}
