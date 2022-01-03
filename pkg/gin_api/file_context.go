package gen_apiV2

import (
	"fmt"
	"github.com/fitan/genapi/pkg/gin_api/plugins"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"log"
	"net/http"
	_ "net/http/pprof"
	"path"
	"strings"
)

type FileContext struct {
	PkgName              string
	Pkg                  *packages.Package
	File                 *ast.File
	CasbinInterfaceType1 *types.Interface
	CasbinInterfaceType2 *types.Interface
	//ImportMsgs map[string]ImportMsg
	Funcs []Func
}

type ParseOption struct {
	ParseTs bool
}

func NewFileContext(pkgName string, pkg *packages.Package, file *ast.File) *FileContext {
	return &FileContext{PkgName: pkgName, Pkg: pkg, File: file}
}

func (c *FileContext) Parse(option ParseOption) {
	//c.ImportMsgs = ParseImport(c.File)
	//c.ParseCasbinPluginer()
	fs := make([]Func, 0, 0)
	for _, fd := range c.FilterFunc() {
		f := c.ParseFunc(fd)
		fs = append(fs, f)
	}
	c.Funcs = fs
	if option.ParseTs {
		c.Func2Ts()
	}
}

func (c *FileContext) Func2Ts() {
	go func() {
		if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
			panic(err)
		}
	}()

	pendNodeRecord := make(map[string]struct{}, 0)
	for index, _ := range c.Funcs {
		inField := c.Funcs[index].Fd.Type.Params.List[1]
		_, _, _, inStruct := FindStructByExpr(c.Pkg, c.File, inField.Type.(*ast.StarExpr).X)
		inTs := NewExtractStruct2Ts(c.Pkg, c.File, inStruct, pendNodeRecord)
		inTs.Parse()
		c.Funcs[index].ParamIn1Ts = inTs.ToTs(func(s string) string {
			return fmt.Sprintf("type %sIn %s", c.Funcs[index].Fd.Name.Name, s)
		})
		outField := c.Funcs[index].Fd.Type.Results.List[0]
		outTs := NewExtractStruct2Ts(c.Pkg, c.File, outField.Type, pendNodeRecord)
		outTs.Parse()
		c.Funcs[index].ResOut0Ts = outTs.ToTs(func(s string) string {
			resutlStr := "type %sOut struct {\nCode int `json:\"code\"`\nData %s `json:\"data\"`\nErr  string `json:\"err\"`}"
			return fmt.Sprintf(resutlStr, c.Funcs[index].Fd.Name.Name, s)
		})
	}
}

func (c *FileContext) ParseFunc(f *ast.FuncDecl) Func {
	fc := Func{
		Fd:       f,
		PkgName:  c.PkgName,
		FuncName: f.Name.Name,
		ResOut0:  "",
		Plugins: Plugins{
			Point: []plugins.PointTemplate{},
		},
	}
	inField := f.Type.Params.List[1]
	outField := f.Type.Results.List[0]

	inPkg, inFile, _, inStruct := FindStructByExpr(c.Pkg, c.File, inField.Type.(*ast.StarExpr).X)
	//_, inStruct := c.FindStruct(inField)
	fc.Bind = c.ParseBind(inPkg, inFile, fc.FuncName, inStruct)
	c.ParseComment(&fc, f.Doc.List, inField, outField)
	fc.ParamIn1 = Node2String(c.Pkg.Fset, Node2SwagType(copyAST(inField.Type), c.File.Name.Name))
	fc.ResOut0 = Node2String(c.Pkg.Fset, Node2SwagType(copyAST(outField.Type), c.File.Name.Name))
	//fmt.Println("enter func %s in", fc.FuncName)
	//inTs := NewExtractStruct2Ts(c.Pkg, c.File, inStruct)
	//inTs.Parse()
	//fc.ParamIn1Ts = inTs.ToTs(func(s string) string {
	//	return fmt.Sprintf("type %sIn %s", f.Name.Name, s)
	//})
	//fmt.Println("enter func %s out", fc.FuncName)
	//outTs := NewExtractStruct2Ts(c.Pkg, c.File, outField.Type)
	//outTs.Parse()
	//fc.ResOut0Ts = outTs.ToTs(func(s string) string {
	//	resutlStr := "type %sOut struct {\nCode int `json:\"code\"`\nData %s `json:\"data\"`\nErr  string `json:\"err\"`}"
	//	return fmt.Sprintf(resutlStr, f.Name.Name, s)
	//})
	return fc
}

//func (c *FileContext) ParseCasbin(fc *Func, inField *ast.Field) {
//	if fc.Plugins.Casbin.Has {
//		casbinKeyser,ok := plugins.FindPluginInterfaceMap[plugins.CasbinKeyserName]
//		if !ok {
//			log.Fatalln("not found casbinKeyser")
//		}
//		casbinListKeyser,ok := plugins.FindPluginInterfaceMap[plugins.CasbinListKeyserName]
//		if !ok {
//			log.Fatalln("not found casbinListKeyser")
//		}
//
//		if types.Implements(c.Pkg.TypesInfo.TypeOf(inField.Type), casbinKeyser) {
//			fc.Plugins.Casbin.ImportPath = plugins.FuncTemplates[plugins.CasbinKeyserName].ImportPath
//			fc.Plugins.Casbin.Raw = fmt.Sprintf(plugins.FuncTemplates[plugins.CasbinKeyserName].Template,fc.Plugins.Casbin.CasbinMark)
//		}
//
//		if types.Implements(c.Pkg.TypesInfo.TypeOf(inField.Type), casbinListKeyser) {
//			fc.Plugins.Casbin.ImportPath = plugins.FuncTemplates[plugins.CasbinListKeyserName].ImportPath
//			fc.Plugins.Casbin.Raw = fmt.Sprintf(plugins.FuncTemplates[plugins.CasbinListKeyserName].Template,fc.Plugins.Casbin.CasbinMark)
//		}
//	}
//}

func (c *FileContext) ParseComment(fc *Func, ms []*ast.Comment, inField *ast.Field, outField *ast.Field) {
	comments := make([]string, 0, 0)
	for _, m := range ms {
		fs := strings.Fields(m.Text)
		if len(fs) < 2 {
			continue
		}
		switch fs[1] {
		case GenMark:
			param := MatchPathParam(fs[2])
			fc.Bind.Uri.Param = param
			router, swagRouter := c.ApiMark2SwagRouter(fs)
			fc.Router = router
			comments = append(comments, swagRouter)
		case CasbinMark:
			temp := plugins.GetCasbinPluginTemplate(fs, c.Pkg.TypesInfo.TypeOf(inField.Type), c.Pkg.TypesInfo.TypeOf(outField.Type))
			fc.Plugins.Point = append(fc.Plugins.Point, temp)
		case plugins.CallBackMark:
			fc.Plugins.CallBack = plugins.GetCallBackTemplate(fs, c.Pkg.TypesInfo.TypeOf(inField.Type), c.Pkg.TypesInfo.TypeOf(outField.Type))
		default:
			comments = append(comments, m.Text)
		}
	}
	fc.Comments = comments
}

func (c *FileContext) ParseBind(inPkg *packages.Package, inFile *ast.File, funcName string, structType *ast.StructType) Bind {
	bind := Bind{}
	for _, field := range structType.Fields.List {
		for _, ident := range field.Names {
			//if _,ok := c.Pkg.TypesInfo.ObjectOf(ident).Type().Underlying().(*types.Struct);!ok {
			//	continue
			//}
			var raw string
			var quoteType QuoteType

			//st, hasStructType := field.Type.(*ast.StructType)
			//if hasStructType {
			//	quoteType = StructType
			//	raw = Node2String(c.Pkg.Fset, st)
			//} else {
			//	quoteType = IdentType
			//	fmt.Println("node2string: ", Node2String(c.Pkg.Fset, field.Type))
			//	raw = Node2String(c.Pkg.Fset, Node2SwagType(field.Type, c.File.Name.Name))
			//}
			raw = Node2String(c.Pkg.Fset, Node2SwagType(field.Type, inFile.Name.Name))
			quoteType = StructType
			switch ident.Name {
			case "Query":
				bind.Query.Has = true
				bind.Query.SwagStructName = "Swag" + funcName + "Query"
				bind.Query.QuoteType = quoteType
				bind.Query.SwagRaw = raw
				bind.Query.Comment = strings.ReplaceAll(field.Doc.Text(), "\n", "\\n")
				bind.Query.SwagObj = bind.Query.SwagStructName
				//if hasStructType {
				//	bind.Query.SwagObj = bind.Query.SwagStructName
				//} else {
				//	bind.Query.SwagObj = bind.Query.SwagRaw
				//}
			case "Body":
				bind.Body.Has = true
				bind.Body.QuoteType = quoteType
				bind.Body.SwagStructName = "Swag" + funcName + "Body"
				bind.Body.SwagRaw = raw
				bind.Body.Comment = strings.ReplaceAll(field.Doc.Text(), "\n", "\\n")
				bind.Body.SwagObj = bind.Body.SwagStructName
				//if hasStructType {
				//	bind.Body.SwagObj = bind.Body.SwagStructName
				//} else {
				//	bind.Body.SwagObj = bind.Body.SwagRaw
				//}
			case "Uri":
				bind.Uri.Has = true
				log.Printf("funcName: %s, file: %s, field: %v", funcName, c.File.Name.Name)
				bind.Uri.TagMsgs = FindTagAndCommentByField(inPkg, inFile, field, "uri")
			case "Header":
				bind.Header.Has = true
				bind.Header.TagMsgs = FindTagAndCommentByField(inPkg, inFile, field, "header")
			case "CtxKey":
				bind.CtxKey.Has = true

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
	fs := make([]*ast.FuncDecl, 0, 0)
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
	GenMarkPath := fields[2]
	TsPath := fields[2]
	TsPath = strings.ReplaceAll(TsPath, "{", "${")
	routerGroupKey := ""
	if len(fields) >= 5 {
		routerGroupKey = fields[4]
		fields[2] = path.Join("/"+routerGroupKey, fields[2])
	}
	return Router{
		Method:         strings.ToUpper(method[1:2]) + method[2:len(method)-1],
		GenMarkPath:    GenMarkPath,
		GinPath:        ginPath,
		TsPath:         TsPath,
		RouterGroupKey: routerGroupKey,
	}, strings.Join(fields[:4], " ")
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
