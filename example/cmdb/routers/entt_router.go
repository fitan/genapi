package routers

import (
	"cmdb/gen/entrest"
	"cmdb/gen/handler/ent"
	"cmdb/gen/handler/logic/casbin"
	"cmdb/gen/handler/logic/tree"
	"cmdb/gen/handler/logic/user"
	"cmdb/middleware/jwt"
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	entrest.NewCURDALL(db)
	ent.Register(GetDefaultRouter())
	user.Register(GetAuthRouter())
	tree.Register(GetDefaultRouter())
	//casbin.Register(GetAuthRouter())
	casbin.Register(GetDefaultRouter())
	GetDefaultRouter().POST("/login", jwt.GetMyJwtMid().LoginHandler)
	GetDefaultRouter().POST("/logout", jwt.GetMyJwtMid().LogoutHandler)
	GetAuthRouter().GET("/refresh_token", jwt.GetMyJwtMid().RefreshHandler)
	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
