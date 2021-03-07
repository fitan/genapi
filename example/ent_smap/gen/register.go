package gen

import "github.com/gin-gonic/gin"

func RegisterAll(c *gin.Engine) {
	c.GET("/post", Hello)
	c.POST("/post",Hello)
}