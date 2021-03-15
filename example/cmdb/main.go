package main

import (
	_ "cmdb/docs"
	"cmdb/ent"
	"cmdb/ent/rolebinding"
	"cmdb/ent/server"
	"cmdb/gen/entt"
	"cmdb/public"
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title cmdbapi
// @version 1.0
// @description RESTful API 文档.
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	db := public.GetDB()
	curdall := entt.NewCURDALL(db)
	r := gin.Default()
	r.GET("test", func(c *gin.Context) {
		all, _ := db.User.Query().WithRoleBindings(func(query *ent.RoleBindingQuery) {
			query.Select(rolebinding.FieldRole).WithService(func(query *ent.ServiceQuery) {
				query.WithServers(func(query *ent.ServerQuery) {
					query.Select(server.FieldIP)
				})
			})
		}).All(context.Background())
		c.JSON(200, all)
	})
	curdall.RegisterRouterALL(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()

}
