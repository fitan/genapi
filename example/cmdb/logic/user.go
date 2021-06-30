package logic

import (
	"cmdb/ent"
	"github.com/gin-gonic/gin"
)

type UserCallIn struct {
	Query struct {
		Id string `json:"id"`
	}
}

func (i UserCallIn) GetCasbinKeys() [][]interface{} {
	return [][]interface{}{}
}

func (i UserCallIn) GetRedisKey() string {
	return i.Query.Id
}




// @GenApi /api/usercall [get]
// @Casbin url UserCall "呼叫User"
// @Casbin object UserCall ""
// @CallBack redis
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	return nil, nil
}
