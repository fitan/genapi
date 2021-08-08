package casbin

import (
	"cmdb/public"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reactivex/rxgo/v2"
	"path"
	"strings"
)

type GetRolesIn struct {
}

type GetRolesOut []Role

type Role struct {
	RoleName string   `json:"role_name"`
	Methods  []string `json:"methods"`
}

// @Summary 获取所有角色
// @GenApi /api/roles [get]
func GetRoles(c *gin.Context, in *GetRolesIn) ([]Role, error) {
	gs := public.GetCasbin().GetNamedGroupingPolicy("g")
	out := make(GetRolesOut, 0, 0)
	observable := rxgo.Create([]rxgo.Producer{func(ctx context.Context, next chan<- rxgo.Item) {
		for index, _ := range gs {
			next <- rxgo.Of(gs[index])
		}
	}})
	<-observable.Map(func(ctx context.Context, i interface{}) (interface{}, error) {
		v := i.([]string)
		return Role{
			RoleName: v[1],
			Methods:  []string{v[0]},
		}, nil
	}).GroupByDynamic(func(item rxgo.Item) string {
		v := item.V.(Role)
		return v.RoleName
	}, rxgo.WithBufferedChannel(len(gs))).DoOnNext(func(i interface{}) {
		g := i.(rxgo.GroupedObservable)
		mergeV, _ := g.Reduce(func(ctx context.Context, i interface{}, i2 interface{}) (interface{}, error) {
			if i == nil {
				return i2, nil
			}
			v1 := i.(Role)
			v2 := i2.(Role)
			return Role{
				RoleName: v1.RoleName,
				Methods:  append(v1.Methods, v2.Methods...),
			}, nil
		}).Get()
		out = append(out, mergeV.V.(Role))
	})
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

// @Summary 创建角色
// @GenApi /api/roles [post]
func AddRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().AddNamedGroupingPolicies("g", in.ToRoles())
}

// @Summary 更新角色
// @GenApi /api/roles [put]
func UpdateRoles(c *gin.Context, in *AddRoleIn) (bool, error) {
	return public.GetCasbin().UpdateFilteredNamedPolicies("g", in.ToRoles(), 1, in.Body.Name)
}

type DeleteRolesIn struct {
	Uri struct {
		Name string `json:"name" uri:"name"`
	} `json:"uri"`
}

// @Summary 删除角色
// @GenApi /api/role/{name} [delete]
func DeleteRoles(c *gin.Context, in *DeleteRolesIn) (bool, error) {
	return public.GetCasbin().RemoveFilteredGroupingPolicy(1, in.Uri.Name)
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
			Role: raw[2],
		})
	}
	return policies
}

type Policy struct {
	User string `json:"user" binding:"required"`
	Resource
	Role string `json:"role" binding:"required"`
}

type Resource struct {
	ProjectId string `json:"project_id" binding:"required"`
	ServiceId string `json:"service_id" binding:"required"`
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
		ps = append(ps, []string{v.User, "/" + path.Join(v.ProjectId, v.ServiceId), v.Role})
	}
	return ps
}

// @Summary 授权
// @GenApi /api/policies/add [post]
func AddPolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	return public.GetCasbin().AddPolicies(in.ToPolicies())
}

type GetPoliciesIn struct {
	Query struct {
		Username string `json:"username" form:"user_name" binding:"required"`
	} `json:"query"`
}

// @Summary 获取授权
// @GenApi /api/policies/get [get]
func GetPolicies(c *gin.Context, in *GetPoliciesIn) ([]Policy, error) {
	//userName, _ := public.GetUserNameByContext(c)
	fmt.Println(in.Query.Username)
	raw := public.GetCasbin().GetFilteredPolicy(0, in.Query.Username)

	return RawPolicies(raw).ToPolicies(), nil

}

// @Summary 删除授权
// @GenApi /api/policies/delete [post]
func DeletePolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	return public.GetCasbin().RemovePolicies(in.ToPolicies())
}
