package controllers

import (
	"context"
	"ent_samp/ent"
	"ent_samp/gen/entt"
	"ent_samp/models/api_models"
	"ent_samp/public"
	"github.com/gin-gonic/gin"
)

// @GenApi /genapi/user [get]
func User(c *gin.Context, in *api_models.In) (*api_models.UserOut, error) {
	ps, err := entt.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}
	all, err := public.GetDB().User.Query().Where(ps...).All(context.Background())
	if err != nil {
		return nil, err
	}
	return &api_models.UserOut{
		User: all,
		Len:  len(all),
	}, err
}

type CarIn struct {
	Query entt.UserNameEQ
}

// @GenApi /genapi/car [get]
func Car(c *gin.Context, in *CarIn) ([]*ent.User, error) {
	ps, err := entt.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}
	public.GetDB().User.Update().RemoveCarIDs()
	return public.GetDB().User.Query().Where(ps...).All(context.Background())
}
