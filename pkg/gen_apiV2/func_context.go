package gen_apiV2


type Func struct {
	Comments []string
	Router Router
	FuncName string
	Bind Bind
	ParamIn1 string
	ResOut0 string
}

type Router struct{
	Method string
	GinPath string
}

type QuoteType string

const StructType QuoteType = "struct"
const IdentType QuoteType = "ident"

type Bind struct {
	Uri struct {
		Has bool
		Param []string
		TagMsgs []TagMsg
	}
	Body struct{
		Has       bool
		QuoteType QuoteType
		SwagStructName string
		SwagRaw   string
		SwagObj string
	}
	Query struct {
		Has       bool
		QuoteType QuoteType
		SwagStructName string
		SwagRaw   string
		SwagObj string
	}
	Header struct {
		Has       bool
		QuoteType QuoteType
		SwagStructName string
		SwagRaw   string
		SwagObj string
		TagMsgs []TagMsg
	}
}