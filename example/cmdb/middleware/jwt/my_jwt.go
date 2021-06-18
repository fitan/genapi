package jwt

import (
	"cmdb/public"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var myJwtMid *jwt.GinJWTMiddleware

type AdminAuthorizator struct {
}

func (a AdminAuthorizator) Authorizator(data interface{}, c *gin.Context) bool {
	path := c.FullPath()
	method := c.Request.Method
	var role string
	if v, ok := data.(string); ok {
		role = v
	} else {
		return false
	}
	//err := models.GetCasbin().LoadPolicy()
	//if err != nil {
	//	public.GetXLog().Error().Err(err).Msg("")
	//	return false
	//}
	c.Set("role",data)
	has, err := public.GetCasbin().Enforce(role, path, method)
	if err != nil {
		public.GetXLog().Error().Err(err).Msg("")
		return false
	}
	if has {
		return true
	}
	return false
}
func GetMyJwtMid() *jwt.GinJWTMiddleware {
	if myJwtMid == nil {
		middleware, err := NewAuthMiddleware(AdminAuthorizator{})
		if err != nil {
			public.GetXLog().Fatal().Err(err).Msg("")
		}
		myJwtMid = middleware
	}
	return myJwtMid
}
