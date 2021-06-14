package routers

import (
	"cmdb/gen/entt2"
	"cmdb/gen2/ent_handler"
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	_ = entt2.NewCURDALL(db)
	ent_handler.RegisterAll(GetDefaultRouter())

	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
