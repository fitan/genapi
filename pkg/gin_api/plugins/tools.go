package plugins

import (
	_ "embed"
	"github.com/fitan/genapi/public"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"strconv"
	"strings"
)


//go:embed interfacepkg/casbin.go
var casbinFile string

type FileMsg struct {
	Name string
	Src string
}

var interfaceFiles = []FileMsg{{
	Name: "casbin.go",
	Src:  casbinFile,
}}

var FindPluginInterfaceMap = make(map[string]*types.Interface)

func init() {
	for _, interfaceFile := range interfaceFiles {

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, interfaceFile.Name, interfaceFile.Src, parser.AllErrors)
		if err != nil {
			log.Panic(err)
		}
		info := &types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue, 0),
			Defs:       make(map[*ast.Ident]types.Object, 0),
			Uses:       make(map[*ast.Ident]types.Object, 0),
			Implicits:  make(map[ast.Node]types.Object, 0),
			Selections: make(map[*ast.SelectorExpr]*types.Selection, 0),
			Scopes:     make(map[ast.Node]*types.Scope, 0),
			InitOrder:  make([]*types.Initializer, 0, 0),
		}
		_, err = new(types.Config).Check(interfaceFile.Name, fset, []*ast.File{f}, info)
		if err != nil {
			log.Panic(err)
		}

		log.Println(FindPluginInterfaceMap)
		findInterface(f, info)
		log.Println("find plugin interface map", FindPluginInterfaceMap)
	}
}

func findInterface(f *ast.File, info *types.Info) {
	ast.Inspect(f, func(node ast.Node) bool {
		if typeSpec, ok := node.(*ast.TypeSpec); ok {
			if interfaceType, ok := typeSpec.Type.(*ast.InterfaceType); ok {
				FindPluginInterfaceMap[typeSpec.Name.Name] = info.TypeOf(interfaceType).(*types.Interface)
				return false
			}
		}
		return true
	})
}

func CheckMatch(match public.Match, docFields []string, inFieldType types.Type, outFieldType types.Type) bool {
	if CheckHasInterface(outFieldType, match.OutInterfaceName) && CheckHasInterface(inFieldType, match.InInterfaceName) && CheckParamMatch(match.Param, docFields) {
		return true
	}
	return false
}

func CheckParamMatch(matchParam []string, docFields []string) bool {
	for _, syntax := range matchParam {
		parseSyntax := strings.Split(syntax, "=")
		if len(parseSyntax) != 2 {
			log.Fatalln("syntax " + syntax +  " not x=x")
		}
		index,err :=  strconv.Atoi(parseSyntax[0])
		if err != nil {
			log.Fatalln("synctx " + syntax + " " + parseSyntax[0] + " inconvertible int")
		}
		key := parseSyntax[1]

		if docFields[index + 2] == key {
			continue
		} else {
			return false
		}
	}
	return true
}

func CheckHasInterface(t types.Type, interfaceNames []string) bool {
	for _, name := range interfaceNames {
		if i, ok := FindPluginInterfaceMap[name]; ok {
			has := types.Implements(t,i)
			if !has {
				return false
			}
		} else {
			log.Fatalln("not found " + name)
		}
	}
	return true
}

type PointTemplate struct {
	Name string
	Has bool
	Keys map[string]string
	BindBefor HandlerTemplate
	BindAfter HandlerTemplate
}

type CallBackTemplate struct {
	Has bool
	Keys map[string]string
	Template HandlerTemplate
}

type HandlerTemplate struct {
	ImportPath string
	Template   string
}
