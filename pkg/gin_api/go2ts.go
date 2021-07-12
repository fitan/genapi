package gen_apiV2

import (
	"fmt"
	"github.com/fatih/structtag"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"reflect"
	"regexp"
	"strings"
)

var Indent = "    "

func getIdent(s string) string {
	switch s {
	case "bool":
		return "boolean"
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"complex64", "complex128":
		return "number"
	}

	return s
}

func writeType(s *strings.Builder, t ast.Expr, depth int) {
	switch t := t.(type) {
	case *ast.StarExpr:
		writeType(s, t.X, depth)
		s.WriteString(" | undefined")
	case *ast.ArrayType:
		if v, ok := t.Elt.(*ast.Ident); ok && v.String() == "byte" {
			s.WriteString("string")
			break
		}
		writeType(s, t.Elt, depth)
		s.WriteString("[]")
	case *ast.StructType:
		s.WriteString("{\n")
		writeFields(s, t.Fields.List, depth+1)

		for i := 0; i < depth+1; i++ {
			s.WriteString(Indent)
		}
		s.WriteByte('}')
	case *ast.Ident:
		s.WriteString(getIdent(t.String()))
	case *ast.SelectorExpr:
		longType := fmt.Sprintf("%s.%s", t.X, t.Sel)
		switch longType {
		case "time.Time":
			s.WriteString("string")
		case "decimal.Decimal":
			s.WriteString("number")
		default:
			s.WriteString(longType)
		}
	case *ast.MapType:
		s.WriteString("{ [key: ")
		writeType(s, t.Key, depth)
		s.WriteString("]: ")
		writeType(s, t.Value, depth)
		s.WriteByte('}')
	case *ast.InterfaceType:
		s.WriteString("any")
	default:
		err := fmt.Errorf("unhandled: %s, %T", t, t)
		fmt.Println(err)
		panic(err)
	}
}

var validJSNameRegexp = regexp.MustCompile(`(?m)^[\pL_][\pL\pN_]*$`)

func validJSName(n string) bool {
	return validJSNameRegexp.MatchString(n)
}

func writeFields(s *strings.Builder, fields []*ast.Field, depth int) {
	for _, f := range fields {
		optional := false

		var fieldName string
		if len(f.Names) != 0 && f.Names[0] != nil && len(f.Names[0].Name) != 0 {
			fieldName = f.Names[0].Name
		}
		if len(fieldName) == 0 || 'A' > fieldName[0] || fieldName[0] > 'Z' {
			continue
		}

		var name string
		if f.Tag != nil {
			tags, err := structtag.Parse(f.Tag.Value[1 : len(f.Tag.Value)-1])
			if err != nil {
				panic(err)
			}

			jsonTag, err := tags.Get("json")
			if err == nil {
				name = jsonTag.Name
				if name == "-" {
					continue
				}

				optional = jsonTag.HasOption("omitempty")
			}
		}

		if len(name) == 0 {
			name = fieldName
		}

		for i := 0; i < depth+1; i++ {
			s.WriteString(Indent)
		}

		quoted := !validJSName(name)

		if quoted {
			s.WriteByte('\'')
		}
		s.WriteString(name)
		if quoted {
			s.WriteByte('\'')
		}

		switch t := f.Type.(type) {
		case *ast.StarExpr:
			optional = true
			f.Type = t.X
		}

		if optional {
			s.WriteByte('?')
		}

		s.WriteString(": ")

		writeType(s, f.Type, depth)

		s.WriteString(";\n")
	}
}

func Convert(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}

	var f ast.Node
	f, err := parser.ParseExprFrom(token.NewFileSet(), "editor.go", s, parser.SpuriousErrors)
	if err != nil {
		s = fmt.Sprintf(`package main
func main() {
	%s
}`, s)

		f, err = parser.ParseFile(token.NewFileSet(), "editor.go", s, parser.SpuriousErrors)
		if err != nil {
			panic(err)
		}
	}

	w := new(strings.Builder)
	name := "MyInterface"

	first := true

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.Ident:
			name = x.Name
		case *ast.StructType:
			if !first {
				w.WriteString("\n\n")
			}

			w.WriteString("declare interface ")
			w.WriteString(name)
			w.WriteString(" {\n")

			writeFields(w, x.Fields.List, 0)

			w.WriteByte('}')

			first = false

			// TODO: allow multiple structs
			return false
		}
		return true
	})

	return w.String()
}

func SpliceTypeV2(pkg *packages.Package, file *ast.File, node ast.Node) bool {
	replace := false
	astutil.Apply(node, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		case *ast.Field:
			fmt.Println("ent ast field")
			if t.Tag == nil {
				fmt.Println("out ast field")
				return false
			}

			if tags, ok := reflect.StructTag(t.Tag.Value[1 : len(t.Tag.Value)-1]).Lookup("json"); ok {
				for _, tag := range strings.Split(tags, ",") {
					if tag == "-" {
						fmt.Println("out ast field")
						return false
					}
				}
			}
			fmt.Println("out ast field")
			return true

		case *ast.Ident:
			//if pkg.TypesInfo.TypeOf(t) != nil {
			//	switch InfoT := pkg.TypesInfo.TypeOf(t).Underlying().(type) {
			//	case *types.Struct:
			//		fmt.Println("struct: ", t.Name, InfoT.String())
			//	case *types.Basic:
			//		switch tt := c.Parent().(type) {
			//		case *ast.Field:
			//			fmt.Println("t: ", t.Name, "parent: ", tt.Names)
			//			fmt.Println("basic: ", t.Name, InfoT.String())
			//		default:
			//			fmt.Println("basic: ", t.Name, InfoT.String())
			//
			//		}
			//	case *types.Named:
			//		fmt.Println("named: ",  t.Name, InfoT.String())
			//
			//
			//
			//
			//
			//	}
			//}
			fmt.Println("t name: ", t.Name, "t obj: ", t.Obj)
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {



					//fmt.Println("ident name: ", t.Name)
					fmt.Println("find need name: ", t.Name)
					f, findTs := FindTypeByName(pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType);ok {
						return false
					}



					file.Imports = append(file.Imports, f.Imports...)
					fmt.Println("find: ", findTs)
					fmt.Println("find ident type: ", findTs.Type, "name: ", t.Name)
					replace = true
					c.Replace(findTs.Type)
				}
			} else {
				if !JudgeBuiltInType(t.Name) {
					f, findTs := FindTypeByName(pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType);ok {
						return false
					}
					file.Imports = append(file.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			}
			//if t.Name == "Code" {
			//	fmt.Println("code obj ", t.Obj)
			//}
			//if t.Name == "AliaseInt" {
			//	fmt.Println("AliaseInt obj", t.Obj)
			//}
			//fmt.Println("enter ident: ", t.Name)
			//if t.Obj == nil {
			//	fmt.Println("obj is nil: ", t.Name)
			//	if !JudgeBuiltInType(t.Name) {
			//		fmt.Println("ident name: ", t.Name)
			//		_, findTs := FindTypeByName(pkg, t.Name)
			//		//	fmt.Println("find ident type: ", findTs.Type)
			//		c.Replace(findTs.Type)
			//	}
			//}
			//if _,ok := c.Parent().(*ast.Field);ok {
			//	return true
			//}
			//fmt.Println("ident Name ", t.Name)
			//if !JudgeBuiltInType(t.Name) {
			//	_, findTs := FindTypeByName(pkg, t.Name)
			//	fmt.Println("find ident type: ", findTs.Type)
			//	c.Replace(findTs.Type)
			//}
		//
		//
		//
		case *ast.SelectorExpr:
			defer fmt.Println("ent ast selector expr")
			fmt.Println("ent ast selector expr")
			if t.X.(*ast.Ident).Name == "time" && t.Sel.Name == "Time" {
				return false
			}
			path := FindImportPath(file.Imports, t.X.(*ast.Ident).Name)
			fmt.Println("file name: ", file.Name, "find pkg path: ", t.X.(*ast.Ident).Name)
			findPkg := pkg.Imports[path]
			fmt.Println("find pkg : ", path)
			if findPkg.Imports != nil {
				for index, importPath := range findPkg.Imports {
					pkg.Imports[index] = importPath
				}
			}
			for _, synx := range findPkg.Syntax {
				pkg.Syntax = append(pkg.Syntax, synx)
			}
			f, findTs := FindTypeByName(findPkg, t.Sel.Name)
			if _, ok := findTs.Type.(*ast.StructType);ok {
				return false
			}

			file.Imports = append(file.Imports, f.Imports...)

			replace = true
			c.Replace(findTs.Type)
			return false
		}
		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})

	//fmt.Println("splice struct :   ", Node2String(pkg.Fset, node))
	return replace
}


type ExtractStruct2Ts struct {
	PendNodes       map[string]ast.Node
	ResolveNodes    []ast.Node
	ResolveMergePkg map[string]interface{}
	Pkg             *packages.Package
	File            *ast.File
	TempType        ast.Node
	EnterType       ast.Node
}

func (e *ExtractStruct2Ts) Pend2Temp() bool {
	if len(e.PendNodes) > 0 {
		for k,v := range e.PendNodes {
			e.TempType = v
			delete(e.PendNodes, k)
			return true
		}
	}
	return false
}
func (e *ExtractStruct2Ts) Temp2ResolveNodes() {
	e.ResolveNodes = append(e.ResolveNodes, e.TempType)
}
func (e *ExtractStruct2Ts) MergePkg(pkg *packages.Package) {
	if pkg.Imports != nil {
		for index, importPath := range pkg.Imports {
			e.Pkg.Imports[index] = importPath
		}
	}
	if _,ok := e.ResolveMergePkg[pkg.PkgPath]; !ok { 
		if e.Pkg.PkgPath  != pkg.PkgPath {
			for _, synx := range pkg.Syntax {
				e.Pkg.Syntax = append(e.Pkg.Syntax, synx)
			}
		}
	}
}

func (e *ExtractStruct2Ts) AddPendNodes(key string, spec *ast.TypeSpec) bool {
	if _, ok := e.PendNodes[key]; ok {
		return false
	} else {
		e.PendNodes[key] = spec
		return true
	}
	return false
}

func (e *ExtractStruct2Ts) SpliceType() bool {
	replace := false
	astutil.Apply(e.TempType, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		case *ast.Field:
			if t.Tag == nil {
				return false
			}

			if tags, ok := reflect.StructTag(t.Tag.Value[1 : len(t.Tag.Value)-1]).Lookup("json"); ok {
				for _, tag := range strings.Split(tags, ",") {
					if tag == "-" {
						return false
					}
				}
			}
			return true

		case *ast.Ident:
			if t.Obj != nil {
				if t.Obj.Kind.String() == "type" {
					f, findTs := FindTypeByName(e.Pkg, t.Name)
					if _,ok := findTs.Type.(*ast.StructType); ok {
						if e.AddPendNodes(string(findTs.Pos()) + string(findTs.End()), findTs) {
							e.File.Imports = append(e.File.Imports, f.Imports...)
						}
						return false
					}
					e.File.Imports = append(e.File.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			}
		case *ast.SelectorExpr:
			findPkg := FindPkgBySelector(e.Pkg, e.File, t)
			e.MergePkg(findPkg)
			f, findTs := FindTypeByName(findPkg, t.Sel.Name)
			if _,ok := findTs.Type.(*ast.StructType); ok {
				if e.AddPendNodes(string(findTs.Pos()) + string(findTs.End()), findTs) {
					e.File.Imports = append(e.File.Imports, f.Imports...)
				}
				return false
			}
			e.File.Imports = append(e.File.Imports, f.Imports...)
			replace = true
			c.Replace(findTs.Type)
		}

		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
	return replace
}

func (e *ExtractStruct2Ts) Parse()  {
	e.TempType = e.EnterType
	ok := e.SpliceType()
	for ok {
		ok = e.SpliceType()
	}
	e.Temp2ResolveNodes()


	for e.Pend2Temp() {
		ok := e.SpliceType()
		for ok {
			ok = e.SpliceType()
		}
		e.Temp2ResolveNodes()
	}
}


