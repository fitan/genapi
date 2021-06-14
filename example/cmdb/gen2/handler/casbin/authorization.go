package casbin

import (
	"cmdb/logic/casbin"

	"github.com/gin-gonic/gin"
)

// @Accept  json
// @Produce  json
// @Param query query casbin.Query false " "
// @Success 200 {object} Result{data=bool}
// @Router /policy [delete]
func DeletePolicy(c *gin.Context) (data interface{}, err error) {
	in := &casbin.DeleteIn{}

	err = c.ShouldBindQuery(&in.Query)
	if err != nil {
		return nil, err
	}

	return casbin.DeletePolicy(c, in)
}
