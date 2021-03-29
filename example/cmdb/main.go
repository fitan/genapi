package main

import (
	_ "cmdb/docs"
	"cmdb/ent"
	"cmdb/ent/rolebinding"
	"cmdb/ent/server"
	"cmdb/gen/entt"
	"cmdb/gen/router"
	"cmdb/public"
	"context"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"
)

// @title cmdbapi
// @version 1.0
// @description RESTful API 文档.
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.IsDebugging() {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	ginLog := zerolog.New(os.Stdout).With().Logger()

	db := public.GetDB()
	curdall := entt.NewCURDALL(db)
	r := gin.Default()
	r.Use(logger.SetLogger(logger.Config{
		Logger:         &ginLog,
		UTC:            false,
		SkipPath:       nil,
		SkipPathRegexp: nil,
	}))
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
	router.RegisterAll(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()

}
