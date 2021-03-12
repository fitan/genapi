package entt

import (
	"ent_samp/ent/user"

	"github.com/gin-gonic/gin"
)

type UserNode struct {
	ID int `json:"id,omitempty"`

	Name string  `json:"name,omitempty"   `
	Age1 int     `json:"age1,omitempty"   `
	En   user.En `json:"en,omitempty"   enums:"1,2,3" binding:"oneof=1,2,3"`
}

type UserNodeNotID struct {
	Name string  `json:"name,omitempty"   `
	Age1 int     `json:"age1,omitempty"   `
	En   user.En `json:"en,omitempty"   enums:"1,2,3" binding:"oneof=1,2,3"`
}

type UserQuery struct {
	UserAge1EQ

	UserPaging
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

// @Summary create list car
// @Accept  json
// @Produce  json
// @Tags Car by User
// @Param id path int true " "
// @Param body body  []CarNodeNotID true " "
// @Success 200 {object} RestReturn{data=[]CarNodeNotID}
// @Router /user/{id}/cars [post]
func CreateListCarsByUser(c *gin.Context) {
}

// @Summary delete one car
// @Accept  json
// @Produce  json
// @Tags Car by User
// @Param id path int true " "
// @Param query query  CarQuery false " "
// @Success 200 {object} RestReturn{data=CarNode}
// @Router /user/{id}/cars [delete]
func DeleteListCarsByUser(c *gin.Context) {
}

// @Summary get list car
// @Accept  json
// @Produce  json
// @Tags Car by User
// @Param id path int true " "
// @Param query query  CarQuery false " "
// @Success 200 {object} RestReturn{data=[]CarNode}
// @Router /user/{id}/cars [get]
func GetListCarsByUser(c *gin.Context) {
}
