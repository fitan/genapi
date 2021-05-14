package casbin

import (
	"cmdb/public"
	"github.com/gin-gonic/gin"
)

const UrlPrefix = "/api"

// @GenApi /policy [get]
func GetPolicyList(c *gin.Context, in *GetListIn) ([][]string,error) {
	return public.GetCasbin().GetFilteredPolicy(0,in.Query.User), nil
}

// @GenApi /policy [post]
func AddPolicies(c *gin.Context, in *AddListIn) (bool, error) {
	ps := make([][]string, 0, len(in.Body))
	for _, v := range in.Body {
		ps = append(ps, append([]string{}, v.User,v.Path,v.Method))
	}
	return public.GetCasbin().AddPolicies(ps)
}

// @GenApi /policy [put]
func UpdatePolicy(c *gin.Context, in *UpdateIn) (bool, error)  {
	return public.GetCasbin().UpdatePolicies(in.Body.Old.Serialize(), in.Body.New.Serialize())
}

// @GenApi /policy [delete]
func DeletePolicy(c *gin.Context, in *DeleteIn) (bool, error) {
	return public.GetCasbin().DeleteUser(in.Query.User)
}





