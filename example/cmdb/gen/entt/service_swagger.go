package entt

import (
	"time"

	"github.com/gin-gonic/gin"
)

type ServiceID struct {
	ID int `json:"id,omitempty"`
}

type ServiceNode struct {
	ServiceID
	ServiceNodeNotID
}

type ServiceEdges struct {
	RoleBindings []*RoleBindingID `json:"role_bindings,omitempty"`

	Servers []*ServerID `json:"servers,omitempty"`

	Project *ProjectID `json:"project,omitempty"`
}

type ServiceNodeNotID struct {
	CreateTime time.Time    `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime time.Time    `json:"update_time,omitempty"  format:"date-time" `
	Name       string       `json:"name,omitempty"   `
	Edges      ServiceEdges `json:"edges"`
}

type ServiceQuery struct {
}

// @Summary create one service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param body body ServiceNodeNotID true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /service [post]
func ServiceCreateOne(c *gin.Context) {
}

// @Summary create list service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param body body []ServiceNode true " "
// @Success 200 {object} RestReturn{data=[]ServiceNodeNotID}
// @Router /services [post]
func ServiceCreateList(c *gin.Context) {
}

// @Summary delete one service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /service/{id} [delete]
func ServiceDeleteOne(c *gin.Context) {
}

// @Summary delete list service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /services [delete]
func ServiceDeleteList(c *gin.Context) {
}

// @Summary get one service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /service/{id} [get]
func ServiceGetOne(c *gin.Context) {
}

// @Summary get list service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param data query ServiceQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetServiceListData}
// @Router /services [get]
func ServiceGetList(c *gin.Context) {
}

// @Summary update one service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param id path int true " "
// @Param body body ServiceNodeNotID true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /service/{id} [put]
func ServiceUpdateOne(c *gin.Context) {
}

// @Summary update list service
// @Accept  json
// @Produce  json
// @Tags Service
// @Param body body []ServiceNode true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /services [put]
func ServiceUpdateList(c *gin.Context) {
}

// @Summary create list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Service
// @Param id path int true " "
// @Param body body  []RoleBindingNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNodeNotID}
// @Router /service/{id}/role_bindings [post]
func CreateListRoleBindingsByService(c *gin.Context) {
}

// @Summary delete one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Service
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /service/{id}/role_bindings [delete]
func DeleteListRoleBindingsByService(c *gin.Context) {
}

// @Summary get list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by Service
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNode}
// @Router /service/{id}/role_bindings [get]
func GetListRoleBindingsByService(c *gin.Context) {
}

// @Summary create list server
// @Accept  json
// @Produce  json
// @Tags Server by Service
// @Param id path int true " "
// @Param body body  []ServerNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]ServerNodeNotID}
// @Router /service/{id}/servers [post]
func CreateListServersByService(c *gin.Context) {
}

// @Summary delete one server
// @Accept  json
// @Produce  json
// @Tags Server by Service
// @Param id path int true " "
// @Param query query  ServerQuery false " "
// @Success 200 {object} RestReturn{data=ServerNode}
// @Router /service/{id}/servers [delete]
func DeleteListServersByService(c *gin.Context) {
}

// @Summary get list server
// @Accept  json
// @Produce  json
// @Tags Server by Service
// @Param id path int true " "
// @Param query query  ServerQuery false " "
// @Success 200 {object} RestReturn{data=[]ServerNode}
// @Router /service/{id}/servers [get]
func GetListServersByService(c *gin.Context) {
}

// @Summary create one project
// @Accept  json
// @Produce  json
// @Tags Project by Service
// @Param id path int true " "
// @Param body body  ProjectNodeNotID true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /service/{id}/project [post]
func CreateOneProjectByService(c *gin.Context) {
}

// @Summary delete one project
// @Accept  json
// @Produce  json
// @Tags Project by Service
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /service/{id}/project [delete]
func DeleteOneProjectByService(c *gin.Context) {
}

// @Summary get one project
// @Accept  json
// @Produce  json
// @Tags Project by Service
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /service/{id}/project [get]
func GetOneProjectByService(c *gin.Context) {
}
