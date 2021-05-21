package gen_apiV2

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/token"
	"log"
	"path"
	"strings"
)


type FileContext struct {
	PkgName string
	Pkg *ast.Package
	Fset *token.FileSet
	File *ast.File
	ImportMsgs map[string]ImportMsg
	Funcs []Func
}

type Func struct {
	Comments []string
	Router Router
	FuncName string
	Bind Bind
	ResOut0 string
}

type Router struct{
	Method string
	GinPath string
}

type Bind struct {
	HasUri bool
	HasQuery bool
	HasBody bool
	HasHeader bool
	QuoteBody string
	QuoteQuery string
	QuoteHeader string
	QuoteUri string
}

func NewFileContext(pkgName string, pkg *ast.Package ,fset *token.FileSet, file *ast.File) *FileContext {
	return &FileContext{PkgName: pkgName, Pkg: pkg,Fset: fset, File: file}
}

func (c *FileContext) Parse() {
	c.ParseImport()
	c.FilterFunc()
	fs := make([]Func,0,0)
	for _, fd := range c.FilterFunc() {
		f := c.ParseFunc(fd)
		fs = append(fs, f)
	}
	c.Funcs = fs
}
func (c *FileContext) ParseFunc(f *ast.FuncDecl) Func {
	fc := Func{
		FuncName: f.Name.Name,
		ResOut0: "",
	}
	inField := f.Type.Params.List[1]
	fset, inStruct := c.FindStruct(inField)
	fc.Bind = c.ParseBind(fset, inStruct)
	c.ParseComment(&fc, f.Doc.List)
	outField := f.Type.Results.List[0]
	fc.ResOut0 = Node2String(c.Fset,Node2SwagType(outField.Type, c.File.Name.Name))
	return fc
}

func (c *FileContext) ParseComment(fc *Func,ms []*ast.Comment)  {
	comments := make([]string,0,0)
	for _, m := range ms {
		fs := strings.Fields(m.Text)

		if fs[1] == GenMark {
			router, swagRouter := c.ApiMark2SwagRouter(fs)
			fc.Router = router
			comments = append(comments, swagRouter)
		}
	}
	fc.Comments = comments
}


func (c *FileContext) FindStruct(field *ast.Field) (*token.FileSet, *ast.StructType) {
	x, sel, has := IsSelector(field)
	if has {
		importMsg, ok := c.ImportMsgs[x]
		if !ok {
			log.Fatalln("not find import pkg: " + x)
		}
		_, fset, _, structType, has := FindStructByDir(importMsg.Dir, sel)
		if !has {
			log.Fatalln("not find struct: " + sel)
		}
		return fset,structType
	} else {
		_, structType, has:= FindStructByPkg(c.Pkg, sel)
		if !has {
			log.Fatalln("not find struct: " + sel)
		}
		return c.Fset, structType
	}
}


func (c *FileContext) ParseBind(fset *token.FileSet, structType *ast.StructType) Bind {
	bind := Bind{}
	for _, field := range structType.Fields.List {
		for _, ident := range field.Names {
			raw := Node2String(c.Fset,Node2SwagType(field.Type, c.File.Name.Name))
			switch ident.Name {
			case "Query":
				bind.HasQuery = true
				//bind.QuoteQuery = c.Struct2Quote(fset, field)
				bind.QuoteQuery = raw
			case "Body":
				bind.HasBody = true
				bind.QuoteBody = raw
			case "Uri":
				bind.HasUri = true
				bind.QuoteUri = raw
			case "Header":
				bind.HasHeader = true
				bind.QuoteHeader = raw
			}
		}
	}
	return bind
}

func (c *FileContext) Struct2Quote(fset *token.FileSet, field *ast.Field) string {
	_, sel, has := IsSelector(field)
	if has {
		return Node2String(fset, field.Type)
	} else {
		return c.PkgName + "." + sel
	}
}

func (c *FileContext) FilterFunc() []*ast.FuncDecl {
	fs := make([]*ast.FuncDecl, 0,0)
	ast.Inspect(c.File, func(node ast.Node) bool {
		if funcDecl, ok := node.(*ast.FuncDecl); ok {
			if c.HasApiMark(funcDecl.Doc) && c.GinFormat(funcDecl) {
				fs = append(fs, funcDecl)
			}
			return false
		}
		return true
	})
	return fs
}

// HasApiMark 注释包含@GenApi的才符合规范
func (c *FileContext) HasApiMark(doc *ast.CommentGroup) bool {
	if doc == nil {
		return false
	}
	for _, comment := range doc.List {
		fs := strings.Fields(comment.Text)
		if len(fs) < 4 {
			continue
		}
		if fs[0] == "//" && fs[1] == GenMark && len(fs[3]) > 2 {
			return true
		}
	}
	return false
}

func (c *FileContext) ApiMark2SwagRouter(fields []string) (Router, string) {
	fields[1] = "@Router"
	method := fields[3]
	ginPath := fields[2]
	ginPath = strings.ReplaceAll(ginPath, "{", ":")
	ginPath = strings.ReplaceAll(ginPath, "}", "")
	return Router{
		Method:     strings.ToUpper(method[1: len(method)-1]),
		GinPath:    ginPath,
	}, strings.Join(fields, " ")
}

// GinFormat 符合 func Name(c *gin.context, in object) (out object, err error)
func (c *FileContext) GinFormat(f *ast.FuncDecl) bool {
	if f.Type.Params.NumFields() != 2 || f.Type.Results.NumFields() != 2 {
		return false
	}
	paramGinContext := f.Type.Params.List[0]
	if selectorExpr, ok := paramGinContext.Type.(*ast.StarExpr).X.(*ast.SelectorExpr); ok {
		if selectorExpr.X.(*ast.Ident).Name == "gin" && selectorExpr.Sel.Name == "Context" {
		}
	} else {
		return false
	}

	paramIn := f.Type.Params.List[1]
	if _, ok := paramIn.Type.(*ast.StarExpr); !ok {
		log.Fatalln(fmt.Sprintf("Func %v %v type not ptr", f.Name.Name, paramIn.Names[0].Name))
		return false
	}

	paramErr := f.Type.Results.List[1]
	if ident, ok := paramErr.Type.(*ast.Ident); ok {
		if ident.Name == "error" {
			return true
		}
	}
	return false
}

type ImportMsg struct {
	Dir string
	AliseName string
	PkgName string
}

func (c *FileContext) ParseImport() {
	m := make(map[string]ImportMsg)
	ast.Inspect(c.File, func(node ast.Node) bool {
		if importSpec, ok := node.(*ast.ImportSpec); ok {
			v := importSpec.Path.Value
			p, err := build.Import(strings.ReplaceAll(v, `"`, ""), "/", build.FindOnly)
			if err != nil {
				log.Fatalln(err)
			}
			name := p.Name
			if name == "" {
				name = path.Base(strings.ReplaceAll(v, `"`, ""))
			}
			msg := ImportMsg{
				Dir:       p.Dir,
				AliseName: "",
				PkgName:   name,
			}
			if importSpec.Name == nil {
				msg.AliseName = name
				m[name] = msg
				return true
			}
			if importSpec.Name.Name == "." || importSpec.Name.Name == "_" {
				return true
			}
			msg.AliseName = importSpec.Name.Name
			m[msg.AliseName] = msg
		}
		return true
	})
	c.ImportMsgs = m
}