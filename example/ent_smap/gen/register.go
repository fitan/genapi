package gen

import (
	"github.com/gin-gonic/gin"
)

func RegisterAll(c *gin.Engine) {

	c.GET("/genapi/car/:id", func(c *gin.Context) {
		GinResult(c, Car)
	})

	c.GET("/user", func(c *gin.Context) {
		GinResult(c, Hello)
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
