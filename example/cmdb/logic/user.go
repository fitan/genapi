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

func (UserCallIn) GetCasbinKeys() []interface{} {
	return []interface{}{}
}

type CasbinPluginer1 interface {
	GetCasbinKeys() []interface{}
}

type CasbinPluginer2 interface {
	GetCasbinKeys() [][]interface{}
}

// @GenApi /api/usercall [get]
// @Casbin UserCall "呼叫User"
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	return nil, nil
}
