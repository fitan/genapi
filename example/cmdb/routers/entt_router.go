package routers

import (
	"cmdb/gen/entrest"
	"cmdb/gen/handler/ent"
	"cmdb/gen/handler/logic"
	"cmdb/gen/handler/logic/casbin"
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	entrest.NewCURDALL(db)
	ent.Register(GetDefaultRouter())
	logic.Register(GetDefaultRouter())
	casbin.Register(GetDefaultRouter())

	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
