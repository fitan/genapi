package ent_handler

import (
	"cmdb/gen/entt2"
	"github.com/gin-gonic/gin"
)


func GetFullPath(c *gin.Context) interface{} {
	return c.FullPath()
}

func GetMehod(c *gin.Context) interface{} {
	return c.Request.Method
}



// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.User}
// @Router /auth/user/:id [get]
func GetOneUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneUserIn{}



	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.UserQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetUserListData}
// @Router /api/users [get]
func GetListUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListUserIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.User true " "
// @Success 200 {object} Result{data=ent.User}
// @Router /user [post]
func CreateOneUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneUserIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Users true " "
// @Success 200 {object} Result{data=ent.Users}
// @Router /auth/users [post]
func CreateListUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListUserIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.User true " "
// @Success 200 {object} Result{data=ent.User}
// @Router /user [put]
func UpdateOneUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneUserIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Users true " "
// @Success 200 {object} Result{data=string}
// @Router /api/users [put]
func UpdateListUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListUserIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /user/:id [delete]
func DeleteOneUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneUserIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneUser(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.UserIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /users [delete]
func DeleteListUser(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListUserIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListUser(c, in)
}
