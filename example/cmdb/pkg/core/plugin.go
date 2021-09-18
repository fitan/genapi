package core

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Option func(c *Core)

func WithResultWrap() Option  {
	return func(c *Core) {
		c.wrapMid = append(c.wrapMid, ResultWrap)
	}
}

type result struct {
	Code int `json:"code"`
	Msg string `json:"msg,omitempty"`
	Data interface{} `json:"data"`
}

func ResultWrap(c *Core)  {
	res := result{Data: c.handlerData.outData}
	if c.handlerData.outErr != nil {
		res.Msg = c.handlerData.outErr.Error()
		res.Code = 5003
		c.GinCtx.JSON(http.StatusInternalServerError, res)
		return
	}

	res.Code = 2000
	c.GinCtx.JSON(http.StatusOK, res)
}

func WithTraceWrap() Option {
	return func(c *Core) {
		c.endBefor = TraceWrap
	}
}

func TraceWrap(c *Core)  {
	l := c.Log.TraceLog(c.GinCtx.Request.Context(), "core trace")
	inData, _ := json.Marshal(c.handlerData.inData)
	outData, _ := json.Marshal(c.handlerData.outData)
	zf := []zap.Field{zap.String("in", string(inData)), zap.String("out", string(outData))}
	if c.handlerData.outErr != nil {
		l.Error(c.handlerData.outErr.Error(), zap.String("in", string(inData)), zap.String("out", string(outData)))
	} else {
		l.Info("handler info", zf...)
	}
	l.End()
}


