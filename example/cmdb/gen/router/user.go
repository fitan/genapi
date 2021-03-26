package router

import (
	"cmdb/controllers"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param query query controllers.UserCallQuery false " "
// @Param id path int true " "
// @Success 200 {object} Result{data=ent.User}
// @Router /api/usercall/{id} [get]
func UserCall(c *gin.Context) (interface{}, error) {
	var err error

	in := &controllers.UserCallIn{}

	err = c.ShouldBindUri(&in.Uri)
	if err != nil {
		return nil, err
	}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return controllers.UserCall(c, in)
}
