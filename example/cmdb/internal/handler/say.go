package handler

import (
	"cmdb/pkg/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Say struct {
	Echo string `json:"echo"`
}

func (s *Say) BindIn(c *gin.Context) (interface{}, error) {
	c.ShouldBind(s)
	return s, nil
}

func (s *Say) BindFn(c *core.Core) (interface{}, error) {
	return SayHello(c, s)
}

func SayHello(c *core.Core, in *Say) (*Say,error) {
	c.Log.Error("fsdf", zap.String("ip", "1.1.1.1"))
	c.Gin.
	return in, nil
}

func SayHandler(c *gin.Context)  {
	core := core.GetCore().(*core.Core)
	core.Gin.BindFn(core, c, &Say{})

}
