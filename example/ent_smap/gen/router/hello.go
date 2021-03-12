package router

import (
	"ent_samp/controllers"
	"ent_samp/models/api_models"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param query query entt.UserNameEQ false " "
// @Success 200 {object} Result{data=[]ent.User}
// @Router /genapi/car [get]
func Car(c *gin.Context) (interface{}, error) {
	var err error

	in := &controllers.CarIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return controllers.Car(c, in)
}

// @Accept  json
// @Produce  json
// @Param query query api_models.Query false " "
// @Success 200 {object} Result{data=api_models.UserOut}
// @Router /genapi/user [get]
func User(c *gin.Context) (interface{}, error) {
	var err error

	in := &api_models.In{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return controllers.User(c, in)
}
