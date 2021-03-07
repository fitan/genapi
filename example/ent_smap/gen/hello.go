package gen

import (
	"ent_samp/api"
	"ent_samp/models/api_models"
	genrest "ent_samp/service"

	"github.com/gin-gonic/gin"
)

func Car(c *gin.Context) {
	var err error

	in := &api.CarIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		genrest.RestReturnFunc(c, "", err)
		return
	}

	data, err := api.Car(c, in)
	genrest.RestReturnFunc(c, data, err)
}

func Hello(c *gin.Context) {
	var err error

	in := &api_models.In{}

	err = c.ShouldBindJSON(&in.Body)
	if err != nil {
		genrest.RestReturnFunc(c, "", err)
		return
	}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		genrest.RestReturnFunc(c, "", err)
		return
	}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		genrest.RestReturnFunc(c, "", err)
		return
	}

	data, err := api.Hello(c, in)
	genrest.RestReturnFunc(c, data, err)
}
