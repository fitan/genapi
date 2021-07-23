package routers

import (
	"cmdb/gen/entrest"
	"cmdb/gen/handler/ent"
	"cmdb/gen/handler/logic"
	"cmdb/gen/handler/logic/casbin"
	"cmdb/middleware/jwt"
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	entrest.NewCURDALL(db)
	ent.Register(GetDefaultRouter())
	logic.Register(GetDefaultRouter())
	casbin.Register(GetAuthRouter())
	GetDefaultRouter().POST("/login", jwt.GetMyJwtMid().LoginHandler)
	GetDefaultRouter().POST("/logout", jwt.GetMyJwtMid().LogoutHandler)
	GetAuthRouter().GET("/refresh_token", jwt.GetMyJwtMid().RefreshHandler)
	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
