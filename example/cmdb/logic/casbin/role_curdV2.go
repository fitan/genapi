package casbin

import (
	"cmdb/ent"
	"cmdb/ent/rolebinding"
	"cmdb/gen/entt"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoleCURDV2 struct {
	Db *ent.Client

	ProjectObj *entt.ProjectCURD

	ServiceObj *entt.ServiceCURD

	RoleBindingObj *entt.RoleBindingCURD
}

func (curd *RoleCURDV2) GetQueryer() *ent.RoleBindingQuery {
	return curd.Db.RoleBinding.Query()
}

func (curd *RoleCURDV2) GetCreater() *ent.RoleBindingCreate {
	return curd.Db.RoleBinding.Create()
}

func (curd *RoleCURDV2) GetDeleter() *ent.RoleBindingDelete {
	return curd.Db.RoleBinding.Delete()
}

func (curd *RoleCURDV2) GetUpdater(id int) *ent.RoleBindingUpdateOne {
	return curd.Db.RoleBinding.UpdateOneID(id)
}

func (curd *RoleCURDV2) GetOneQueryer(id int) *ent.RoleBindingQuery {
	return curd.GetQueryer().Where(rolebinding.IDEQ(id))
}

func (curd *RoleCURDV2) GetOne(id int) (*ent.RoleBinding, error) {
	return curd.GetOneQueryer(id).Only(context.Background())
}

func (curd *RoleCURDV2) SetListQuery(queryer *ent.RoleBindingQuery, query *entt.RoleBindingDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}
	entt.RoleBindingSelete(queryer)
	curd.DefaultOrder(queryer)
	return nil
}

func (curd *RoleCURDV2) SetListCountQuery(queryer *ent.RoleBindingQuery, query *entt.RoleBindingDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(rolebinding.And(ps...))
	return nil
}

func (curd *RoleCURDV2) GetList(query *ent.RoleBindingQuery) (*entt.GetRoleBindingListData, error) {
	listQueryer := curd.GetQueryer()
	countQueryer := curd.GetQueryer()
	return curd.GetListByQueryer(listQueryer, countQueryer, query)
}

func (curd *RoleCURDV2) GetListByQueryer(listQueryer, countQueryer *ent.RoleBindingQuery, query *entt.RoleBindingDefaultQuery) (*entt.GetRoleBindingListData, error) {
	bg := context.Background()
	err := curd.SetListQuery(listQueryer, query)
	if err != nil {
		return nil, err
	}
	list, err := listQueryer.All(bg)
	if err != nil {
		return nil, err
	}
	err = curd.SetListCountQuery(countQueryer, query)
	if err != nil {
		return nil, err
	}
	count, err := countQueryer.Count(bg)
	if err != nil {
		return nil, err
	}
	return &entt.GetRoleBindingListData{count, list}, nil
}

func (curd *RoleCURDV2) optionalOrder(c *gin.Context, queryer *ent.RoleBindingQuery) error {
	var expect = map[string]int{}
	orderFunc, err := entt.BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *RoleCURDV2) DefaultOrder(queryer *ent.RoleBindingQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *RoleCURDV2) CreateOne(rolebinding *ent.RoleBinding) (*ent.RoleBinding, error) {
	creater := curd.GetCreater()
	entt.RoleBindingCreateMutation(creater.Mutation(), rolebinding)
	return creater.Save(context.Background())
}

func (curd *RoleCURDV2) GetBulk(rolebindings []*ent.RoleBinding) []*ent.RoleBindingCreate {
	bulk := make([]*ent.RoleBindingCreate, 0, len(rolebindings))
	for _, v := range rolebindings {
		creater := curd.GetCreater()
		entt.RoleBindingCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *RoleCURDV2) CreateList(rolebindings []*ent.RoleBinding) ([]*ent.RoleBinding, error) {
	return curd.Db.RoleBinding.CreateBulk(curd.GetBulk(rolebindings)...).Save(context.Background())
}

func (curd *RoleCURDV2) UpdateOne(rolebinding *ent.RoleBinding) (*ent.RoleBinding, error) {
	updater := curd.GetUpdater(rolebinding.ID)
	entt.RoleBindingUpdateMutation(updater.Mutation(), rolebinding)
	return updater.Save(context.Background())
}

func (curd *RoleCURDV2) UpdateList(rolebindings []*ent.RoleBinding) (string, error) {
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return "", err
	}
	for _, v := range rolebindings {
		updater := tx.RoleBinding.UpdateOneID(v.ID)
		entt.RoleBindingUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(bg)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return "", fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return "ok", tx.Commit()
}

func (curd *RoleCURDV2) DeleteOne(id int) (int, error) {
	return curd.GetDeleter().Where(rolebinding.IDEQ(id)).Exec(context.Background())
}

func (curd *RoleCURDV2) DeleteList(ids []int) (int, error) {
	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDIn(ids...)).Exec(context.Background())
}
