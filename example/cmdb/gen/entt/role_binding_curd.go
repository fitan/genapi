package entt

import (
	"cmdb/ent"
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/service"
	"cmdb/ent/user"
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

		r.POST(curd.CreateOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
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

		r.POST(curd.CreateOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneProjectByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneProjectByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneServiceByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneServiceByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneUserByRoleBindingIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneUserByRoleBindingId(c)
			RestReturnFunc(c, data, err)
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

func (curd *RoleBindingCURD) BaseGetOneQueryer(c *gin.Context) (*ent.RoleBindingQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.RoleBinding.Query().Where(rolebinding.IDEQ(id.ID)), nil
}

func (curd *RoleBindingCURD) GetOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) GetOne(c *gin.Context) (*ent.RoleBinding, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	RoleBindingSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *RoleBindingCURD) defaultGetListCount(queryer *ent.RoleBindingQuery, query *RoleBindingDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(rolebinding.And(ps...))
	return nil
}

func (curd *RoleBindingCURD) defaultGetListQueryer(queryer *ent.RoleBindingQuery, query *RoleBindingDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	RoleBindingSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *RoleBindingCURD) BaseGetListQueryer(c *gin.Context) (*ent.RoleBindingQuery, *ent.RoleBindingQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.RoleBinding.Query()

	err = curd.defaultGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.RoleBinding.Query()
	err = curd.defaultGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *RoleBindingCURD) GetListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) GetList(c *gin.Context) (*GetRoleBindingListData, error) {
	getListQueryer, countQueryer, err := curd.BaseGetListQueryer(c)
	if err != nil {
		return nil, err
	}

	bg := context.Background()
	count, err := countQueryer.Count(bg)
	if err != nil {
		return nil, err
	}

	res, err := getListQueryer.All(bg)
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

func (curd *RoleBindingCURD) BaseCreateOneCreater(c *gin.Context) (*ent.RoleBindingCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	creater := curd.Db.RoleBinding.Create()
	RoleBindingCreateMutation(creater.Mutation(), body)
	return creater, nil
}

func (curd *RoleBindingCURD) CreateOneRoutePath() string {
	return "/role_binding"
}

func (curd *RoleBindingCURD) CreateOne(c *gin.Context) (*ent.RoleBinding, error) {
	creater, err := curd.BaseCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *RoleBindingCURD) BaseCreateListBulk(c *gin.Context) ([]*ent.RoleBindingCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	bulk := make([]*ent.RoleBindingCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.RoleBinding.Create()
		RoleBindingCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk, nil
}

func (curd *RoleBindingCURD) BaseCreateList(c *gin.Context) (*ent.RoleBindingCreateBulk, error) {
	bulk, err := curd.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.RoleBinding.CreateBulk(bulk...), nil
}

func (curd *RoleBindingCURD) CreateListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) CreateList(c *gin.Context) ([]*ent.RoleBinding, error) {
	creater, err := curd.BaseCreateList(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *RoleBindingCURD) BaseUpdateOneUpdater(c *gin.Context) (*ent.RoleBindingUpdateOne, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	updater := curd.Db.RoleBinding.UpdateOneID(id.ID)
	RoleBindingUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *RoleBindingCURD) UpdateOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) UpdateOne(c *gin.Context) (*ent.RoleBinding, error) {
	updater, err := curd.BaseUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *RoleBindingCURD) BaseUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
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

func (curd *RoleBindingCURD) UpdateListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.BaseUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *RoleBindingCURD) BaseDeleteOneDeleter(c *gin.Context) (*ent.RoleBindingDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDEQ(id.ID)), nil
}

func (curd *RoleBindingCURD) DeleteOneRoutePath() string {
	return "/role_binding/:id"
}

func (curd *RoleBindingCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *RoleBindingCURD) BaseDeleteListDeleter(c *gin.Context) (*ent.RoleBindingDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDIn(ids.Ids...)), nil
}

func (curd *RoleBindingCURD) DeleteListRoutePath() string {
	return "/role_bindings"
}

func (curd *RoleBindingCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *RoleBindingCURD) GetOneProjectByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/project"
}

func (curd *RoleBindingCURD) GetOneProjectByRoleBindingId(c *gin.Context) (*ent.Project, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	return queryer.QueryProject().First(context.Background())
}

// M2O
func (curd *RoleBindingCURD) CreateOneProjectByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/project"
}

func (curd *RoleBindingCURD) CreateOneProjectByRoleBindingId(c *gin.Context) (*ent.Project, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}

	projectCreater, err := curd.ProjectObj.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return projectCreater.AddRoleBindingIDs(id.ID).Save(context.Background())
}

func (curd *RoleBindingCURD) DeleteOneProjectByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/project"
}

func (curd *RoleBindingCURD) DeleteOneProjectByRoleBindingId(c *gin.Context) (int, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	id, err := queryer.QueryProject().OnlyID(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Project.Delete().Where(project.IDEQ(id)).Exec(context.Background())
}

func (curd *RoleBindingCURD) GetOneServiceByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/service"
}

func (curd *RoleBindingCURD) GetOneServiceByRoleBindingId(c *gin.Context) (*ent.Service, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	return queryer.QueryService().First(context.Background())
}

// M2O
func (curd *RoleBindingCURD) CreateOneServiceByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/service"
}

func (curd *RoleBindingCURD) CreateOneServiceByRoleBindingId(c *gin.Context) (*ent.Service, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}

	serviceCreater, err := curd.ServiceObj.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return serviceCreater.AddRoleBindingIDs(id.ID).Save(context.Background())
}

func (curd *RoleBindingCURD) DeleteOneServiceByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/service"
}

func (curd *RoleBindingCURD) DeleteOneServiceByRoleBindingId(c *gin.Context) (int, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	id, err := queryer.QueryService().OnlyID(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Service.Delete().Where(service.IDEQ(id)).Exec(context.Background())
}

func (curd *RoleBindingCURD) GetOneUserByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/user"
}

func (curd *RoleBindingCURD) GetOneUserByRoleBindingId(c *gin.Context) (*ent.User, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	return queryer.QueryUser().First(context.Background())
}

// M2O
func (curd *RoleBindingCURD) CreateOneUserByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/user"
}

func (curd *RoleBindingCURD) CreateOneUserByRoleBindingId(c *gin.Context) (*ent.User, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}

	userCreater, err := curd.UserObj.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return userCreater.AddRoleBindingIDs(id.ID).Save(context.Background())
}

func (curd *RoleBindingCURD) DeleteOneUserByRoleBindingIdRoutePath() string {
	return "/role_binding/:id/user"
}

func (curd *RoleBindingCURD) DeleteOneUserByRoleBindingId(c *gin.Context) (int, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	id, err := queryer.QueryUser().OnlyID(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.User.Delete().Where(user.IDEQ(id)).Exec(context.Background())
}
