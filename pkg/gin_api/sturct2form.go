package gen_apiV2

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func formWriteType(name string, s *strings.Builder, t ast.Expr, opt ...string) {
	switch t := t.(type) {

	case *ast.StarExpr:
		formWriteType(name, s, t.X)

	case *ast.ArrayType:
		s.WriteString(fmt.Sprintf(`%s: {
		  type: 'array',
		  title: '%s',
		  items:`, name, name))
		formWriteType(name, s, t.Elt)
		s.WriteString(`}`)

	case *ast.StructType:
		s.WriteString(fmt.Sprintf(`{
		type: 'object',
		properties: {`))
		formWriteFields(s, t.Fields.List)
		s.WriteString(`}}`)
	case *ast.Ident:
		s.WriteString(fmt.Sprintf(`%s: {
		  type: '%s',	
		  title: '%s'	
		}`, name, getIdent(t.String()), name))
		if len(opt) == 0 {
			s.WriteString(",")
		}
	//case *ast.SelectorExpr:
	//	longType := fmt.Sprintf("%s.%s", t.X, t.Sel)
	//	switch longType {
	//	case "time.Time":
	//		s.WriteString("string")
	//	case "decimal.Decimal":
	//		s.WriteString("number")
	//	default:
	//		s.WriteString(longType)
	//	}
	//case *ast.MapType:
	//	s.WriteString("{ [key: ")
	//	writeType(s, t.Key, depth)
	//	s.WriteString("]: ")
	//	writeType(s, t.Value, depth)
	//	s.WriteByte('}')
	//case *ast.InterfaceType:
	//	s.WriteString("any")
	default:
		//err := fmt.Errorf("unhandled: %s, %T", t, t)
		//fmt.Println(err)
		//panic(err)
	}
}

func formWriteFields(s *strings.Builder, fields []*ast.Field) {
	fLen := len(fields)
	for _, f := range fields {

		for index, nameIdent := range f.Names {
			if nameIdent.IsExported() {
				if fLen-1 == index {
					formWriteType(nameIdent.Name, s, f.Type, "last")
				} else {
					formWriteType(nameIdent.Name, s, f.Type)
				}
			}
		}

		//var fieldName string
		//if len(f.Names) != 0 && f.Names[0] != nil && len(f.Names[0].Name) != 0 {
		//	fieldName = f.Names[0].Name
		//}
		//if len(fieldName) == 0 || 'A' > fieldName[0] || fieldName[0] > 'Z' {
		//	continue
		//}
		//
		//var name string
		//if f.Tag != nil {
		//	tags, err := structtag.Parse(f.Tag.Value[1 : len(f.Tag.Value)-1])
		//	if err != nil {
		//		panic(err)
		//	}
		//
		//	jsonTag, err := tags.Get("json")
		//	if err == nil {
		//		name = jsonTag.Name
		//		if name == "-" {
		//			continue
		//		}
		//
		//	}
		//
		//}
		//
		//if len(name) == 0 {
		//	name = fieldName
		//}

		//for i := 0; i < depth+1; i++ {
		//	s.WriteString(Indent)
		//}
		//
		//quoted := !validJSName(name)
		//
		//if quoted {
		//	s.WriteByte('\'')
		//}
		//s.WriteString(name)
		//if quoted {
		//	s.WriteByte('\'')
		//}
		//
		//switch t := f.Type.(type) {
		//case *ast.StarExpr:
		//	f.Type = t.X
		//}
		//
		//
		//s.WriteString(": ")
		//
		//writeType(s, f.Type, depth)
		//
		//s.WriteString("\n")
	}
}

func FormConvert(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	var f ast.Node
	f, err := parser.ParseExprFrom(token.NewFileSet(), "editor.go", s, parser.SpuriousErrors|parser.ParseComments)
	if err != nil {
		s = fmt.Sprintf(`package main
func main() {
	type %s
}`, s)

		f, err = parser.ParseFile(token.NewFileSet(), "editor.go", s, parser.SpuriousErrors)
		if err != nil {
			panic(err)
		}
	}

	w := new(strings.Builder)

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
		case *ast.StructType:
			formWriteType("", w, x)

			return false
		}
		return true
	})

	return w.String()
}
