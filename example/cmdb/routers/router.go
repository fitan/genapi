package routers

import (
	"cmdb/middleware/jwt"
	"cmdb/public"
	"github.com/gin-gonic/gin"
	"sync"
)

var defaultRouter *gin.Engine
var defaultRouterLock sync.Mutex
var authRouter *gin.RouterGroup
var authRouterLock sync.Mutex

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		public.GetXLog().Info().Msg(c.FullPath())
	})
	return r
}

func GetDefaultRouter() *gin.Engine {
	if defaultRouter == nil {
		defaultRouterLock.Lock()
		defer defaultRouterLock.Unlock()
		defaultRouter = NewRouter()
	}
	return defaultRouter
}

func GetAuthRouter() *gin.RouterGroup {
	if authRouter == nil {
		authRouterLock.Lock()
		defer authRouterLock.Unlock()
		authRouter = GetDefaultRouter().Group("/auth")
		authRouter.Use(jwt.GetMyJwtMid().MiddlewareFunc())
	}
	return authRouter
}
