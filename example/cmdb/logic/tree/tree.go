package tree

import (
	"cmdb/ent"
	"cmdb/ent/servicetree"
	"cmdb/gen/entrest"
	"cmdb/public"
	"context"
	"github.com/gin-gonic/gin"
)

type GetServiceTreeIn struct {
	Query struct {
		entrest.ServiceTreePaging 
	} `json:"query"`
}

// @Description
// @Summary 获取服务树
// @GenApi /api/tree [get]
func GetServiceTree(c *gin.Context, in *GetServiceTreeIn) ([]*ent.ServiceTree, error) {
	queryer := public.GetDB().ServiceTree.Query()
	in.Query.BindPagingServiceTree(queryer)

	return queryer.Where(servicetree.Not(servicetree.HasProject())).WithService().All(context.Background())
}

type CreateProjectIn struct {
	Body ent.ServiceTree `json:"body"`
}


// @Summary 创建project
// @GenApi /api/project [post]
func CreateProject(c *gin.Context, in *CreateProjectIn) (*ent.ServiceTree, error) {
	return public.GetDB().ServiceTree.Create().SetName(in.Body.Name).SetNote(in.Body.Note).SetType(in.Body.Type).Save(context.Background())
}

type CreateServiceIn struct {
	Uri struct {
		Id int `json:"id" uri:"id"`
	} `json:"uri"`
	Body ent.ServiceTree `json:"body"`
}

// @Summary 创建服务
// @GenApi /api/project/{id}/service [post]
func CreateService(c *gin.Context, in *CreateServiceIn) (*ent.ServiceTree, error) {
	return public.GetDB().ServiceTree.Create().SetProjectID(in.Uri.Id).SetName(in.Body.Name).SetNote(in.Body.Note).SetType(in.Body.Type).Save(context.Background())
}
