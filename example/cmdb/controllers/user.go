package controllers

import (
	"cmdb/ent"
	"cmdb/ent/user"
	"cmdb/gen/entt"
	"cmdb/public"
	"context"
	"github.com/gin-gonic/gin"
)

type UserCallQuery struct {
	entt.UserIncludes
	entt.UserNameEQ
}

type UserCallIn struct {
	Query UserCallQuery
}

// @GenApi /api/usercall [get]
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	db := public.GetDB()
	query := db.User.Query()
	ps, err := entt.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}

	query.Where(user.And(ps...))
	entt.QueryerIncludes(query, in.Query.Includes)
	return query.All(context.Background())

}
