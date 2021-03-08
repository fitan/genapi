package gen

import (
	"ent_samp/api"
	"ent_samp/models/api_models"

	"github.com/gin-gonic/gin"
)

func Car(c *gin.Context) (interface{}, error) {
	var err error

	in := &api.CarIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return api.Car(c, in)
}

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
