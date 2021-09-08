package routers

import (
	"cmdb/middleware/jwt"
	"cmdb/public"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"sync"
)

var defaultRouter *gin.Engine
var defaultRouterLock sync.Mutex
var authRouter *gin.RouterGroup
var authRouterLock sync.Mutex

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		public.GetXLog().Info().Str("handlername", c.HandlerName()).Strs("handlernames", c.HandlerNames()).Msg(c.FullPath())
	})
	r.Use(otelgin.Middleware("my-cmdb"))
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

var apiRouter *gin.RouterGroup

func GetApiRouter() *gin.RouterGroup {
	if apiRouter == nil {
		return GetDefaultRouter().Group("/api")
	}
	return apiRouter
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
