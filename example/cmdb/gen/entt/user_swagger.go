package entt

import (
	"cmdb/ent/user"
	"time"

	"github.com/gin-gonic/gin"
)

type UserID struct {
	ID int `json:"id,omitempty"`
}

type UserNode struct {
	UserID
	UserNodeNotID
}

type UserEdges struct {
	RoleBindings []*RoleBindingID `json:"role_bindings,omitempty"`

	Alerts []*AlertID `json:"alerts,omitempty"`
}

type UserNodeNotID struct {
	CreateTime time.Time `json:"create_time,omitempty"  format:"date-time" `
	UpdateTime time.Time `json:"update_time,omitempty"  format:"date-time" `
	Name       string    `json:"name,omitempty"   `
	Password   string    `json:"password,omitempty"   `
	Email      string    `json:"email,omitempty"   `
	Phone      string    `json:"phone,omitempty"   `
	Role       user.Role `json:"role,omitempty"   enums:"user,admin,tourist" binding:"oneof=user,admin,tourist"`
	Edges      UserEdges `json:"edges"`
}

type UserQuery struct {
}

// @Summary create one user
// @Accept  json
// @Produce  json
// @Tags User
// @Param body body UserNodeNotID true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /user [post]
func UserCreateOne(c *gin.Context) {
}

// @Summary create list user
// @Accept  json
// @Produce  json
// @Tags User
// @Param body body []UserNode true " "
// @Success 200 {object} RestReturn{data=[]UserNodeNotID}
// @Router /users [post]
func UserCreateList(c *gin.Context) {
}

// @Summary delete one user
// @Accept  json
// @Produce  json
// @Tags User
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /user/{id} [delete]
func UserDeleteOne(c *gin.Context) {
}

// @Summary delete list user
// @Accept  json
// @Produce  json
// @Tags User
// @Param ids query IdsQuery true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /users [delete]
func UserDeleteList(c *gin.Context) {
}

// @Summary get one user
// @Accept  json
// @Produce  json
// @Tags User
// @Param id path int true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /user/{id} [get]
func UserGetOne(c *gin.Context) {
}

// @Summary get list user
// @Accept  json
// @Produce  json
// @Tags User
// @Param data query UserQuery true " "
// @Header 200 {string} Count "The total amount"
// @Success 200 {object} RestReturn{data=GetUserListData}
// @Router /users [get]
func UserGetList(c *gin.Context) {
}

// @Summary update one user
// @Accept  json
// @Produce  json
// @Tags User
// @Param id path int true " "
// @Param body body UserNodeNotID true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /user/{id} [put]
func UserUpdateOne(c *gin.Context) {
}

// @Summary update list user
// @Accept  json
// @Produce  json
// @Tags User
// @Param body body []UserNode true " "
// @Success 200 {object} RestReturn{data=UserNode}
// @Router /users [put]
func UserUpdateList(c *gin.Context) {
}

// @Summary create list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by User
// @Param id path int true " "
// @Param body body  []RoleBindingNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNodeNotID}
// @Router /user/{id}/role_bindings [post]
func CreateListRoleBindingsByUser(c *gin.Context) {
}

// @Summary delete one role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by User
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=RoleBindingNode}
// @Router /user/{id}/role_bindings [delete]
func DeleteListRoleBindingsByUser(c *gin.Context) {
}

// @Summary get list role_binding
// @Accept  json
// @Produce  json
// @Tags RoleBinding by User
// @Param id path int true " "
// @Param query query  RoleBindingQuery false " "
// @Success 200 {object} RestReturn{data=[]RoleBindingNode}
// @Router /user/{id}/role_bindings [get]
func GetListRoleBindingsByUser(c *gin.Context) {
}

// @Summary create list alert
// @Accept  json
// @Produce  json
// @Tags Alert by User
// @Param id path int true " "
// @Param body body  []AlertNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]AlertNodeNotID}
// @Router /user/{id}/alerts [post]
func CreateListAlertsByUser(c *gin.Context) {
}

// @Summary delete one alert
// @Accept  json
// @Produce  json
// @Tags Alert by User
// @Param id path int true " "
// @Param query query  AlertQuery false " "
// @Success 200 {object} RestReturn{data=AlertNode}
// @Router /user/{id}/alerts [delete]
func DeleteListAlertsByUser(c *gin.Context) {
}

// @Summary get list alert
// @Accept  json
// @Produce  json
// @Tags Alert by User
// @Param id path int true " "
// @Param query query  AlertQuery false " "
// @Success 200 {object} RestReturn{data=[]AlertNode}
// @Router /user/{id}/alerts [get]
func GetListAlertsByUser(c *gin.Context) {
}
