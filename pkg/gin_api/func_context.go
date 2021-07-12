package gen_apiV2

import "github.com/fitan/genapi/pkg/gin_api/plugins"

type Func struct {
	PkgName    string
	Comments   []string
	Router     Router
	FuncName   string
	Bind       Bind
	ParamIn1   string
	ParamIn1Ts []string
	ResOut0    string
	ResOut0Ts  []string
	Plugins    Plugins
}

type Plugins struct {
	Point    []plugins.PointTemplate
	CallBack plugins.CallBackTemplate
}

type Router struct {
	Method         string
	GinPath        string
	RouterGroupKey string
}

type QuoteType string

const StructType QuoteType = "struct"
const IdentType QuoteType = "ident"

type Bind struct {
	Uri struct {
		Has     bool
		Param   []string
		TagMsgs []TagMsg
	}
	Body struct {
		Has            bool
		QuoteType      QuoteType
		Comment        string
		SwagStructName string
		SwagRaw        string
		SwagObj        string
	}
	Query struct {
		Has            bool
		QuoteType      QuoteType
		Comment        string
		SwagStructName string
		SwagRaw        string
		SwagObj        string
	}
	Header struct {
		Has     bool
		TagMsgs []TagMsg
	}
}
