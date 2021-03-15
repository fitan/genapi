package entt

import (
	"cmdb/ent"
	"cmdb/ent/rolebinding"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type RoleBindingCURD struct {
	Db *ent.Client

	ProjectObj *ProjectCURD

	ServiceObj *ServiceCURD

	UserObj *UserCURD
}

func (curd *RoleBindingCURD) RegisterRouter(router interface{}) {
	switch router.(type) {
	case *gin.Engine:
		r := router.(*gin.Engine)

		r.POST(curd.CreateOneRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOne(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateList(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOne(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteList(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOne(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoutePath(), func(c *gin.Context) {
			data, err := curd.GetList(c)
			RestReturnFunc(c, data, err)
		})

		r.PUT(curd.UpdateOneRoutePath(), func(c *gin.Context) {
			data, err := curd.UpdateOne(c)
			RestReturnFunc(c, data, err)
		})

		r.PUT(curd.UpdateListRoutePath(), func(c *gin.Context) {
			err := curd.UpdateList(c)
			RestReturnFunc(c, "", err)
		})

	case *gin.RouterGroup:
		r := router.(*gin.RouterGroup)

		r.POST(curd.CreateOneRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOne(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateList(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOne(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteList(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOne(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoutePath(), func(c *gin.Context) {
			data, err := curd.GetList(c)
			RestReturnFunc(c, data, err)
		})

		r.PUT(curd.UpdateOneRoutePath(), func(c *gin.Context) {
			data, err := curd.UpdateOne(c)
			RestReturnFunc(c, data, err)
		})

		r.PUT(curd.UpdateListRoutePath(), func(c *gin.Context) {
			err := curd.UpdateList(c)
			RestReturnFunc(c, "", err)
		})

	}
}

func (curd *RoleBindingCURD) BindObj(c *gin.Context) (*ent.RoleBinding, error) {
	body := new(ent.RoleBinding)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *RoleBindingCURD) BindObjs(c *gin.Context) (ent.RoleBindings, error) {
	body := make(ent.RoleBindings, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *RoleBindingCURD) BindDefaultQuery(c *gin.Context) (*RoleBindingDefaultQuery, error) {
	body := new(RoleBindingDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *RoleBindingCURD) GetIDs(role_bindings ent.RoleBindings) []int {
	IDs := make([]int, 0, len(role_bindings))
	for _, role_binding := range role_bindings {
		IDs = append(IDs, role_binding.ID)
	}
	return IDs
}

func (curd *RoleBindingCURD) BaseGetOneQueryer(id int) (*ent.RoleBindingQuery, error) {
	return curd.Db.RoleBinding.Query().Where(rolebinding.IDEQ(id)), nil
}

func (curd *RoleBindingCURD) defaultGetOneQueryer(c *gin.Context) (*ent.RoleBindingQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *RoleBindingCURD) GetOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) GetOne(c *gin.Context) (*ent.RoleBinding, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	RoleBindingSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *RoleBindingCURD) BaseGetListCount(queryer *ent.RoleBindingQuery, query *RoleBindingDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(rolebinding.And(ps...))
	return nil
}

func (curd *RoleBindingCURD) BaseGetListQueryer(queryer *ent.RoleBindingQuery, query *RoleBindingDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	RoleBindingSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *RoleBindingCURD) defaultGetListQueryer(c *gin.Context) (*ent.RoleBindingQuery, *ent.RoleBindingQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.RoleBinding.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.RoleBinding.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *RoleBindingCURD) GetListRoutePath() string {
	return "/role_bindings"
}

type GetRoleBindingListData struct {
	Count  int
	Result []*ent.RoleBinding
}

func (curd *RoleBindingCURD) GetList(c *gin.Context) (*GetRoleBindingListData, error) {
	getListQueryer, countQueryer, err := curd.defaultGetListQueryer(c)
	if err != nil {
		return nil, err
	}

	count, err := countQueryer.Count(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := getListQueryer.All(context.Background())
	if err != nil {
		return nil, err
	}

	return &GetRoleBindingListData{count, res}, nil
}

func (curd *RoleBindingCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *RoleBindingCURD) defaultOrder(queryer *ent.RoleBindingQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *RoleBindingCURD) BaseCreateOneCreater(body *ent.RoleBinding) *ent.RoleBindingCreate {
	creater := curd.Db.RoleBinding.Create()
	RoleBindingCreateMutation(creater.Mutation(), body)
	return creater
}

func (curd *RoleBindingCURD) defaultCreateOneCreater(c *gin.Context) (*ent.RoleBindingCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *RoleBindingCURD) CreateOneRoutePath() string {
	return "/role_binding"
}

func (curd *RoleBindingCURD) CreateOne(c *gin.Context) (*ent.RoleBinding, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *RoleBindingCURD) BaseCreateListBulk(body ent.RoleBindings) []*ent.RoleBindingCreate {
	bulk := make([]*ent.RoleBindingCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.RoleBinding.Create()
		RoleBindingCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *RoleBindingCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.RoleBindingCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *RoleBindingCURD) CreateListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) CreateList(c *gin.Context) ([]*ent.RoleBinding, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.RoleBinding.CreateBulk(bulk...).Save(context.Background())
}

func (curd *RoleBindingCURD) BaseUpdateOneUpdater(id int, body *ent.RoleBinding) (*ent.RoleBindingUpdateOne, error) {
	updater := curd.Db.RoleBinding.UpdateOneID(id)
	RoleBindingUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *RoleBindingCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.RoleBindingUpdateOne, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseUpdateOneUpdater(id.ID, body)
}

func (curd *RoleBindingCURD) UpdateOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) UpdateOne(c *gin.Context) (*ent.RoleBinding, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *RoleBindingCURD) BaseUpdateListUpdater(body ent.RoleBindings) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.RoleBinding.UpdateOneID(v.ID)
		RoleBindingUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *RoleBindingCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *RoleBindingCURD) UpdateListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *RoleBindingCURD) BaseDeleteOneDeleter(id int) *ent.RoleBindingDelete {
	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDEQ(id))
}

func (curd *RoleBindingCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.RoleBindingDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *RoleBindingCURD) DeleteOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *RoleBindingCURD) BaseDeleteListDeleter(ids []int) *ent.RoleBindingDelete {
	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDIn(ids...))
}

func (curd *RoleBindingCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.RoleBindingDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *RoleBindingCURD) DeleteListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}
