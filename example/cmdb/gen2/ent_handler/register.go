package ent_handler

import (
	"cmdb/public"

	"github.com/gin-gonic/gin"
)

// @ui button
func Register(r gin.IRouter) {

	r.GET("/alert/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneAlert)
	})

	r.GET("/alerts", func(c *gin.Context) {
		public.GinResult(c, GetListAlert)
	})

	r.POST("/alert", func(c *gin.Context) {
		public.GinResult(c, CreateOneAlert)
	})

	r.POST("/alerts", func(c *gin.Context) {
		public.GinResult(c, CreateListAlert)
	})

	r.PUT("/alert", func(c *gin.Context) {
		public.GinResult(c, UpdateOneAlert)
	})

	r.PUT("/alerts", func(c *gin.Context) {
		public.GinResult(c, UpdateListAlert)
	})

	r.DELETE("/alert/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneAlert)
	})

	r.DELETE("/alerts", func(c *gin.Context) {
		public.GinResult(c, DeleteListAlert)
	})

	r.GET("/project/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneProject)
	})

	r.GET("/projects", func(c *gin.Context) {
		public.GinResult(c, GetListProject)
	})

	r.POST("/project", func(c *gin.Context) {
		public.GinResult(c, CreateOneProject)
	})

	r.POST("/projects", func(c *gin.Context) {
		public.GinResult(c, CreateListProject)
	})

	r.PUT("/project", func(c *gin.Context) {
		public.GinResult(c, UpdateOneProject)
	})

	r.PUT("/projects", func(c *gin.Context) {
		public.GinResult(c, UpdateListProject)
	})

	r.DELETE("/project/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneProject)
	})

	r.DELETE("/projects", func(c *gin.Context) {
		public.GinResult(c, DeleteListProject)
	})

	r.GET("/rolebinding/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneRoleBinding)
	})

	r.GET("/rolebindings", func(c *gin.Context) {
		public.GinResult(c, GetListRoleBinding)
	})

	r.POST("/rolebinding", func(c *gin.Context) {
		public.GinResult(c, CreateOneRoleBinding)
	})

	r.POST("/rolebindings", func(c *gin.Context) {
		public.GinResult(c, CreateListRoleBinding)
	})

	r.PUT("/rolebinding", func(c *gin.Context) {
		public.GinResult(c, UpdateOneRoleBinding)
	})

	r.PUT("/rolebindings", func(c *gin.Context) {
		public.GinResult(c, UpdateListRoleBinding)
	})

	r.DELETE("/rolebinding/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneRoleBinding)
	})

	r.DELETE("/rolebindings", func(c *gin.Context) {
		public.GinResult(c, DeleteListRoleBinding)
	})

	r.GET("/server/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneServer)
	})

	r.GET("/servers", func(c *gin.Context) {
		public.GinResult(c, GetListServer)
	})

	r.POST("/server", func(c *gin.Context) {
		public.GinResult(c, CreateOneServer)
	})

	r.POST("/servers", func(c *gin.Context) {
		public.GinResult(c, CreateListServer)
	})

	r.PUT("/server", func(c *gin.Context) {
		public.GinResult(c, UpdateOneServer)
	})

	r.PUT("/servers", func(c *gin.Context) {
		public.GinResult(c, UpdateListServer)
	})

	r.DELETE("/server/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneServer)
	})

	r.DELETE("/servers", func(c *gin.Context) {
		public.GinResult(c, DeleteListServer)
	})

	r.GET("/service/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneService)
	})

	r.GET("/services", func(c *gin.Context) {
		public.GinResult(c, GetListService)
	})

	r.POST("/service", func(c *gin.Context) {
		public.GinResult(c, CreateOneService)
	})

	r.POST("/services", func(c *gin.Context) {
		public.GinResult(c, CreateListService)
	})

	r.PUT("/service", func(c *gin.Context) {
		public.GinResult(c, UpdateOneService)
	})

	r.PUT("/services", func(c *gin.Context) {
		public.GinResult(c, UpdateListService)
	})

	r.DELETE("/service/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneService)
	})

	r.DELETE("/services", func(c *gin.Context) {
		public.GinResult(c, DeleteListService)
	})

	r.POST("/user", func(c *gin.Context) {
		public.GinResult(c, CreateOneUser)
	})

	r.PUT("/user", func(c *gin.Context) {
		public.GinResult(c, UpdateOneUser)
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		public.GinResult(c, DeleteOneUser)
	})

	r.DELETE("/users", func(c *gin.Context) {
		public.GinResult(c, DeleteListUser)
	})

}

func RegisterAPI(r gin.IRouter) {

	r.GET("/users", func(c *gin.Context) {
		public.GinResult(c, GetListUser)
	})

	r.PUT("/users", func(c *gin.Context) {
		public.GinResult(c, UpdateListUser)
	})

}

func RegisterAuth(r gin.IRouter) {

	r.GET("/user/:id", func(c *gin.Context) {
		public.GinResult(c, GetOneUser)
	})

	r.POST("/users", func(c *gin.Context) {
		public.GinResult(c, CreateListUser)
	})

}

type Result struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Err  string      `json:"err"`
}

func GinResult(c *gin.Context, fc func(c *gin.Context) (data interface{}, err error)) {
	data, err := fc(c)
	res := Result{
		Code: 0,
		Data: data,
		Err:  "",
	}
	if err != nil {
		res.Code = 503
		res.Err = err.Error()
	} else {
		res.Code = 200
	}
	c.JSON(200, res)
}
