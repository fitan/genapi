package core

import (
	"cmdb/pkg/log"
	"github.com/gin-gonic/gin"
)

type WrapFn func(core *Core)


type Core struct {
	Log     *log.Xlog

	GinCtx  *gin.Context

	Api     *Api

	wrapMid []WrapFn

	endBefor WrapFn

	handlerData struct{
		inData interface{}
		outData interface{}
		outErr error
	}

}

func (c *Core) Execute() {
	for _, wrapMid := range c.wrapMid {
		wrapMid(c)
	}

	c.endBefor(c)
}

func (c *Core) SetGinCtx(ginCtx *gin.Context) {
	c.GinCtx = ginCtx
}

func (c *Core) SetHandlerIn(in interface{})  {
	c.handlerData.inData = in
}

func (c *Core) SetHandlerOut(data interface{}, err error) {
	c.handlerData.outData = data
	c.handlerData.outErr = err
}

func (c *Core) unsetHandlerData() {
	c.handlerData.outData = nil
	c.handlerData.outErr = nil
	c.handlerData.inData = nil
}

func (c *Core) unsetGinCtx() {
	c.GinCtx = nil
}

func (c *Core) UnSet() {
	c.unsetGinCtx()
	c.unsetHandlerData()
	corePool.Put(c)
}

func GetCore() interface{} {
	return corePool.Get()
}

