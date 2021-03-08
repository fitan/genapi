package gen

import (
	"github.com/gin-gonic/gin"
)

func RegisterAll(c *gin.Engine) {

	c.GET("/car", func(c *gin.Context) {
		GinResult(c, Car)
	})

	c.GET("/user", func(c *gin.Context) {
		GinResult(c, Hello)
	})

}

type Result struct {
	Code int `json:"code"`
	Data interface{}
	Err  error
}

func GinResult(c *gin.Context, fc func(c *gin.Context) (data interface{}, err error)) {
	data, err := fc(c)
	res := Result{
		Code: 0,
		Data: data,
		Err:  err,
	}
	if err != nil {
		res.Code = 503
	} else {
		res.Code = 200
	}
}
