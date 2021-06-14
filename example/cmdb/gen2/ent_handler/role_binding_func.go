package ent_handler

import (
	"cmdb/gen/entt2"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=ent.RoleBinding}
// @Router /rolebinding/:id [get]
func GetOneRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetOneRoleBindingIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.GetOneRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.RoleBindingQueryOption false " "
// @Success 200 {object} Result{data=entt2.GetRoleBindingListData}
// @Router /rolebindings [get]
func GetListRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.GetListRoleBindingIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.GetListRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.RoleBinding true " "
// @Success 200 {object} Result{data=ent.RoleBinding}
// @Router /rolebinding [post]
func CreateOneRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateOneRoleBindingIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateOneRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.RoleBindings true " "
// @Success 200 {object} Result{data=ent.RoleBindings}
// @Router /rolebindings [post]
func CreateListRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.CreateListRoleBindingIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.CreateListRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.RoleBinding true " "
// @Success 200 {object} Result{data=ent.RoleBinding}
// @Router /rolebinding [put]
func UpdateOneRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateOneRoleBindingIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateOneRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body ent.RoleBindings true " "
// @Success 200 {object} Result{data=string}
// @Router /rolebindings [put]
func UpdateListRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.UpdateListRoleBindingIn{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	return entt2.UpdateListRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param id path string true " "
// @Success 200 {object} Result{data=int}
// @Router /rolebinding/:id [delete]
func DeleteOneRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteOneRoleBindingIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteOneRoleBinding(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query entt2.RoleBindingIDs false " "
// @Success 200 {object} Result{data=int}
// @Router /rolebindings [delete]
func DeleteListRoleBinding(c *gin.Context) (data interface{}, err error) {
	in := &entt2.DeleteListRoleBindingIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return entt2.DeleteListRoleBinding(c, in)
}
