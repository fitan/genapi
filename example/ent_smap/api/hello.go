package api

import (
	"context"
	"ent_samp/ent"
	"ent_samp/ent/user"
	"ent_samp/genent"
	"ent_samp/models/api_models"
	"ent_samp/public"
	"github.com/gin-gonic/gin"
)

// @GenApi /genapi/car [get]
func Hello(c *gin.Context, in *api_models.In) ([]*ent.User, error) {
	return public.GetDB().User.Query().All(context.Background())
}

type CarIn struct {
	Uri   genent.IdUri
	Query genent.UserNameEQ
}

// @GenApi /genapi/car/{id} [get]
func Car(c *gin.Context, in *CarIn) ([]*ent.User, error) {
	ps, err := genent.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}
	return public.GetDB().User.Query().Where(user.ID(in.Uri.ID)).Where(ps...).All(context.Background())
}
