package ent_handler

import (
	"cmdb/gen/entt2"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.Server}
// @Router /server/:id [get]
func GetOneServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneServerIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ServerQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetServerListData}
// @Router /servers [get]
func GetListServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListServerIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Server true " "
// @Success 200 {object} Result{data=ent.Server}
// @Router /server [post]
func CreateOneServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneServerIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Servers true " "
// @Success 200 {object} Result{data=ent.Servers}
// @Router /servers [post]
func CreateListServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListServerIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Server true " "
// @Success 200 {object} Result{data=ent.Server}
// @Router /server [put]
func UpdateOneServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneServerIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Servers true " "
// @Success 200 {object} Result{data=string}
// @Router /servers [put]
func UpdateListServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListServerIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /server/:id [delete]
func DeleteOneServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneServerIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneServer(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ServerIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /servers [delete]
func DeleteListServer(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListServerIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListServer(c, in)
}
