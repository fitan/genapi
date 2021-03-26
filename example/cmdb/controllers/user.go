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
	Uri   entt.IdUri
}

// @GenApi /api/usercall/{id} [get]
func UserCall(c *gin.Context, in *UserCallIn) (*ent.User, error) {
	qureyer := public.DB.User.Query()
	entt.QueryerIncludes(qureyer, in.Query.Includes)
	return qureyer.Where(user.IDEQ(in.Uri.ID)).Only(context.Background())
}
