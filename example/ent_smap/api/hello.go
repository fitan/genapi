package api

import (
	"context"
	"ent_samp/ent"
	"ent_samp/models/api_models"
	"ent_samp/public"
	genrest "ent_samp/service"
	"github.com/gin-gonic/gin"
)

// @GenApi /user [get]
func Hello(c *gin.Context, in *api_models.In) ([]*ent.User, error) {
	ps, err := genrest.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}
	return public.GetDB().User.Query().Where(ps...).All(context.Background())
}

type CarIn struct {
	Query genrest.UserNameEQ
}

// @GenApi /car [get]
func Car(c *gin.Context, in *CarIn) ([]*ent.User, error) {
	return public.GetDB().User.Query().All(context.Background())
}
