package main

import (
	_ "cmdb/docs"
	"cmdb/public"
	"cmdb/routers"
)

// @title cmdbapi
// @version 1.0
// @description RESTful API 文档.
// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	routers.GetDefaultRouter().Run(public.GetConf().App.Host + ":" + public.GetConf().App.Port)
}
