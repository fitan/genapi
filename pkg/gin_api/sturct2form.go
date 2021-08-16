package gen_apiV2

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

const NGForm = "ngform"

type ngTag struct {
	tagS string
	nameField string
	typeField string
	titleField string
	formatField string
}

func (n *ngTag) parse() {

	jsonFormTag := reflect.StructTag(n.tagS).Get("json")
	for _, v := range strings.Split(jsonFormTag, ",") {
		n.nameField = v
		break
	}
	ngFormTag := reflect.StructTag(n.tagS).Get(NGForm)
	if ngFormTag != "" {
		l := strings.Split(ngFormTag, ",")
		for _, v := range l {
			field := strings.Split(v, "=")
			if len(field) >= 2 {
				switch field[0] {
				case "type":
					n.typeField = field[1]
				case "title":
					n.titleField = field[1]
					if n.titleField == "" {
						n.titleField = n.nameField
					}
				case "format":
					n.formatField = field[1]
				}
			}
		}
	}

}

func NGTAG(s string) ngTag {
	t := ngTag{tagS: s}
	t.parse()
	return t
}

func formWriteType(tag ngTag, s *strings.Builder, t ast.Expr,  opt ...string) {
	switch t := t.(type) {

	case *ast.StarExpr:
		formWriteType(tag, s, t.X)

	case *ast.ArrayType:
		switch tt := t.Elt.(type) {
		case *ast.StructType:
			s.WriteString(fmt.Sprintf(`%s: {
			type: 'array',
			title: '%s',
			items:`, tag.nameField, tag.titleField))
			formWriteType(tag, s, t.Elt)
			s.WriteString(`}`)
		case *ast.Ident:
			formWriteType(tag, s, tt)
		}

	case *ast.StructType:
		s.WriteString(fmt.Sprintf(`{
		type: 'object',
		properties: {`))
		formWriteFields(s, t.Fields.List)
		s.WriteString(`}}`)
	case *ast.Ident:
		ty := tag.typeField
		title := tag.titleField
		format := tag.formatField
		name := tag.nameField

		if ty == "" {
			ty = getIdent(t.String())
		}


		s.WriteString(fmt.Sprintf(`%s: {
		  type: '%s',	
		  title: '%s'`, name, ty, title))
		if format != "" {
			s.WriteString(fmt.Sprintf(`,format: '%s'`, format))
		}
		s.WriteString(`}`)
		if len(opt) == 0 {
			s.WriteString(",")
		}

	default:
		//err := fmt.Errorf("unhandled: %s, %T", t, t)
		//fmt.Println(err)
		//panic(err)
	}
}

func formWriteFields(s *strings.Builder, fields []*ast.Field) {
	fLen := len(fields)
	for _, f := range fields {
		var ngTag ngTag
		if f.Tag == nil {
			ngTag = NGTAG("")
		} else {
			ngTag = NGTAG(f.Tag.Value[1 : len(f.Tag.Value)-1])
		}

		for index, nameIdent := range f.Names {
			if nameIdent.IsExported() {
				if fLen-1 == index {
					formWriteType(ngTag, s, f.Type, "last")
				} else {
					formWriteType(ngTag, s, f.Type)
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
			formWriteType(NGTAG(""), w, x)

			return false
		}
		return true
	})

	return w.String()
}
