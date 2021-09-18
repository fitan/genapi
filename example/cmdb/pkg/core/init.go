package core

import (
	"cmdb/pkg/httpclient"
	"cmdb/pkg/log"
	"cmdb/pkg/trace"
	"github.com/go-resty/resty/v2"
	"go.uber.org/zap"
	"sync"
)

var xlog *log.Xlog

func NewXlog() *log.Xlog {
	if xlog != nil {
		return xlog
	}

	return log.NewXlog(log.WithTrace(trace.GetTp(), zap.InfoLevel))
}

var api *Api

type Api struct {
	Cmdb *resty.Client
}

func NewApi() *Api {
	if api != nil {
		return api
	}

	return &Api{
		Cmdb: httpclient.NewClient(httpclient.WithHost("http://www.baidu.com")),
	}
}

var corePool = sync.Pool{
	New: func() interface{} {
		return New(WithResultWrap())
	},
}

func New(fs ...Option) *Core {
	core := &Core{Log: NewXlog(),Api: NewApi()}
	for _,f := range fs {
		f(core)
	}
	return core
}
