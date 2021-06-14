package ent_handler

import (
	"cmdb/gen/entt2"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.Service}
// @Router /service/:id [get]
func GetOneService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneServiceIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneService(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ServiceQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetServiceListData}
// @Router /services [get]
func GetListService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListServiceIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListService(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Service true " "
// @Success 200 {object} Result{data=ent.Service}
// @Router /service [post]
func CreateOneService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneServiceIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneService(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Services true " "
// @Success 200 {object} Result{data=ent.Services}
// @Router /services [post]
func CreateListService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListServiceIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListService(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Service true " "
// @Success 200 {object} Result{data=ent.Service}
// @Router /service [put]
func UpdateOneService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneServiceIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneService(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Services true " "
// @Success 200 {object} Result{data=string}
// @Router /services [put]
func UpdateListService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListServiceIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListService(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /service/:id [delete]
func DeleteOneService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneServiceIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneService(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ServiceIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /services [delete]
func DeleteListService(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListServiceIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListService(c, in)
}
