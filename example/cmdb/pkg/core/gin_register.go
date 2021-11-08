package core

import (
	"github.com/gin-gonic/gin"
)

type BindHandlerInterface interface {
	BindIn(c *gin.Context) (interface{}, error)
	BindFn(c *Core) (interface{}, error)
}

type Gin struct {
	*gin.Context
	beforMid    []GinMid
	afterMid    []GinMid
	handlerData struct {
		inData  interface{}
		outData interface{}
		err     error
	}
	resultWrap []Option
}

func (g *Gin) BindFn(core *Core, c *gin.Context, i BindHandlerInterface) {
	defer g.result(core)

	g.setCtx(c)

	err := core.Gin.bindBefor(core)
	if err != nil {
		core.Gin.setResult(nil, err)
		return
	}

	in, err := i.BindIn(c)
	if err != nil {
		core.Gin.setResult(nil, err)
		return
	}

	g.setIn(in)

	err = g.bindAfter(core)
	if err != nil {
		core.Gin.setResult(nil, err)
		return
	}

	g.setResult(i.BindFn(core))
	core.Release()
}

func (g *Gin) setCtx(c *gin.Context) {
	g.Context = c
}

func (g *Gin) Ctx() *gin.Context {
	return g.Context
}

func (g *Gin) setIn(data interface{}) {
	g.handlerData.inData = data
}

func (g *Gin) setResult(data interface{}, err error) {
	g.handlerData.outData = data
	g.handlerData.err = err
}

func (g *Gin) result(c *Core) {
	for _, r := range g.resultWrap {
		r(c)
	}
}

func (g *Gin) bindBefor(c *Core) error {
	for _, m := range g.beforMid {
		err := m(c)
		if err != nil {
			g.handlerData.err = err
			return err
		}
	}
	return nil
}

func (g *Gin) bindAfter(c *Core) error {
	for _, m := range g.afterMid {
		err := m(c)
		if err != nil {
			g.handlerData.err = err
			return err
		}
	}
	return nil
}

type GinOption func(g *Gin)

func NewGin(fs ...GinOption) *Gin {
	g := &Gin{}
	for _, f := range fs {
		f(g)
	}
	return g
}

type ginRegister struct {
}

func (g *ginRegister) Set(c *Core) {
	c.Gin = NewGin(WithWrap(GinResultWrap, GinTraceWrap))
}

func (g *ginRegister) Unset(c *Core) {
	c.Gin.Context = nil
	c.Gin.handlerData.inData = nil
	c.Gin.handlerData.outData = nil
	c.Gin.handlerData.err = nil
}
