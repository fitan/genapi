package routers

import (
	"cmdb/gen/entrest"
	"cmdb/gen/handler/ent"
	"cmdb/gen/handler/logic/casbin"
	"cmdb/gen/handler/logic/tree"
	"cmdb/gen/handler/logic/user"
	"cmdb/middleware/jwt"
	"cmdb/public"
	public2 "github.com/fitan/genapi/public"
	"github.com/gin-gonic/gin"
)

func init() {
	db := public.GetDB()
	entrest.NewCURDALL(db)
	role_method := make([]public2.CasbinRoleMethod,0,0)
	ent.Register(GetDefaultRouter(), &role_method)
	user.Register(GetAuthRouter(), &role_method)
	tree.Register(GetDefaultRouter(), &role_method)
	//casbin.Register(GetAuthRouter())
	casbin.Register(GetDefaultRouter(), &role_method)

	GetDefaultRouter().GET("/role_method", func(c *gin.Context) {
		c.JSON(200, role_method)
	})
	GetDefaultRouter().POST("/login", jwt.GetMyJwtMid().LoginHandler)
	GetDefaultRouter().POST("/logout", jwt.GetMyJwtMid().LogoutHandler)
	GetAuthRouter().GET("/refresh_token", jwt.GetMyJwtMid().RefreshHandler)
	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
