package casbin

import (
	"cmdb/public"
	"github.com/gin-gonic/gin"
)

type GetRolesIn struct {

}

type GetRolesOut map[string][]string

func GetRoles(c *gin.Context, in *GetRolesIn) map[string][]string {
	out := make(map[string][]string, 0)
	gs := public.GetCasbin().GetNamedGroupingPolicy("g")
	for _, g := range gs {
		Name := g[1]
		if l, ok := out[Name]; ok {
			out[Name] = append(l, g[0])
		} else {
			out[Name] = []string{g[0]}
		}
	}
	return out
}


type AddRoleIn struct {
	Body struct{
		// 角色
		Name string
		// 允许的方法
		Action []string
	}
}

func (a *AddRoleIn) ToRoles() [][]string {
	roles := make([][]string, 0 ,0)
	for _, action := range a.Body.Action {
		roles = append(roles, []string{action ,a.Body.Name})
	}
	return roles
}



func AddRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().AddNamedGroupingPolicies("g",in.ToRoles())
}

func UpdateRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().UpdateFilteredNamedPolicies("g", in.ToRoles(), 1, in.Body.Name)
}

type DeleteRolesIn struct {
	Uri struct{
		Name string
	}
}
func DeleteRoles(c *gin.Context, in *DeleteRolesIn) (bool, error) {
	return public.GetCasbin().RemoveFilteredGroupingPolicy(1, in.Uri.Name)
}







