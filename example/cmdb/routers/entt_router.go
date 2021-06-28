package routers

import (
	"cmdb/gen/entrest"
	"cmdb/gen/handler/ent"
	"cmdb/gen/handler/logic"
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	entrest.NewCURDALL(db)
	ent.Register(GetDefaultRouter())
	logic.Register(GetDefaultRouter())

	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
