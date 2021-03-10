package main

import (
	_ "ent_samp/docs"
	"ent_samp/gen"
	"ent_samp/genent"
	"ent_samp/public"
	"flag"
	"github.com/davecgh/go-spew/spew"
	"github.com/fitan/genapi/pkg"
	"github.com/pyroscope-io/pyroscope/pkg/agent/profiler"

	//genrest "ent_samp/service"
	//"github.com/fitan/genapi/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	var b bool
	flag.BoolVar(&b, "ast", false, "")
	flag.Parse()
	if b {
		parse := pkg.ParseFuncApi("ent_smap", "./api", "./gen")
		spew.Dump(parse.ApiMap)
		return
	}
	//return
	//pkg.Load("./ent/schema", "./genent")
	profiler.Start(profiler.Config{
		ApplicationName: "backend.ent_smap",
		ServerAddress:   "http://10.143.131.148:4040",
	})
	r := gin.Default()
	curd := genent.NewCURDALL(public.GetDB())
	curd.RegisterRouterALL(r)
	gen.RegisterAll(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
