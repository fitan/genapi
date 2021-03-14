package main

import (
	_ "cmdb/docs"
	"cmdb/gen/entt"
	"cmdb/public"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	db := public.GetDB()
	curdall := entt.NewCURDALL(db)
	r := gin.Default()
	curdall.RegisterRouterALL(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()

}
