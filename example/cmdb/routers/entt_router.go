package routers

import (
	"cmdb/gen/entt"
	"cmdb/public"
)

func init()  {
	db := public.GetDB()
	curdall := entt.NewCURDALL(db)
	curdall.RegisterRouterALL(GetDefaultRouter())
}
