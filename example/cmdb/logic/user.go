package logic

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
	Id int `json:"id" uri:"id"`
	ent.UserQuery

	Uri struct{
		Id int `json:"id"`
	}
	Body struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	Query UserCallQuery `json:"query" genapi:"query"`
}

// @GenApi /api/usercall [get]
func UserCall(c *gin.Context, in *UserCallIn) {
	db := public.GetDB()
	query := db.User.Query()
	ps, err := entt.UserPredicatesExec(in.Query.BindUserNameEQ)
	if err != nil {
		return nil, err
	}

	query.Where(user.And(ps...))
	entt.QueryerIncludes(query, in.Query.Includes)
	//return query.All(context.Background())
	return in.Body,nil
}
