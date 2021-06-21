package handler

import (
	"cmdb/logic"
	"cmdb/public"

	"github.com/gin-gonic/gin"
)

type SwagUserCallQuery struct {
	Id string `json:"id"`
}

// @Accept  json
// @Produce  json
// @Param query query SwagUserCallQuery false " "
// @Success 200 {object} Result{data=[]ent.User}
// @Router /api/usercall [get]
func UserCall(c *gin.Context) (data interface{}, err error) {
	in := &logic.UserCallIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	data, err = public.CheckKeysCasbin(c, "UserCall", in.GetCasbinKeys())
	if err != nil {
		return data, err
	}

	return logic.UserCall(c, in)
}
