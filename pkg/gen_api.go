package pkg

import (
	"bytes"
	_ "embed"
	"entgo.io/ent/entc/gen"
	"github.com/fitan/genapi/pkg/gen_apiV2"
	"github.com/fitan/genapi/public"
	"github.com/marcinwyszynski/directory_tree"
	"log"
	"path"
	"path/filepath"
	"text/template"
)

//go:embed internal/templateV2/gen_api.tmpl
var gen_api_tmpl string

//go:embed gen_apiV2/template/handler.tmpl
var gen_api_tmplV2 string

//go:embed gen_apiV2/template/register.tmpl
var register_tmplV2 string

//go:embed internal/templateV2/register.tmpl
var register_tmpl string

//go:embed internal/templateV2/pkg_name.tmpl
var pkg_name_tmpl string

//const GenMark string = "@GenApi"
//
//func ParseFuncApi(src string, dest string) ParseContext {
//	fset, pkgs := LoadPkgs(src)
//	context := ParseContext{
//		LocalModuleName: GetModuleName(),
//		ParsePath:       src,
//		Fset:            fset,
//		Pkgs:            pkgs,
//	}
//	context.Enter()
//	GenApi(context.ApiMap, dest)
//	return context
//}
//
//type InHas struct {
//	HasBody   bool
//	BodyMsg   FieldMsg
//	HasUri    bool
//	UriMsg    FieldMsg
//	HasHeader bool
//	HeaderMsg FieldMsg
//	HasQuery  bool
//	QueryMsg  FieldMsg
//}
//
//type FuncRouter struct {
//	SwagRouter string
//	Method     string
//	GinPath    string
//}
//
//func ApiRouterToGinRouter(fields []string) FuncRouter {
//	fields[1] = "@Router"
//	method := fields[3]
//	swagPath := fields[2]
//	swagPath = strings.ReplaceAll(swagPath, "{", ":")
//	swagPath = strings.ReplaceAll(swagPath, "}", "")
//
//	return FuncRouter{
//		SwagRouter: strings.Join(fields, " "),
//		Method:     strings.ToUpper(method[1 : len(method)-1]),
//		GinPath:    swagPath,
//	}
//}
//
//type ApiMsg struct {
//	SrcPkgName string
//	FuncRouter FuncRouter
//	Doc        *ast.CommentGroup
//	In         struct {
//		ObjectMsg ObjectMsg
//		InHas     InHas
//	}
//	Out ObjectMsg
//}
//
//type ObjectMsg struct {
//	PkgName string
//	RawName string
//	// Selector 中的Sel
//	SelectorSel string
//	// Selector 中的X
//	SelectorX string
//	// true 则x sel都有，false只有 sel
//	IsSelector bool
//	// isselector为true 则寻找引用包路径
//	ObjectImportMsg ObjectImportMsg
//}
//type ParseContext struct {
//	LocalModuleName string
//	ParsePath       string
//	Fset            *token.FileSet
//	// pakName fileName funcName
//	ApiMap map[string]map[string]map[string]ApiMsg
//	Pkgs   map[string]*ast.Package
//}
//
//func (c *ParseContext) Enter() {
//
//	for pkgName, pkg := range c.Pkgs {
//		apiMap := make(map[string]map[string]map[string]ApiMsg, 0)
//		apiMap[pkgName] = make(map[string]map[string]ApiMsg)
//		for filePath, file := range pkg.Files {
//			apiMap[pkgName][filePath] = make(map[string]ApiMsg, 0)
//			filterFuncs := c.FileEnter(file)
//			parseImport := ParseImport(file, c.Fset)
//			for _, f := range filterFuncs {
//				inObjectMsg := c.FindInObjectMsg(f.Type.Params.List[1], file, filePath, parseImport)
//				outObjectMsg := c.FindOutObjectMsg(f.Type.Results.List[0], file, filePath, parseImport)
//				outObjectMsg.RawName = strings.ReplaceAll(outObjectMsg.RawName, "*", "")
//				_, funcRouter := c.FuncHasSwaggerRouter(f.Doc)
//				apiMsg := ApiMsg{
//					SrcPkgName: pkgName,
//					FuncRouter: funcRouter,
//					Doc:        f.Doc,
//					In: struct {
//						ObjectMsg ObjectMsg
//						InHas     InHas
//					}{
//						ObjectMsg: inObjectMsg,
//					},
//					Out: outObjectMsg,
//				}
//				apiMsg.In.InHas = c.FindInHas(apiMsg)
//
//				apiMap[pkgName][filePath][f.Name.Name] = apiMsg
//			}
//		}
//		c.ApiMap = apiMap
//	}
//}
//
//func (c *ParseContext) FileEnter(f *ast.File) []*ast.FuncDecl {
//	filterFunc := make([]*ast.FuncDecl, 0)
//	ast.Inspect(f, func(node ast.Node) bool {
//		if funcDecl, ok := node.(*ast.FuncDecl); ok {
//			hasRouter, _ := c.FuncHasSwaggerRouter(funcDecl.Doc)
//			hasConform := c.ConformFormat(funcDecl)
//			log.Println(fmt.Sprintf("func %v: hasRouter %v, hasConform %v", funcDecl.Name.Name, hasRouter, hasConform))
//			if hasRouter && hasConform {
//				filterFunc = append(filterFunc, funcDecl)
//			}
//			return false
//		}
//		return true
//	})
//	return filterFunc
//}
//
//func (c *ParseContext) FuncHasSwaggerRouter(doc *ast.CommentGroup) (bool, FuncRouter) {
//	if doc == nil {
//		return false, FuncRouter{}
//	}
//	for _, comment := range doc.List {
//		fields := strings.Fields(comment.Text)
//		if len(fields) < 4 {
//			continue
//		}
//		if fields[0] == "//" && fields[1] == GenMark && len(fields[3]) > 2 {
//			return true, ApiRouterToGinRouter(fields)
//		}
//	}
//	return false, FuncRouter{}
//}
//
//// 符合 func Name(c *gin.context, in object) (out object, err error)
//func (c *ParseContext) ConformFormat(f *ast.FuncDecl) bool {
//	if f.Type.Params.NumFields() != 2 || f.Type.Results.NumFields() != 2 {
//		return false
//	}
//	paramGinContext := f.Type.Params.List[0]
//	if selectorExpr, ok := paramGinContext.Type.(*ast.StarExpr).X.(*ast.SelectorExpr); ok {
//		if selectorExpr.X.(*ast.Ident).Name == "gin" && selectorExpr.Sel.Name == "Context" {
//		}
//	} else {
//		return false
//	}
//
//	paramIn := f.Type.Params.List[1]
//	if _, ok := paramIn.Type.(*ast.StarExpr); !ok {
//		log.Fatalln(fmt.Sprintf("Func %v %v type not ptr", f.Name.Name, paramIn.Names[0].Name))
//		return false
//	}
//
//	paramErr := f.Type.Results.List[1]
//	if ident, ok := paramErr.Type.(*ast.Ident); ok {
//		if ident.Name == "error" {
//			return true
//		}
//	}
//	return false
//}
//func (c *ParseContext) FindOutObjectMsg(field *ast.Field, file *ast.File, filePath string, msgMap map[string]ObjectImportMsg) ObjectMsg {
//	msg := ObjectMsg{
//		PkgName: file.Name.Name,
//		//RawName:     NodeString(c.Fset, field.Type),
//		RawName:     "",
//		SelectorSel: "",
//		SelectorX:   "",
//		IsSelector:  false,
//	}
//
//	has, x, sel := FieldIsSelector(field)
//	if has {
//		msg.IsSelector = true
//		msg.SelectorX = x
//		msg.SelectorSel = sel
//		if fileImportMsg, ok := msgMap[x]; !ok {
//			log.Fatalf("没有找到x %v \n", x)
//		} else {
//			msg.ObjectImportMsg = fileImportMsg
//		}
//		msg.RawName = NodeString(c.Fset, Res2SwagModel(c.Fset, field.Type, file.Name.Name))
//		return msg
//	}
//	//fmt.Println(NodeString())
//	//msg.SelectorSel = field.Type.(*ast.StarExpr).X.(*ast.Ident).Name
//	msg.ObjectImportMsg = ObjectImportMsg{RawImport: path.Join(c.LocalModuleName, c.ParsePath)}
//	msg.RawName = NodeString(c.Fset, Res2SwagModel(c.Fset, field.Type, file.Name.Name))
//	return msg
//}
//
//func (c *ParseContext) FindInObjectMsg(field *ast.Field, file *ast.File, filePath string, msgMap map[string]ObjectImportMsg) ObjectMsg {
//	msg := ObjectMsg{
//		PkgName: file.Name.Name,
//		//RawName:     NodeString(c.Fset, field.Type),
//		RawName:     "",
//		SelectorSel: "",
//		SelectorX:   "",
//		IsSelector:  false,
//	}
//
//	has, x, sel := FieldIsSelector(field)
//	if has {
//		msg.IsSelector = true
//		msg.SelectorX = x
//		msg.SelectorSel = sel
//		if fileImportMsg, ok := msgMap[x]; !ok {
//			log.Fatalf("没有找到x %v \n", x)
//		} else {
//			msg.ObjectImportMsg = fileImportMsg
//		}
//		msg.RawName = NodeString(c.Fset, Res2SwagModel(c.Fset, field.Type, file.Name.Name))
//		return msg
//	}
//	//fmt.Println(NodeString())
//	msg.SelectorSel = field.Type.(*ast.StarExpr).X.(*ast.Ident).Name
//	msg.ObjectImportMsg = ObjectImportMsg{RawImport: path.Join(c.LocalModuleName, c.ParsePath)}
//	msg.RawName = NodeString(c.Fset, Res2SwagModel(c.Fset, field.Type, file.Name.Name))
//	return msg
//}
//
//func FieldIsSelector(field *ast.Field) (bool, string, string) {
//	var has bool
//	var x string
//	var sel string
//	ast.Inspect(field, func(node ast.Node) bool {
//		if selectorExpr, ok := node.(*ast.SelectorExpr); ok {
//			has = true
//			x = selectorExpr.X.(*ast.Ident).Name
//			sel = selectorExpr.Sel.Name
//			return false
//		}
//		return true
//	})
//	return has, x, sel
//}
//
//func (c *ParseContext) LocalPackagePath(file *ast.File) {
//
//}
//
//type ObjectImportMsg struct {
//	RawImport string
//	PathDir   string
//	AliseName string
//	PkgName   string
//}
//
//func ParseImport(file *ast.File, fset *token.FileSet) map[string]ObjectImportMsg {
//	fileImportMsgs := make(map[string]ObjectImportMsg, 0)
//	ast.Inspect(file, func(node ast.Node) bool {
//		if importSpec, ok := node.(*ast.ImportSpec); ok {
//			dir, name := FindPkg(importSpec.Path.Value)
//			msg := ObjectImportMsg{PathDir: dir, PkgName: name, RawImport: NodeString(fset, node)}
//			if importSpec.Name == nil {
//				msg.AliseName = name
//				fileImportMsgs[name] = msg
//				return true
//			}
//			if importSpec.Name.Name == "." || importSpec.Name.Name == "_" {
//				return true
//			}
//			msg.AliseName = importSpec.Name.Name
//			fileImportMsgs[importSpec.Name.Name] = msg
//			return true
//		}
//		return true
//	})
//	return fileImportMsgs
//}
//func NodeString(fset *token.FileSet, node interface{}) string {
//	var buf bytes.Buffer
//	err := printer.Fprint(&buf, fset, node)
//	if err != nil {
//		spew.Dump(node)
//		log.Fatalln(err.Error())
//	}
//	return buf.String()
//}
//
//func FindPkg(pkgPath string) (string, string) {
//	p, err := build.Import(strings.ReplaceAll(pkgPath, `"`, ""), "/", build.FindOnly)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	if p.Name == "" {
//		return p.Dir, path.Base(strings.ReplaceAll(pkgPath, `"`, ""))
//	}
//	return p.Dir, p.Name
//}
//
//func LoadPkgs(path string) (*token.FileSet, map[string]*ast.Package) {
//	fset := token.NewFileSet()
//	fmt.Println(path)
//	pkgs, err := parser.ParseDir(fset, path, func(info os.FileInfo) bool {
//		return true
//	}, parser.ParseComments)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	return fset, pkgs
//}
//
//func (c *ParseContext) FindInHas(api ApiMsg) InHas {
//	if api.In.ObjectMsg.IsSelector {
//		fset, pkgs := LoadPkgs(api.In.ObjectMsg.ObjectImportMsg.PathDir)
//		fileAst, structType, has := FindStructByPkg(pkgs[api.In.ObjectMsg.SelectorX], api.In.ObjectMsg.SelectorSel)
//		if !has {
//			log.Fatalln(fmt.Sprintf("not find struct %v by pkg %v", api.In.ObjectMsg.SelectorSel, api.In.ObjectMsg.SelectorX))
//		}
//		parseImport := ParseImport(fileAst, fset)
//		inHas := FindInByStructType(structType, fset, pkgs, parseImport, api.In.ObjectMsg.ObjectImportMsg.RawImport)
//		return inHas
//	} else {
//		_, fileAst, structType, has := FindStructByPkgs(c.Pkgs, api.In.ObjectMsg.SelectorSel)
//		if !has {
//			log.Fatalln(fmt.Sprintf("not find struct %v by pkg %v", api.In.ObjectMsg.SelectorSel, api.In.ObjectMsg.PkgName))
//		}
//		parseImport := ParseImport(fileAst, c.Fset)
//		inHas := FindInByStructType(structType, c.Fset, c.Pkgs, parseImport, path.Join(c.LocalModuleName, path.Dir(c.ParsePath)))
//		return inHas
//
//	}
//
//}
//
//type FieldTag struct {
//	FieldName string
//	TagName   string
//}
//
//func FindStructByPkgs(pkgs map[string]*ast.Package, structName string) (string, *ast.File, *ast.StructType, bool) {
//	for pkgName, pkg := range pkgs {
//		fileType, structType, has := FindStructByPkg(pkg, structName)
//		if has {
//			return pkgName, fileType, structType, has
//		}
//	}
//	return "", nil, nil, false
//}
//
//func FindStructByPkg(pkg *ast.Package, structName string) (*ast.File, *ast.StructType, bool) {
//	var structType *ast.StructType
//	var fileAst *ast.File
//
//	for _, file := range pkg.Files {
//		ast.Inspect(file, func(node ast.Node) bool {
//			if typeSpec, ok := node.(*ast.TypeSpec); ok {
//				if typeSpec.Name.Name == structName {
//					if st, ok := typeSpec.Type.(*ast.StructType); ok {
//						structType = st
//						fileAst = file
//						return false
//					}
//				}
//			}
//			if structType != nil {
//				return false
//			}
//			return true
//		})
//	}
//	//if structType == nil {
//	//	log.Fatalln(fmt.Sprintf("pkg %v中, 没有找到%v struct", pkgName, structName))
//	//}
//	if structType == nil {
//		return fileAst, structType, false
//	}
//	return fileAst, structType, true
//}
//
//type ObjectSrc struct {
//	Fset    *token.FileSet
//	Pkgs    map[string]*ast.Package
//	File    *ast.File
//	Imports map[string]ObjectImportMsg
//}
//
//type StructTypeTools struct {
//	ObjectSrc
//	StructType *ast.StructType
//}
//
//func (s *StructTypeTools) Parse(tagName string) []FieldTag {
//	fieldTags := make([]FieldTag, 0, 0)
//	for _, field := range s.StructType.Fields.List {
//		if field.Names == nil {
//			switch field.Type.(type) {
//			case *ast.SelectorExpr:
//				selectorExpr := field.Type.(*ast.SelectorExpr)
//				x := selectorExpr.X.(*ast.Ident).Name
//				sel := selectorExpr.Sel.Name
//				sturctTypeTool := s.FindSturctType(s.Imports[x].PathDir, sel)
//				fieldTags = append(fieldTags, sturctTypeTool.Parse(tagName)...)
//			case *ast.Ident:
//				ident := field.Type.(*ast.Ident)
//				sel := ident.Name
//				_, file, structType, has := FindStructByPkgs(s.Pkgs, sel)
//				if !has {
//					log.Fatalln(fmt.Sprintf("not find struct name %v", sel))
//				}
//				sturctTypeTool := StructTypeTools{
//					ObjectSrc: ObjectSrc{
//						Fset:    s.Fset,
//						Pkgs:    s.Pkgs,
//						File:    file,
//						Imports: ParseImport(file, s.Fset),
//					},
//					StructType: structType,
//				}
//				fieldTags = append(fieldTags, sturctTypeTool.Parse(tagName)...)
//			default:
//				log.Fatalln(fmt.Sprintf("未知类型%v", NodeString(s.Fset, field)))
//			}
//			continue
//		}
//
//		if field.Tag == nil {
//			continue
//		}
//		tagTool := reflect.StructTag(field.Tag.Value[1 : len(field.Tag.Value)-1])
//		value, ok := tagTool.Lookup(tagName)
//		if ok {
//			fieldTags = append(fieldTags, FieldTag{
//				FieldName: field.Names[0].Name,
//				TagName:   value,
//			})
//		}
//	}
//	return fieldTags
//}
//
//func (s *StructTypeTools) FindSturctType(path string, structName string) StructTypeTools {
//	fset, pkgs := LoadPkgs(path)
//	_, file, structType, has := FindStructByPkgs(pkgs, structName)
//	if !has {
//		log.Fatalln(fmt.Sprintf("not find struct name %v", structName))
//	}
//	return StructTypeTools{
//		ObjectSrc: ObjectSrc{
//			Fset:    fset,
//			Pkgs:    pkgs,
//			File:    file,
//			Imports: ParseImport(file, fset),
//		},
//		StructType: structType,
//	}
//
//}
//
//func FindInByStructType(structType *ast.StructType, fset *token.FileSet, pkgs map[string]*ast.Package, importMsgs map[string]ObjectImportMsg, byFindImportPath string) InHas {
//	inHas := InHas{}
//	for _, field := range structType.Fields.List {
//		for _, ident := range field.Names {
//			switch ident.Name {
//			case "Query":
//				inHas.HasQuery = true
//				inHas.QueryMsg = FieldTypeIsStruct(field, fset, pkgs, importMsgs, byFindImportPath, "form")
//			case "Body":
//				inHas.HasBody = true
//				inHas.BodyMsg = FieldTypeIsStruct(field, fset, pkgs, importMsgs, byFindImportPath, "")
//			case "Uri":
//				inHas.HasUri = true
//				inHas.UriMsg = FieldTypeIsStruct(field, fset, pkgs, importMsgs, byFindImportPath, "uri")
//			case "Header":
//				inHas.HasHeader = true
//				inHas.HeaderMsg = FieldTypeIsStruct(field, fset, pkgs, importMsgs, byFindImportPath, "header")
//			}
//		}
//
//	}
//	return inHas
//}
//
//type FieldMsg struct {
//	IsSelector bool
//	PkgName    string
//	Raw        string
//	FieldTags  []FieldTag
//	ImportPath string
//}
//
//func FieldTypeIsStruct(field *ast.Field, fset *token.FileSet, pkgs map[string]*ast.Package, importMsg map[string]ObjectImportMsg, byFindImportPath string, tagMark string) FieldMsg {
//	fieldMsg := FieldMsg{Raw: NodeString(fset, field.Type)}
//	switch field.Type.(type) {
//	case *ast.SelectorExpr:
//		selectorExpr, _ := field.Type.(*ast.SelectorExpr)
//		x := selectorExpr.X.(*ast.Ident).Name
//		//sel := selectorExpr.Sel.Name
//		fieldMsg.IsSelector = true
//		fieldMsg.ImportPath = importMsg[x].RawImport
//		targetFset, targetPkgs := LoadPkgs(importMsg[x].PathDir)
//		_, targetFile, structType, has := FindStructByPkgs(targetPkgs, selectorExpr.Sel.Name)
//		if !has {
//			log.Fatalln(fmt.Sprintf("import pkg not found struct %v", x))
//		}
//		structTypeTools := StructTypeTools{
//			ObjectSrc: ObjectSrc{
//				Fset:    targetFset,
//				Pkgs:    targetPkgs,
//				File:    targetFile,
//				Imports: ParseImport(targetFile, targetFset),
//			},
//			StructType: structType,
//		}
//		fieldMsg.FieldTags = structTypeTools.Parse(tagMark)
//		return fieldMsg
//	case *ast.Ident:
//		ident, _ := field.Type.(*ast.Ident)
//		name := ident.Name
//		pkgName, targetFile, targetStructType, has := FindStructByPkgs(pkgs, name)
//		if !has {
//			log.Fatalln(fmt.Sprintf("local pkg not found struct %v", ident.Name))
//		}
//		structTypeTools := StructTypeTools{
//			ObjectSrc: ObjectSrc{
//				Fset:    fset,
//				Pkgs:    pkgs,
//				File:    targetFile,
//				Imports: ParseImport(targetFile, fset),
//			},
//			StructType: targetStructType,
//		}
//		fieldMsg.FieldTags = structTypeTools.Parse(tagMark)
//		fieldMsg.PkgName = pkgName
//		fieldMsg.ImportPath = byFindImportPath
//		return fieldMsg
//	}
//
//	log.Fatalln("field not structType")
//	return fieldMsg
//}

//func GenApi(apiMap map[string]map[string]map[string]ApiMsg, dest string) {
//	parse, err := template.New("gen_api").Parse(pkg_name_tmpl)
//	if err != nil {
//		log.Panicln(err.Error())
//	}
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	assets := assets{
//		dirs: []string{
//			filepath.Join(dest),
//		},
//	}
//
//	for _, fileMap := range apiMap {
//		for fileName, funcMap := range fileMap {
//			tpl, err := parse.Parse(gen_api_tmpl)
//			if err != nil {
//				log.Fatalln(err.Error())
//			}
//			b := bytes.NewBuffer(nil)
//			err = tpl.Execute(b, struct {
//				PkgName string
//				FuncMap map[string]ApiMsg
//			}{
//				PkgName: path.Base(dest),
//				FuncMap: funcMap,
//			})
//
//			if err != nil {
//				log.Fatalln(err.Error())
//			}
//			assets.files = append(assets.files, file{
//				path:    filepath.Join(dest, path.Base(fileName)),
//				content: b.Bytes(),
//			})
//		}
//	}
//
//	tpl, err := parse.New("register").Parse(register_tmpl)
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	b := bytes.NewBuffer(nil)
//	err = tpl.Execute(b, struct {
//		PkgName string
//		ApiMap  map[string]map[string]map[string]ApiMsg
//	}{
//		PkgName: path.Base(dest),
//		ApiMap:  apiMap,
//	})
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//	assets.files = append(assets.files, file{
//		path:    filepath.Join(dest, path.Base("register.go")),
//		content: b.Bytes(),
//	})
//
//	if err := assets.write(); err != nil {
//		log.Fatalln(err.Error())
//	}
//
//	err = assets.formatGo()
//	if err != nil {
//		log.Fatalln(err.Error())
//	}
//}

func GenApiV2(apiMap map[string]*gen_apiV2.FileContext, ReginsterMap map[string][]gen_apiV2.Func, baseConf public.BaseConf, dest string) {
	parse, err := template.New("gen_api").Funcs(gen.Funcs).Funcs(FM).Parse(pkg_name_tmpl)
	if err != nil {
		log.Panicln(err.Error())
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	assets := assets{
		dirs: []string{
			filepath.Join(dest),
		},
	}

	for fileName, fileContext := range apiMap {
		tpl, err := parse.Parse(gen_api_tmplV2)
		if err != nil {
			log.Fatalln(err.Error())
		}
		b := bytes.NewBuffer(nil)
		err = tpl.Execute(b, struct {
			PkgName string
			Funcs   []gen_apiV2.Func
			BaseConf public.BaseConf

		}{
			PkgName: path.Base(dest),
			Funcs:   fileContext.Funcs,
			BaseConf: baseConf,
		})

		if err != nil {
			log.Fatalln(err.Error())
		}
		assets.files = append(assets.files, file{
			path:    filepath.Join(dest, path.Base(fileName)),
			content: b.Bytes(),
		})
	}

	tpl, err := parse.New("register").Parse(register_tmplV2)
	if err != nil {
		log.Fatalln(err.Error())
	}
	b := bytes.NewBuffer(nil)
	err = tpl.Execute(b, struct {
		PkgName      string
		ApiMap       map[string]*gen_apiV2.FileContext
		ReginsterMap map[string][]gen_apiV2.Func
		BaseConf public.BaseConf
	}{
		PkgName:      path.Base(dest),
		ApiMap:       apiMap,
		ReginsterMap: ReginsterMap,
		BaseConf: baseConf,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	assets.files = append(assets.files, file{
		path:    filepath.Join(dest, path.Base("register.go")),
		content: b.Bytes(),
	})

	if err := assets.write(); err != nil {
		log.Fatalln(err.Error())
	}

	err = assets.formatGo()
	if err != nil {
		log.Fatalln(err.Error())
	}
}


func DepthGen(src, dir string)  {
	tree, err := directory_tree.NewTree(src)
	if err != nil {
		log.Panicln(err)
	}

	depthGen(tree, dir)
}

func depthGen(tree *directory_tree.Node, Dir string) {
	context := gen_apiV2.NewApiContext()
	context.Load(tree.FullPath)
	context.Parse()
	for _, file := range context.Files {
		if len(file.Funcs) != 0 {
			GenApiV2(context.Files, context.ReginsterMap, public.GetGenConf().BaseConf,Dir)
			break
		}
	}

	for _, node := range tree.Children {
		if node.Info.IsDir {
			depthGen(node, path.Join(Dir, node.Info.Name))
		}
	}
}