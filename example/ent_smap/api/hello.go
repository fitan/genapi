package api

import (
	"context"
	"ent_samp/ent"
	"ent_samp/ent/user"
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
	Uri   genrest.IdUri
	Query genrest.UserNameEQ
}

// @GenApi /genapi/car/{id} [get]
func Car(c *gin.Context, in *CarIn) ([]*ent.User, error) {
	ps, err := genrest.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}
	return public.GetDB().User.Query().Where(user.ID(in.Uri.ID)).Where(ps...).All(context.Background())
}
