package routers

import (
	"cmdb/middleware/jwt"
	"github.com/gin-gonic/gin"
)

func init()  {
	GetDefaultRouter().POST("/login", jwt.GetMyJwtMid().LoginHandler)
	GetAuthRouter().GET("/refresh_token", jwt.GetMyJwtMid().RefreshHandler)
	GetAuthRouter().GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"err": "",
			"msg": "hello",
		})
	})
}