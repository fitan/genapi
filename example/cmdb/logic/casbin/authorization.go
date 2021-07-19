package casbin

import (
	"cmdb/public"
	"github.com/gin-gonic/gin"
)

type GetRolesIn struct {
}

type GetRolesOut map[string][]string

// @GenApi /api/roles [get]
func GetRoles(c *gin.Context, in *GetRolesIn) (map[string][]string, error) {
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
	return out, nil
}

type AddRoleIn struct {
	Body struct {
		// 角色
		Name string `json:"name"`
		// 允许的方法
		Action []string `json:"action"`
	} `json:"body"`
}

func (a *AddRoleIn) ToRoles() [][]string {
	roles := make([][]string, 0, 0)
	for _, action := range a.Body.Action {
		roles = append(roles, []string{action, a.Body.Name})
	}
	return roles
}

// @GenApi /api/roles [post]
func AddRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().AddNamedGroupingPolicies("g", in.ToRoles())
}

// @GenApi /api/roles [put]
func UpdateRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().UpdateFilteredNamedPolicies("g", in.ToRoles(), 1, in.Body.Name)
}

type DeleteRolesIn struct {
	Uri struct {
		Name string `json:"name" uri:"name"`
	} `json:"uri"`
}

// @GenApi /api/role/{name} [delete]
func DeleteRoles(c *gin.Context, in *DeleteRolesIn) (bool, error) {
	return public.GetCasbin().RemoveFilteredGroupingPolicy(1, in.Uri.Name)
}

type PoliciesIn struct {
	Body []struct {
		User      string `json:"user"`
		ServiceId string `json:"service_id"`
		Role      string `json:"role"`
	} `json:"body"`
}

func (p *PoliciesIn) ToPolicies() [][]string {
	ps := make([][]string, 0, 0)
	for _, v := range p.Body {
		ps = append(ps, []string{v.User, v.ServiceId, v.Role})
	}
	return ps
}

// @GenApi /api/policies/add [post]
func AddPolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	return public.GetCasbin().AddPolicies(in.ToPolicies())
}

type GetPoliciesIn struct {
}

// @GenApi /api/policies/get [get]
func GetPolicies(c *gin.Context, in *GetPoliciesIn) ([][]string, error) {
	return public.GetCasbin().GetPolicy(), nil
}

// @GenApi /api/policies/delete [post]
func DeletePolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	return public.GetCasbin().RemovePolicies(in.ToPolicies())
}
