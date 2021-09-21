package core

import (
	"cmdb/pkg/httpclient"
	"github.com/go-resty/resty/v2"
)

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

type apiRegister struct {
}

func (a apiRegister) Set(c *Core) {
	c.Api = NewApi()
}

func (a apiRegister) Unset(c *Core) {
	c.Api = nil
}
