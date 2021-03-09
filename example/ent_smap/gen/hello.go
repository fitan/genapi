package gen

import (
	"ent_samp/api"
	"ent_samp/models/api_models"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param query query genrest.UserNameEQ false " "
// @Param id path int true " "
// @Success 200 {object} Result{data=[]ent.User}
// @Router /genapi/car/{id} [get]
func Car(c *gin.Context) (interface{}, error) {
	var err error

	in := &api.CarIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return api.Car(c, in)
}

// @Accept  json
// @Produce  json
// @Param body body api_models.Body true " "
// @Param query query api_models.Query false " "
// @Param id path int true " "
// @Success 200 {object} Result{data=[]ent.User}
// @Router /user [get]
func Hello(c *gin.Context) (interface{}, error) {
	var err error

	in := &api_models.In{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		return nil, err
	}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return api.Hello(c, in)
}
