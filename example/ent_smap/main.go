package main

import (
	_ "ent_samp/docs"
	"ent_samp/gen/entt"
	"ent_samp/gen/router"
	"ent_samp/public"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"

	//genrest "ent_samp/service"
	//"github.com/fitan/genapi/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	profiler.Start(profiler.Config{
		ApplicationName: "backend.ent_smap",
		ServerAddress:   "http://10.143.131.148:4040",
	})
	r := gin.Default()
	curd := entt.NewCURDALL(public.GetDB())
	curd.RegisterRouterALL(r)
	router.RegisterAll(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
