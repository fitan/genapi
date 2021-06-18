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
	//MUri map[string]struct{
	//	Id int
	//}
	Uri struct{
		// 查询id
		// 多次查询id 的结果
		Id int `json:"id" uri:"id"`
	}
	// 这是body
	Body struct{
		Name string `json:"name"`
		Age int `json:"age"`
	}
	// 这是query
	Query UserCallQuery `json:"query" genapi:"query"`
	Header public.Header
}

func GetId()  {
	
}




// @Tag fsdf
// @GenApi /api/usercall [get]
// @Casbin UserCall "呼叫User"
// @Redis
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
	//return in.Body,nil
}
