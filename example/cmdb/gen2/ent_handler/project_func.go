package ent_handler

import (
	"cmdb/gen/entt2"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.Project}
// @Router /project/:id [get]
func GetOneProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneProjectIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ProjectQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetProjectListData}
// @Router /projects [get]
func GetListProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListProjectIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Project true " "
// @Success 200 {object} Result{data=ent.Project}
// @Router /project [post]
func CreateOneProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneProjectIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Projects true " "
// @Success 200 {object} Result{data=ent.Projects}
// @Router /projects [post]
func CreateListProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListProjectIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Project true " "
// @Success 200 {object} Result{data=ent.Project}
// @Router /project [put]
func UpdateOneProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneProjectIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Projects true " "
// @Success 200 {object} Result{data=string}
// @Router /projects [put]
func UpdateListProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListProjectIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /project/:id [delete]
func DeleteOneProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneProjectIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneProject(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.ProjectIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /projects [delete]
func DeleteListProject(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListProjectIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListProject(c, in)
}
