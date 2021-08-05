package user

import (
	"cmdb/ent"
	"github.com/gin-gonic/gin"
)

type UserCallIn struct {
	Query struct {
		Id string `json:"id"`
	} `json:"query"`
}


func (i UserCallIn) GetCasbinKeys() [][]interface{} {
	k := make([][]interface{},0,0)
	k = append(k, []interface{}{i.Query.Id})
	return k
}

func (i UserCallIn) GetRedisKey() string {
	return i.Query.Id
}




// @GenApi /api/usercall [get]
// @CallBack redis get
// @Casbin UserCall 呼叫usercall
func UserCall(c *gin.Context, in *UserCallIn) ([]*ent.User, error) {
	return nil, nil
}
