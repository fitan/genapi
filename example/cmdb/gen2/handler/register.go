package handler

import (
	"cmdb/public"

	"github.com/gin-gonic/gin"
)

func RegisterAll(r gin.IRouter) {

	r.GET("/api/usercall", func(c *gin.Context) {
		public.GinResult(c, UserCall)
	})

}

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

func GinResult(c *gin.Context, fc func(c *gin.Context) (data interface{}, err error)) {
	data, err := fc(c)
	res := Result{
		Code: 0,
		Data: data,
		Err:  "",
	}
	if err != nil {
		res.Code = 503
		res.Err = err.Error()
	} else {
		res.Code = 200
	}
	c.JSON(200, res)
}