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




// @GenApi /api/usercall [get]
// @Casbin UserCall "呼叫User"
// @Redis
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	return nil, nil
}
