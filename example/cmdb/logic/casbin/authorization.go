package casbin

import (
	"cmdb/ent"
	"cmdb/gen/entrest"
	"cmdb/public"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/reactivex/rxgo/v2"
)

// @Summary 获取所有角色
// @GenApi /api/roles [get]
func GetListRole(c *gin.Context, in *entrest.GetListRoleBindingIn) (*entrest.GetRoleBindingListData, error) {
	return entrest.GetListRoleBinding(c, in)
}

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
			RoleID:      v[1],
			Permissions: []string{v[0]},
		}, nil
	}).GroupByDynamic(func(item rxgo.Item) string {
		v := item.V.(Role)
		return v.RoleID
	}, rxgo.WithBufferedChannel(len(gs))).DoOnNext(func(i interface{}) {
		g := i.(rxgo.GroupedObservable)
		mergeV, _ := g.Reduce(func(ctx context.Context, i interface{}, i2 interface{}) (interface{}, error) {
			if i == nil {
				return i2, nil
			}
			v1 := i.(Role)
			v2 := i2.(Role)
			return Role{
				RoleID:      v1.RoleID,
				Permissions: append(v1.Permissions, v2.Permissions...),
			}, nil
		}).Get()
		out = append(out, mergeV.V.(Role))
	})
	return out, nil
}

// @Summary 创建角色
// @GenApi /api/role [post]
func AddRole(c *gin.Context, in *entrest.CreateOneRoleBindingIn) (*ent.RoleBinding, error) {
	tx, err := public.GetDB().Tx(context.Background())
	if err != nil {
		return nil, err
	}

	creater := tx.RoleBinding.Create()
	save, err := entrest.GetRoleBindingCURD().CreateOne(creater, &in.Body)
	if err != nil {
		return save, entrest.Rollback(tx, err)
	}

	roles := (&AddRoleIn{
		RoleID:      in.Body.RoleID,
		Permissions: in.Body.Permissions,
	}).ToRoles()

	_, err = public.GetCasbin().AddGroupingPolicies(roles)
	if err != nil {
		return save, entrest.Rollback(tx, err)
	}

	return save, tx.Commit()
}

// @Summary 更新角色
// @GenApi /api/roles [put]
func UpdateRoles(c *gin.Context, in *entrest.UpdateOneRoleBindingIn) (*ent.RoleBinding, error) {
	tx, err := public.GetDB().Tx(context.Background())
	if err != nil {
		return nil, err
	}
	save, err := entrest.GetRoleBindingCURD().UpdateOne(entrest.GetRoleBindingCURD().GetUpdaterById(in.Uri.ID), &in.Body)
	if err != nil {
		return save, entrest.Rollback(tx, err)
	}

	if !in.Body.Status {
		_, err := public.GetCasbin().RemoveFilteredNamedGroupingPolicy("g", 0, in.Body.RoleID)
		if err != nil {
			return save, entrest.Rollback(tx, err)
		}
		return save, err
	}

	roles := (&AddRoleIn{
		RoleID:      in.Body.RoleID,
		Permissions: in.Body.Permissions,
	}).ToRoles()
	_, err = public.GetCasbin().UpdateFilteredNamedPolicies("g", roles, 0, in.Body.RoleID)
	if err != nil {
		return save, entrest.Rollback(tx, err)
	}
	return save, tx.Commit()
}

// @Summary 删除角色
// @GenApi /api/role/{id} [delete]
func DeleteRoles(c *gin.Context, in *DeleteRolesIn) (bool, error) {
	tx, client, err := entrest.GetRoleBindingCURD().GetTx()
	if err != nil {
		return false, err
	}

	data, err := client.Get(context.Background(), in.Uri.ID)

	if err != nil {
		if !ent.IsNotFound(err) {
			return false, entrest.Rollback(tx, err)
		}
	} else {
		err := client.DeleteOneID(in.Uri.ID).Exec(context.Background())
		if err != nil {
			return false, entrest.Rollback(tx, err)
		}
	}

	_, err = public.GetCasbin().RemoveFilteredNamedGroupingPolicy("g", 0, data.RoleID)
	if err != nil {
		return false, entrest.Rollback(tx, err)
	}
	return true, tx.Commit()
}

// @Summary 授权
// @GenApi /api/policies/add [post]
func AddPolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	for _, policy := range in.Body {
		exist, err := policy.IsExist()
		if err != nil {
			return false, err
		}
		if !exist {
			return false, fmt.Errorf("project: %s, service: %s 不存在", policy.ProjectId, policy.ServiceId)
		}
	}
	return public.GetCasbin().AddPolicies(in.ToPolicies())
}

// @Summary 获取授权
// @GenApi /api/policies/get [get]
func GetPolicies(c *gin.Context, in *GetPoliciesIn) ([]Policy, error) {
	raw := public.GetCasbin().GetFilteredPolicy(0, in.Query.Username)
	policies := RawPolicies(raw).ToPolicies()
	var filterP []Policy
	for _, p := range policies {
		if public.GetCasbin().HasGroupingPolicy(p.RoleID, "authorization") {
			filterP = append(filterP, p)
		}
	}

	return filterP, nil

}

// @Summary 删除授权
// @GenApi /api/policies/delete [post]
func DeletePolicies(c *gin.Context, in *PoliciesIn) (bool, error) {
	return public.GetCasbin().RemovePolicies(in.ToPolicies())
}
