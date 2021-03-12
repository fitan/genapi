package main

import (
	"cmdb/gen/entt"
	"cmdb/public"
	"github.com/gin-gonic/gin"
)

func main() {
	db := public.GetDB()

	curd := entt.NewCURDALL(db)

	r := gin.Default()
	curd.RegisterRouterALL(r)
	r.Run()
}
