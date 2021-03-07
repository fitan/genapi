package api

import (
	"context"
	"ent_samp/ent"
	"ent_samp/models/api_models"
	"ent_samp/public"
	genrest "ent_samp/service"
	"github.com/gin-gonic/gin"
)

// @GenApi /user [post]
func Hello(c *gin.Context, in *api_models.In) ([]*ent.User, error) {
	return public.GetDB().User.Query().All(context.Background())
}

type CarIn struct {
	Query genrest.UserNameEQ
}

// @GenApi /car [get]
func Car(c *gin.Context, in *CarIn) ([]*ent.Car, error) {
	return public.GetDB().Car.Query().All(context.Background())
}

例子