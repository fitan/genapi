package routers

import (
	"cmdb/public"
)

func init() {
	db := public.GetDB()
	_ = entt2.NewCURDALL(db)
	ent_handler.Register(GetApiRouter())

	//curdall := entt.NewCURDALL(db)
	//curdall.RegisterRouterALL(GetDefaultRouter())
}
