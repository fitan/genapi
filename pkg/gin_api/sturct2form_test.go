package gen_apiV2

import (
	"fmt"
	"testing"
)

func TestFormConvert(t *testing.T) {

	_, pkg, _ := LoadPackages("./TestData")
	_, findTs := FindTypeByName(pkg, "Form")
	s := Node2String(pkg.Fset, findTs)
	fmt.Println(s)
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Form",
			args: args{
				s: s,
			},
			want: "",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FormConvert(tt.args.s)
			fmt.Println(got)
		})
	}
}
