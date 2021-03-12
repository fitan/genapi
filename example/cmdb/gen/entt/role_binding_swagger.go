package entt

import (
	"cmdb/ent/rolebinding"
	"time"

	"github.com/gin-gonic/gin"
)

type RoleBindingNode struct {
	ID int `json:"id,omitempty"`

	CreateTime time.Time        `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime time.Time        `json:"update_time,omitempty"  format:"date-time" `
	Role       rolebinding.Role `json:"role,omitempty"   enums:"admin,user" binding:"oneof=admin,user"`
}

type RoleBindingNodeNotID struct {
	CreateTime time.Time        `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime time.Time        `json:"update_time,omitempty"  format:"date-time" `
	Role       rolebinding.Role `json:"role,omitempty"   enums:"admin,user" binding:"oneof=admin,user"`
}

type RoleBindingQuery struct {
}

// @Summary create one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param body body RoleBindingNodeNotID true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_binding [post]
func RoleBindingCreateOne(c *gin.Context) {
}

// @Summary create list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param body body []RoleBindingNode true " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNodeNotID}
// @Router /role_bindings [post]
func RoleBindingCreateList(c *gin.Context) {
}

// @Summary delete one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_binding/{id} [delete]
func RoleBindingDeleteOne(c *gin.Context) {
}

// @Summary delete list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_bindings [delete]
func RoleBindingDeleteList(c *gin.Context) {
}

// @Summary get one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_binding/{id} [get]
func RoleBindingGetOne(c *gin.Context) {
}

// @Summary get list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param data query RoleBindingQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetRoleBindingListData}
// @Router /role_bindings [get]
func RoleBindingGetList(c *gin.Context) {
}

// @Summary update one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param id path int true " "
// @Param body body RoleBindingNodeNotID true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_binding/{id} [put]
func RoleBindingUpdateOne(c *gin.Context) {
}

// @Summary update list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding
// @Param body body []RoleBindingNode true " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /role_bindings [put]
func RoleBindingUpdateList(c *gin.Context) {
}

// @Summary create one project
// @Accept  json
// @Produce  json
// @Tags Project by RoleBinding
// @Param id path int true " "
// @Param body body  ProjectNodeNotID true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /role_binding/{id}/project [post]
func CreateOneProjectByRoleBinding(c *gin.Context) {
}

// @Summary delete one project
// @Accept  json
// @Produce  json
// @Tags Project by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /role_binding/{id}/project [delete]
func DeleteOneProjectByRoleBinding(c *gin.Context) {
}

// @Summary get one project
// @Accept  json
// @Produce  json
// @Tags Project by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ProjectNode}
// @Router /role_binding/{id}/project [get]
func GetOneProjectByRoleBinding(c *gin.Context) {
}

// @Summary create one service
// @Accept  json
// @Produce  json
// @Tags Service by RoleBinding
// @Param id path int true " "
// @Param body body  ServiceNodeNotID true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /role_binding/{id}/service [post]
func CreateOneServiceByRoleBinding(c *gin.Context) {
}

// @Summary delete one service
// @Accept  json
// @Produce  json
// @Tags Service by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /role_binding/{id}/service [delete]
func DeleteOneServiceByRoleBinding(c *gin.Context) {
}

// @Summary get one service
// @Accept  json
// @Produce  json
// @Tags Service by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=ServiceNode}
// @Router /role_binding/{id}/service [get]
func GetOneServiceByRoleBinding(c *gin.Context) {
}

// @Summary create one user
// @Accept  json
// @Produce  json
// @Tags User by RoleBinding
// @Param id path int true " "
// @Param body body  UserNodeNotID true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /role_binding/{id}/user [post]
func CreateOneUserByRoleBinding(c *gin.Context) {
}

// @Summary delete one user
// @Accept  json
// @Produce  json
// @Tags User by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /role_binding/{id}/user [delete]
func DeleteOneUserByRoleBinding(c *gin.Context) {
}

// @Summary get one user
// @Accept  json
// @Produce  json
// @Tags User by RoleBinding
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /role_binding/{id}/user [get]
func GetOneUserByRoleBinding(c *gin.Context) {
}
