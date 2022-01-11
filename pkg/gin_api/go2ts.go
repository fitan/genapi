package gen_apiV2

import (
	"fmt"
	"github.com/fatih/structtag"
	"go/ast"
	"go/parser"
	"go/token"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var Indent = "  "

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

//func writeType(s *strings.Builder, t ast.Expr, depth int, optionalParens bool) {
//	switch t := t.(type) {
//	case *ast.StarExpr:
//		if optionalParens {
//			s.WriteByte('(')
//		}
//		writeType(s, t.X, depth, false)
//		s.WriteString(" | undefined")
//		if optionalParens {
//			s.WriteByte(')')
//		}
//	case *ast.ArrayType:
//		if v, ok := t.Elt.(*ast.Ident); ok && v.String() == "byte" {
//			s.WriteString("string")
//			break
//		}
//		writeType(s, t.Elt, depth, true)
//		s.WriteString("[]")
//	case *ast.StructType:
//		s.WriteString("{\n")
//		writeFields(s, t.Fields.List, depth+1)
//
//		for i := 0; i < depth+1; i++ {
//			s.WriteString(Indent)
//		}
//		s.WriteByte('}')
//	case *ast.Ident:
//		s.WriteString(getIdent(t.String()))
//	case *ast.SelectorExpr:
//		longType := fmt.Sprintf("%s.%s", t.X, t.Sel)
//		switch longType {
//		case "time.Time":
//			s.WriteString("string")
//		case "decimal.Decimal":
//			s.WriteString("number")
//		default:
//			s.WriteString(longType)
//		}
//	case *ast.MapType:
//		s.WriteString("{ [key: ")
//		writeType(s, t.Key, depth, false)
//		s.WriteString("]: ")
//		writeType(s, t.Value, depth, false)
//		s.WriteByte('}')
//	case *ast.InterfaceType:
//		s.WriteString("any")
//	default:
//		err := fmt.Errorf("unhandled: %s, %T", t, t)
//		fmt.Println(err)
//		panic(err)
//	}
//}

func writeType(s *strings.Builder, t ast.Expr, depth int, opt ...string) {
	switch t := t.(type) {
	case *ast.StarExpr:
		writeType(s, t.X, depth)
		if len(opt) != 0 {
			s.WriteString(opt[0] + " | undefined")
		} else {
			s.WriteString(" | undefined")
		}
	case *ast.ArrayType:
		if v, ok := t.Elt.(*ast.Ident); ok && v.String() == "byte" {
			s.WriteString("string")
			break
		}
		writeType(s, t.Elt, depth, "[]")
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
		optional := true

		if f.Doc != nil {
			for _, m := range f.Doc.List {
				s.WriteString(Indent + m.Text)
				s.WriteString("\n")
				fmt.Println("doc: ", m.Text)
			}
		}

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

				//optional = jsonTag.HasOption("omitempty")
			}

			bindingTag, err := tags.Get("binding")
			if err == nil {
				optional = !bindingTag.HasOption("required")
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
	f, err := parser.ParseExprFrom(token.NewFileSet(), "editor.go", s, parser.SpuriousErrors|parser.ParseComments)
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

			w.WriteString("export interface ")
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

type NodeInfo struct {
	Pkg  *packages.Package
	File *ast.File
	Node ast.Node
}

type ExtractStruct2Ts struct {
	PendNodeRecord  map[string]struct{}
	PendNodes       []NodeInfo
	ResolveNodes    []NodeInfo
	ResolveMergePkg map[string]struct{}
	TempNodeInfo    NodeInfo
	EnterType       NodeInfo
}

func NewExtractStruct2Ts(pkg *packages.Package, file *ast.File, node ast.Node, pendNodeRecord map[string]struct{}) *ExtractStruct2Ts {
	e := &ExtractStruct2Ts{EnterType: NodeInfo{
		Pkg:  pkg,
		File: file,
		Node: node,
	}, PendNodeRecord: pendNodeRecord, ResolveMergePkg: map[string]struct{}{}}
	e.PendNodeRecord[e.nodeRecordKey(pkg, file, node)] = struct{}{}
	return e
}

func (e *ExtractStruct2Ts) Pend2Temp() bool {
	if len(e.PendNodes) > 0 {
		e.TempNodeInfo = e.PendNodes[0]
		e.PendNodes = e.PendNodes[1:]
		return true
	}
	return false
}

func (e *ExtractStruct2Ts) Temp2ResolveNodes() {
	e.ResolveNodes = append(e.ResolveNodes, e.TempNodeInfo)
}

func (e *ExtractStruct2Ts) MergePkg(pkg *packages.Package) {
	log.Println("mergePkg", pkg)
	log.Printf("merge imports: %v, syntax: %v", len(e.TempNodeInfo.Pkg.Imports), len(e.TempNodeInfo.Pkg.Syntax))
	if pkg.Imports != nil {
		for index, importPath := range pkg.Imports {
			e.TempNodeInfo.Pkg.Imports[index] = importPath
		}
	}
	if _, ok := e.ResolveMergePkg[pkg.PkgPath]; !ok {
		if e.TempNodeInfo.Pkg.PkgPath != pkg.PkgPath {
			for _, synx := range pkg.Syntax {
				e.TempNodeInfo.Pkg.Syntax = append(e.TempNodeInfo.Pkg.Syntax, synx)
				//e.TempNodeInfo.Pkg.Syntax = append(e.TempNodeInfo.Pkg.Syntax, synx)
			}
		}
	} else {
		e.ResolveMergePkg[pkg.PkgPath] = struct {
		}{}
	}
}

func (e *ExtractStruct2Ts) AddPendNodes(pkg *packages.Package, file *ast.File, node *ast.TypeSpec) bool {
	key := e.nodeRecordKey(pkg, file, node)

	if _, ok := e.PendNodeRecord[key]; ok {
		return false
	} else {
		e.PendNodes = append(e.PendNodes, NodeInfo{
			Pkg:  pkg,
			File: file,
			Node: node,
		})
		e.PendNodeRecord[key] = struct {
		}{}
		return true
	}
}

func (e *ExtractStruct2Ts) SpliceTypeV2() bool {
	replace := false
	newNode := astutil.Apply(e.TempNodeInfo.Node, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		case *ast.Field:

			//匿名字段
			if t.Names == nil {
				switch tt := t.Type.(type) {
				case *ast.SelectorExpr:
					log.Printf("FindPkgBySelector.pkgName: %v, filePath: %v, type: %v", e.TempNodeInfo.Pkg.Name, GetFileNameByPos(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File.Pos()), Node2String(e.TempNodeInfo.Pkg.Fset, tt))
					findPkg := FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, tt)
					log.Printf("FindTypeByName: %v", Node2String(e.TempNodeInfo.Pkg.Fset, tt))
					findFile, findTs := FindTypeByName(findPkg, tt.Sel.Name)
					//e.AddPendNodes(findPkg, findFile, findTs)
					//if e.TempNodeInfo.Pkg.PkgPath != findPkg.PkgPath {
					//	e.MergePkg(findPkg)
					//	e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, findFile.Imports...)
					//}
					_, ok := findTs.Type.(*ast.StructType)
					if ok {
						e.AddPendNodes(findPkg, findFile, findTs)
						tmpField := &ast.Field{
							Doc:     t.Doc,
							Names:   []*ast.Ident{ast.NewIdent(tt.Sel.Name)},
							Type:    ast.NewIdent(tt.Sel.Name),
							Tag:     t.Tag,
							Comment: t.Comment,
						}
						//t.Type = ast.NewIdent(tt.Sel.Name)

						c.Replace(tmpField)
						//for _, f := range structType.Fields.List {
						//	c.InsertBefore(f)
						//}
						//c.Delete()
						//replace = true
						return false
						//return false
					}
					ct := t
					ident := ast.NewIdent(findTs.Name.Name)
					ident.Obj = ast.NewObj(ast.Var, findTs.Name.Name)
					ct.Names = []*ast.Ident{ident}
					ct.Type = findTs.Type
					ct.Doc = findTs.Doc
					c.Replace(ct)
					replace = true
					return false

				case *ast.Ident:
					findFile, findTs := FindTypeByName(e.TempNodeInfo.Pkg, tt.Name)
					_, ok := findTs.Type.(*ast.StructType)
					if ok {
						e.AddPendNodes(e.TempNodeInfo.Pkg, findFile, findTs)
						return false
						//for _, f := range structType.Fields.List {
						//	c.InsertBefore(f)
						//}
						//c.Delete()
						//replace = true
						//return false
					}
					ct := t
					ident := ast.NewIdent(findTs.Name.Name)
					ident.Obj = ast.NewObj(ast.Var, findTs.Name.Name)
					ct.Names = []*ast.Ident{ident}
					ct.Type = findTs.Type
					ct.Doc = findTs.Doc
					c.Replace(ct)
					replace = true
					return false
					//case *ast.SelectorExpr:
					//	_, findTs := findtype
					//
				}
			}

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
					f, findTs := FindTypeByName(e.TempNodeInfo.Pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType); ok {
						e.AddPendNodes(e.TempNodeInfo.Pkg, f, findTs)
						return false
					}
					//e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			} else {
				if !JudgeBuiltInType(t.Name) {
					log.Printf("findtypeByName. pkgName: %v, typeName: %v.", e.TempNodeInfo.Pkg.Name, t.Name)
					f, findTs := FindTypeByName(e.TempNodeInfo.Pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType); ok {
						e.AddPendNodes(e.TempNodeInfo.Pkg, f, findTs)
						return false
					}
					//e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			}
		case *ast.SelectorExpr:
			if n := e.selectorCover(t); n != nil {
				c.Replace(n)
				return false
			}

			fmt.Println(t.End()+t.Pos(), Node2String(e.TempNodeInfo.Pkg.Fset, t))
			//if t.End() + t.Pos() == 6 {
			//	fmt.Println(t.End()+t.Pos(), Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			//}

			var findPkg *packages.Package

			if e.TempNodeInfo.Pkg.Name != t.X.(*ast.Ident).Name {
				findPkg = FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, t)
				if findPkg == nil {
					log.Fatalln("file: ", e.TempNodeInfo.File.Name.Name, " type: ", Node2String(e.TempNodeInfo.Pkg.Fset, t))
				}
				//e.MergePkg(findPkg)
			} else {
				findPkg = e.TempNodeInfo.Pkg
			}

			//fmt.Println("this pkg name: ", e.TempNodeInfo.Pkg.Name)
			//findPkg := FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, t)

			//fmt.Println("find need merge pkg: ", e.TempNodeInfo.Pkg.Name, t.X.(*ast.Ident).Name,t.Sel.Name)
			//fmt.Println(Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			//e.MergePkg(findPkg)
			FindFile, findTs := FindTypeByName(findPkg, t.Sel.Name)
			if _, ok := findTs.Type.(*ast.StructType); ok {
				e.AddPendNodes(findPkg, FindFile, findTs)
				tmpNode := ast.NewIdent(t.Sel.Name)
				c.Replace(tmpNode)
				return false
			}
			replace = true
			//fmt.Println("replace selector: ", Node2String(e.TempNodeInfo.Pkg.Fset, findTs.Type))
			c.Replace(findTs.Type)
			//fmt.Println("replace after: ", Node2String(e.TempNodeInfo.Pkg.Fset, c.Node()))
			//fmt.Println("echo file: ", Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			return false
		}

		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
	fmt.Println(Node2String(e.TempNodeInfo.Pkg.Fset, newNode))
	//fmt.Println("new node: ", Node2String(e.TempNodeInfo.Pkg.Fset,newNode))
	e.TempNodeInfo.Node = newNode
	return replace
}

func (e *ExtractStruct2Ts) SpliceType() bool {
	replace := false
	newNode := astutil.Apply(e.TempNodeInfo.Node, func(c *astutil.Cursor) bool {
		switch t := c.Node().(type) {
		case *ast.Field:

			if t.Names == nil {
				switch tt := t.Type.(type) {
				case *ast.SelectorExpr:
					log.Printf("FindPkgBySelector.pkgName: %v, filePath: %v, type: %v", e.TempNodeInfo.Pkg.Name, GetFileNameByPos(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File.Pos()), Node2String(e.TempNodeInfo.Pkg.Fset, tt))
					findPkg := FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, tt)
					log.Printf("FindTypeByName: %v", Node2String(e.TempNodeInfo.Pkg.Fset, tt))
					findFile, findTs := FindTypeByName(findPkg, tt.Sel.Name)
					if e.TempNodeInfo.Pkg.PkgPath != findPkg.PkgPath {
						e.MergePkg(findPkg)
						e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, findFile.Imports...)
					}
					structType, ok := findTs.Type.(*ast.StructType)
					if ok {
						for _, f := range structType.Fields.List {
							c.InsertBefore(f)
						}
						c.Delete()
						replace = true
						return false
					}
					ct := t
					ident := ast.NewIdent(findTs.Name.Name)
					ident.Obj = ast.NewObj(ast.Var, findTs.Name.Name)
					ct.Names = []*ast.Ident{ident}
					ct.Type = findTs.Type
					ct.Doc = findTs.Doc
					c.Replace(ct)
					replace = true
					return false

				case *ast.Ident:
					_, findTs := FindTypeByName(e.TempNodeInfo.Pkg, tt.Name)
					structType, ok := findTs.Type.(*ast.StructType)
					if ok {
						for _, f := range structType.Fields.List {
							c.InsertBefore(f)
						}
						c.Delete()
						replace = true
						return false
					}
					ct := t
					ident := ast.NewIdent(findTs.Name.Name)
					ident.Obj = ast.NewObj(ast.Var, findTs.Name.Name)
					ct.Names = []*ast.Ident{ident}
					ct.Type = findTs.Type
					ct.Doc = findTs.Doc
					c.Replace(ct)
					replace = true
					return false
					//case *ast.SelectorExpr:
					//	_, findTs := findtype
					//
				}
			}

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
					f, findTs := FindTypeByName(e.TempNodeInfo.Pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType); ok {
						e.AddPendNodes(e.TempNodeInfo.Pkg, f, findTs)
						return false
					}
					e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			} else {
				if !JudgeBuiltInType(t.Name) {
					log.Printf("findtypeByName. pkgName: %v, typeName: %v.", e.TempNodeInfo.Pkg.Name, t.Name)
					f, findTs := FindTypeByName(e.TempNodeInfo.Pkg, t.Name)
					if _, ok := findTs.Type.(*ast.StructType); ok {
						e.AddPendNodes(e.TempNodeInfo.Pkg, f, findTs)
						return false
					}
					e.TempNodeInfo.File.Imports = append(e.TempNodeInfo.File.Imports, f.Imports...)
					replace = true
					c.Replace(findTs.Type)
				}
			}
		case *ast.SelectorExpr:
			if n := e.selectorCover(t); n != nil {
				c.Replace(n)
				return false
			}

			fmt.Println(t.End()+t.Pos(), Node2String(e.TempNodeInfo.Pkg.Fset, t))
			//if t.End() + t.Pos() == 6 {
			//	fmt.Println(t.End()+t.Pos(), Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			//}

			var findPkg *packages.Package

			if e.TempNodeInfo.Pkg.Name != t.X.(*ast.Ident).Name {
				findPkg = FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, t)
				if findPkg == nil {
					log.Fatalln("file: ", e.TempNodeInfo.File.Name.Name, " type: ", Node2String(e.TempNodeInfo.Pkg.Fset, t))
				}
				e.MergePkg(findPkg)
			} else {
				findPkg = e.TempNodeInfo.Pkg
			}

			//fmt.Println("this pkg name: ", e.TempNodeInfo.Pkg.Name)
			//findPkg := FindPkgBySelector(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, t)

			//fmt.Println("find need merge pkg: ", e.TempNodeInfo.Pkg.Name, t.X.(*ast.Ident).Name,t.Sel.Name)
			//fmt.Println(Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			//e.MergePkg(findPkg)
			FindFile, findTs := FindTypeByName(findPkg, t.Sel.Name)
			if _, ok := findTs.Type.(*ast.StructType); ok {
				e.AddPendNodes(findPkg, FindFile, findTs)
				tmpNode := ast.NewIdent(t.Sel.Name)
				c.Replace(tmpNode)
				return false
			}
			replace = true
			//fmt.Println("replace selector: ", Node2String(e.TempNodeInfo.Pkg.Fset, findTs.Type))
			c.Replace(findTs.Type)
			//fmt.Println("replace after: ", Node2String(e.TempNodeInfo.Pkg.Fset, c.Node()))
			//fmt.Println("echo file: ", Node2String(e.TempNodeInfo.Pkg.Fset, e.TempNodeInfo.File))
			return false
		}

		return true
	}, func(c *astutil.Cursor) bool {
		return true
	})
	//fmt.Println("new node: ", Node2String(e.TempNodeInfo.Pkg.Fset,newNode))
	e.TempNodeInfo.Node = newNode
	return replace
}
func (e *ExtractStruct2Ts) nodeRecordKey(pkg *packages.Package, file *ast.File, node ast.Node) string {
	return strconv.Itoa(int(node.Pos())) + strconv.Itoa(int(node.End()))
}

func (e *ExtractStruct2Ts) selectorCover(expr *ast.SelectorExpr) *ast.Ident {
	if expr.X.(*ast.Ident).Name == "time" && expr.Sel.Name == "Time" {
		return ast.NewIdent("string")
	}
	return nil
}

func (e *ExtractStruct2Ts) Parse() {

	//e.TempNodeInfo = e.EnterType
	//ok := e.SpliceType()
	//for ok {
	//	ok = e.SpliceType()
	//}
	//e.EnterType = e.TempNodeInfo
	//
	//for e.Pend2Temp() {
	//	ok := e.SpliceType()
	//	for ok {
	//		ok = e.SpliceType()
	//	}
	//	e.Temp2ResolveNodes()
	//}

	e.TempNodeInfo = e.EnterType
	ctx := NewDepthContext(e.TempNodeInfo.Pkg, e.TempNodeInfo.File, e.TempNodeInfo.Node)
	node := DepthType(ctx)
	fmt.Println("depthTyp: ", node)
	return
	e.SpliceTypeV2()

	e.EnterType = e.TempNodeInfo

	for e.Pend2Temp() {
		e.SpliceTypeV2()
		e.Temp2ResolveNodes()
	}
}

func (e *ExtractStruct2Ts) toTsFormat(s string) string {
	return fmt.Sprintf("type Obj %s", s)
}

func (e *ExtractStruct2Ts) ToTs(format func(string) string) []string {
	tss := make([]string, 0, 0)
	tss = append(tss, Convert(format(Node2String(e.EnterType.Pkg.Fset, e.EnterType.Node))))
	for _, node := range e.ResolveNodes {
		tss = append(tss, Convert("type "+Node2String(node.Pkg.Fset, node.Node)))
	}
	return tss
}
