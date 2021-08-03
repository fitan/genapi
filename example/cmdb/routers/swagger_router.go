package routers

import (
	_ "cmdb/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")

	GetDefaultRouter().GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

