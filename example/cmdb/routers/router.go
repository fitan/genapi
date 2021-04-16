package routers

import (
	"cmdb/middleware/jwt"
	"github.com/gin-gonic/gin"
)

var defaultRouter *gin.Engine
var authRouter *gin.RouterGroup

func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}

func GetDefaultRouter() *gin.Engine {
	if defaultRouter == nil {
		defaultRouter = NewRouter()
	}
	return defaultRouter
}

func GetAuthRouter() *gin.RouterGroup {
	if authRouter == nil {
		authRouter = GetDefaultRouter().Group("/auth")
		authRouter.Use(jwt.GetMyJwtMid().MiddlewareFunc())
	}
	return authRouter
}
