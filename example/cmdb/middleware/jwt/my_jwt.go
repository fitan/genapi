package jwt

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
)

var myJwtMid *jwt.GinJWTMiddleware

type AdminAuthorizator struct {

}

func (a AdminAuthorizator)Authorizator(data interface{}, c *gin.Context) bool {
	if v, ok := data.(string); ok && v == "admin" {
		return true
	}
	return false
}
func GetMyJwtMid() *jwt.GinJWTMiddleware {
	if myJwtMid == nil {
		middleware, err := NewAuthMiddleware(AdminAuthorizator{})
		if err != nil {
			log.Fatalln(err)
		}
		myJwtMid = middleware
	}
	return myJwtMid
}
