package core

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type GinMid func(c *Core) error

func WithAfterMid(ops ...GinMid) GinOption {
	mid := make([]GinMid, 0, len(ops))
	for _, f := range ops {
		mid = append(mid, f)
	}
	return func(g *Gin) {
		g.afterMid = mid
	}
}

func WithBeforMid(ops ...GinMid) GinOption {
	mid := make([]GinMid, 0, len(ops))
	for _, f := range ops {
		mid = append(mid, f)
	}
	return func(g *Gin) {
		g.beforMid = mid
	}
}

func WithWrap(ops ...Option) GinOption {
	wrap := make([]Option, 0, len(ops))
	for _, o := range ops {
		wrap = append(wrap, o)
	}
	return func(g *Gin) {
		g.resultWrap = wrap
	}
}

func GinResultWrap(c *Core) {
	res := result{Data: c.Gin.handlerData.outData}
	if c.Gin.handlerData.outData != nil {
		res.Msg = c.Gin.handlerData.err.Error()
		res.Code = 5003
		c.Gin.GinContext.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = 2000
	c.Gin.GinContext.JSON(http.StatusOK, res)
}

type result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

func GinTraceWrap(c *Core) {
	l := c.Log.TraceLog("GinTraceWrap")
	defer l.End()
	inData, _ := json.Marshal(c.Gin.handlerData.inData)
	outData, _ := json.Marshal(c.Gin.handlerData.outData)
	zf := []zap.Field{zap.String("in", string(inData)), zap.String("out", string(outData))}
	if c.Gin.handlerData.err != nil {
		l.Error(c.Gin.handlerData.err.Error(),zf...)
	} else {
		l.Info("handler info", zf...)
	}
}
