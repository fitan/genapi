package casbin

import (
	"cmdb/ent/servicetree"
	"cmdb/public"
	"context"
	"path"
	"strings"
)

type GetRolesIn struct {
}

type GetRolesOut []Role

type Role struct {
	RoleID      string   `json:"role_name"`
	Permissions []string `json:"permissions"`
}

type AddRoleIn struct {
	// 角色
	RoleID string `json:"name"`
	// 允许的方法
	Permissions []string `json:"action"`
}

func (a *AddRoleIn) ToRoles() [][]string {
	roles := make([][]string, 0, 0)
	for _, action := range a.Permissions {
		roles = append(roles, []string{a.RoleID, action})
	}
	return roles
}

type DeleteRolesIn struct {
	Uri struct {
		ID int `json:"id" uri:"id"`
	} `json:"uri"`
}

type RawPolicies [][]string

func (r RawPolicies) ToPolicies() []Policy {
	policies := make([]Policy, 0, 0)
	for _, raw := range r {
		resources := strings.Split(raw[1][1:], "/")
		policies = append(policies, Policy{
			User: raw[0],
			Resource: Resource{
				ProjectId: resources[0],
				ServiceId: resources[1],
			},
			RoleID: raw[2],
		})
	}
	return policies
}

type Policy struct {
	User string `json:"user" binding:"required"`
	Resource
	RoleID string `json:"role" binding:"required"`
}

type Resource struct {
	ProjectId string `json:"project_id" binding:"required"`
	ServiceId string `json:"service_id" binding:"required"`
}

func (r *Resource) IsExist() (bool, error) {
	bg := context.Background()
	return public.GetDB().ServiceTree.Query().Where(servicetree.NameEQ(r.ProjectId)).QueryService().Where(servicetree.NameEQ(r.ServiceId)).Exist(bg)
}

func (r *Resource) EncodeCasbinResource() string {
	return "/" + r.ProjectId + "/" + r.ServiceId
}

type PoliciesIn struct {
	Body []Policy `json:"body"`
}

func (p *PoliciesIn) GetCasbinListKeys() [][]interface{} {
	keys := make([][]interface{}, 0, 0)
	for _, b := range p.Body {
		keys = append(keys, []interface{}{b.EncodeCasbinResource()})
	}
	return keys
}

func (p *PoliciesIn) ToPolicies() [][]string {
	ps := make([][]string, 0, 0)
	for _, v := range p.Body {
		ps = append(ps, []string{v.User, "/" + path.Join(v.ProjectId, v.ServiceId), v.RoleID})
	}
	return ps
}

type GetPoliciesIn struct {
	Query struct {
		Username string `json:"username" form:"username" binding:"required"`
	} `json:"query"`
}
