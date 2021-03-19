package entt

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ProjectID struct {
	ID int `json:"id,omitempty"`
}

type ProjectNode struct {
	ProjectID
	ProjectNodeNotID
}

type ProjectEdges struct {
}

type ProjectNodeNotID struct {
	CreateTime time.Time    `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime time.Time    `json:"update_time,omitempty"  format:"date-time" `
	Name       string       `json:"name,omitempty"   `
	Edges      ProjectEdges `json:"edges"`
}

type ProjectQuery struct {
	Includes []string `json:"includes" form:"includes" enums:"role_binding.service,role_binding.user,service.server,service.role_binding,role_binding,role_binding.service.server,role_binding.user.alert,service,service.role_binding.user,service.role_binding.user.alert"`
}

// @Summary create one project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param body body ProjectNodeNotID true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /project [post]
func ProjectCreateOne(c *gin.Context) {
}

// @Summary create list project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param body body []ProjectNode true " "
// @Success 200 {object} RestReturn{data=[]ProjectNodeNotID}
// @Router /projects [post]
func ProjectCreateList(c *gin.Context) {
}

// @Summary delete one project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /project/{id} [delete]
func ProjectDeleteOne(c *gin.Context) {
}

// @Summary delete list project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /projects [delete]
func ProjectDeleteList(c *gin.Context) {
}

// @Summary get one project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /project/{id} [get]
func ProjectGetOne(c *gin.Context) {
}

// @Summary get list project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param data query ProjectQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetProjectListData}
// @Router /projects [get]
func ProjectGetList(c *gin.Context) {
}

// @Summary update one project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param id path int true " "
// @Param body body ProjectNodeNotID true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /project/{id} [put]
func ProjectUpdateOne(c *gin.Context) {
}

// @Summary update list project
// @Accept  json
// @Produce  json
// @Tags Project
// @Param body body []ProjectNode true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /projects [put]
func ProjectUpdateList(c *gin.Context) {
}

// @Summary create list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Project
// @Param id path int true " "
// @Param body body  []RoleBindingNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNodeNotID}
// @Router /project/{id}/role_bindings [post]
func CreateListRoleBindingsByProject(c *gin.Context) {
}

// @Summary delete one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Project
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /project/{id}/role_bindings [delete]
func DeleteListRoleBindingsByProject(c *gin.Context) {
}

// @Summary get list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Project
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNode}
// @Router /project/{id}/role_bindings [get]
func GetListRoleBindingsByProject(c *gin.Context) {
}

// @Summary create list service
// @Accept  json
// @Produce  json
// @Tags Service by Project
// @Param id path int true " "
// @Param body body  []ServiceNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]ServiceNodeNotID}
// @Router /project/{id}/services [post]
func CreateListServicesByProject(c *gin.Context) {
}

// @Summary delete one service
// @Accept  json
// @Produce  json
// @Tags Service by Project
// @Param id path int true " "
// @Param query query  ServiceQuery false " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /project/{id}/services [delete]
func DeleteListServicesByProject(c *gin.Context) {
}

// @Summary get list service
// @Accept  json
// @Produce  json
// @Tags Service by Project
// @Param id path int true " "
// @Param query query  ServiceQuery false " "
// @Success 200 {object} RestReturn{data=[]ServiceNode}
// @Router /project/{id}/services [get]
func GetListServicesByProject(c *gin.Context) {
}
