package main

import (
	_ "ent_samp/docs"
	"ent_samp/gen"
	"ent_samp/public"
	genrest "ent_samp/service"
	"github.com/davecgh/go-spew/spew"
	"github.com/fitan/genapi/pkg"

	//genrest "ent_samp/service"
	//"github.com/fitan/genapi/pkg"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	parse := pkg.ParseFuncApi("ent_smap", "./api", "./gen")
	spew.Dump(parse.ApiMap)
	return
	//pkg.Load("./ent/schema", "./", nil)
	r := gin.Default()
	curd := genrest.NewCURDALL(public.GetDB())
	curd.RegisterRouterALL(r)
	gen.RegisterAll(r)
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.Run()
}
