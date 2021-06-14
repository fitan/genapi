package ent_handler

import (
	"cmdb/gen/entt2"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.Alert}
// @Router /alert/:id [get]
func GetOneAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneAlertIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.AlertQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetAlertListData}
// @Router /alerts [get]
func GetListAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListAlertIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Alert true " "
// @Success 200 {object} Result{data=ent.Alert}
// @Router /alert [post]
func CreateOneAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneAlertIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Alerts true " "
// @Success 200 {object} Result{data=ent.Alerts}
// @Router /alerts [post]
func CreateListAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListAlertIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Alert true " "
// @Success 200 {object} Result{data=ent.Alert}
// @Router /alert [put]
func UpdateOneAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneAlertIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.Alerts true " "
// @Success 200 {object} Result{data=string}
// @Router /alerts [put]
func UpdateListAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListAlertIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /alert/:id [delete]
func DeleteOneAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneAlertIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneAlert(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.AlertIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /alerts [delete]
func DeleteListAlert(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListAlertIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListAlert(c, in)
}
