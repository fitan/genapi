package entt

import (
	"cmdb/ent"
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/server"
	"cmdb/ent/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServiceCURD struct {
	Db *ent.Client

	RoleBindingObj *RoleBindingCURD

	ServerObj *ServerCURD

	ProjectObj *ProjectCURD
}

func (curd *ServiceCURD) RegisterRouter(router interface{}) {
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

		r.POST(curd.CreateListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneProjectByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneProjectByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneProjectByServiceId(c)
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

		r.POST(curd.CreateListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServersByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServersByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateOneProjectByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteOneProjectByServiceId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetOneProjectByServiceIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOneProjectByServiceId(c)
			RestReturnFunc(c, data, err)
		})

	}
}

func (curd *ServiceCURD) BindObj(c *gin.Context) (*ent.Service, error) {
	body := new(ent.Service)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *ServiceCURD) BindObjs(c *gin.Context) (ent.Services, error) {
	body := make(ent.Services, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *ServiceCURD) BindDefaultQuery(c *gin.Context) (*ServiceDefaultQuery, error) {
	body := new(ServiceDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *ServiceCURD) BaseGetOneQueryer(c *gin.Context) (*ent.ServiceQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Service.Query().Where(service.IDEQ(id.ID)), nil
}

func (curd *ServiceCURD) GetOneRoutePath() string {
	return "/service/:id"
}

func (curd *ServiceCURD) GetOne(c *gin.Context) (*ent.Service, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	ServiceSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *ServiceCURD) defaultGetListCount(queryer *ent.ServiceQuery, query *ServiceDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(service.And(ps...))
	return nil
}

func (curd *ServiceCURD) defaultGetListQueryer(queryer *ent.ServiceQuery, query *ServiceDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	ServiceSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *ServiceCURD) BaseGetListQueryer(c *gin.Context) (*ent.ServiceQuery, *ent.ServiceQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.Service.Query()

	err = curd.defaultGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.Service.Query()
	err = curd.defaultGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *ServiceCURD) GetListRoutePath() string {
	return "/services"
}

func (curd *ServiceCURD) GetList(c *gin.Context) (*GetServiceListData, error) {
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

	return &GetServiceListData{count, res}, nil
}

func (curd *ServiceCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *ServiceCURD) defaultOrder(queryer *ent.ServiceQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *ServiceCURD) BaseCreateOneCreater(c *gin.Context) (*ent.ServiceCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	creater := curd.Db.Service.Create()
	ServiceCreateMutation(creater.Mutation(), body)
	return creater, nil
}

func (curd *ServiceCURD) CreateOneRoutePath() string {
	return "/service"
}

func (curd *ServiceCURD) CreateOne(c *gin.Context) (*ent.Service, error) {
	creater, err := curd.BaseCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *ServiceCURD) BaseCreateListBulk(c *gin.Context) ([]*ent.ServiceCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	bulk := make([]*ent.ServiceCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.Service.Create()
		ServiceCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk, nil
}

func (curd *ServiceCURD) BaseCreateList(c *gin.Context) (*ent.ServiceCreateBulk, error) {
	bulk, err := curd.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Service.CreateBulk(bulk...), nil
}

func (curd *ServiceCURD) CreateListRoutePath() string {
	return "/services"
}

func (curd *ServiceCURD) CreateList(c *gin.Context) ([]*ent.Service, error) {
	creater, err := curd.BaseCreateList(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *ServiceCURD) BaseUpdateOneUpdater(c *gin.Context) (*ent.ServiceUpdateOne, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	updater := curd.Db.Service.UpdateOneID(id.ID)
	ServiceUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *ServiceCURD) UpdateOneRoutePath() string {
	return "/service/:id"
}

func (curd *ServiceCURD) UpdateOne(c *gin.Context) (*ent.Service, error) {
	updater, err := curd.BaseUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *ServiceCURD) BaseUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
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
		updater := tx.Service.UpdateOneID(v.ID)
		ServiceUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *ServiceCURD) UpdateListRoutePath() string {
	return "/services"
}

func (curd *ServiceCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.BaseUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *ServiceCURD) BaseDeleteOneDeleter(c *gin.Context) (*ent.ServiceDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Service.Delete().Where(service.IDEQ(id.ID)), nil
}

func (curd *ServiceCURD) DeleteOneRoutePath() string {
	return "/service/:id"
}

func (curd *ServiceCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *ServiceCURD) BaseDeleteListDeleter(c *gin.Context) (*ent.ServiceDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.Db.Service.Delete().Where(service.IDIn(ids.Ids...)), nil
}

func (curd *ServiceCURD) DeleteListRoutePath() string {
	return "/services"
}

func (curd *ServiceCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *ServiceCURD) GetListRoleBindingsByServiceIdRoutePath() string {
	return "/service/:id/role_bindings"
}

func (curd *ServiceCURD) GetListRoleBindingsByServiceId(c *gin.Context) ([]*ent.RoleBinding, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	tmpQueryer := queryer.QueryRoleBindings()

	query, err := curd.RoleBindingObj.BindDefaultQuery(c)
	if err != nil {
		return nil, err
	}
	err = query.Exec(tmpQueryer)
	if err != nil {
		return nil, err
	}
	RoleBindingSelete(tmpQueryer)
	curd.RoleBindingObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *ServiceCURD) CreateListRoleBindingsByServiceIdRoutePath() string {
	return "/service/:id/role_bindings"
}

func (curd *ServiceCURD) CreateListRoleBindingsByServiceId(c *gin.Context) ([]*ent.RoleBinding, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.RoleBindingObj.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}

	role_bindings, err := func() ([]*ent.RoleBinding, error) {
		if err != nil {
			return nil, err
		}
		role_bindings, err := tx.RoleBinding.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}
		_, err = tx.Service.UpdateOneID(id.ID).AddRoleBindings(role_bindings...).Save(bg)
		if err != nil {
			return nil, err
		}

		return role_bindings, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return role_bindings, tx.Commit()
}

func (curd *ServiceCURD) DeleteListRoleBindingsByServiceIdRoutePath() string {
	return "/service/:id/role_bindings"
}

func (curd *ServiceCURD) DeleteListRoleBindingsByServiceId(c *gin.Context) (int, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	query, err := curd.RoleBindingObj.BindDefaultQuery(c)

	if err != nil {
		return 0, err
	}

	ps, err := query.PredicatesExec()
	if err != nil {
		return 0, err
	}
	ids, err := queryer.QueryRoleBindings().Where(ps...).IDs(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.RoleBinding.Delete().Where(rolebinding.IDIn(ids...)).Exec(context.Background())
}

func (curd *ServiceCURD) GetListServersByServiceIdRoutePath() string {
	return "/service/:id/servers"
}

func (curd *ServiceCURD) GetListServersByServiceId(c *gin.Context) ([]*ent.Server, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	tmpQueryer := queryer.QueryServers()

	query, err := curd.ServerObj.BindDefaultQuery(c)
	if err != nil {
		return nil, err
	}
	err = query.Exec(tmpQueryer)
	if err != nil {
		return nil, err
	}
	ServerSelete(tmpQueryer)
	curd.ServerObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// M2M
func (curd *ServiceCURD) CreateListServersByServiceIdRoutePath() string {
	return "/service/:id/servers"
}

func (curd *ServiceCURD) CreateListServersByServiceId(c *gin.Context) ([]*ent.Server, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.ServerObj.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}

	servers, err := func() ([]*ent.Server, error) {
		if err != nil {
			return nil, err
		}
		servers, err := tx.Server.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}
		_, err = tx.Service.UpdateOneID(id.ID).AddServers(servers...).Save(bg)
		if err != nil {
			return nil, err
		}

		return servers, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return servers, tx.Commit()
}

func (curd *ServiceCURD) DeleteListServersByServiceIdRoutePath() string {
	return "/service/:id/servers"
}

func (curd *ServiceCURD) DeleteListServersByServiceId(c *gin.Context) (int, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	query, err := curd.ServerObj.BindDefaultQuery(c)

	if err != nil {
		return 0, err
	}

	ps, err := query.PredicatesExec()
	if err != nil {
		return 0, err
	}
	ids, err := queryer.QueryServers().Where(ps...).IDs(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Server.Delete().Where(server.IDIn(ids...)).Exec(context.Background())
}

func (curd *ServiceCURD) GetOneProjectByServiceIdRoutePath() string {
	return "/service/:id/project"
}

func (curd *ServiceCURD) GetOneProjectByServiceId(c *gin.Context) (*ent.Project, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	return queryer.QueryProject().First(context.Background())
}

// M2O
func (curd *ServiceCURD) CreateOneProjectByServiceIdRoutePath() string {
	return "/service/:id/project"
}

func (curd *ServiceCURD) CreateOneProjectByServiceId(c *gin.Context) (*ent.Project, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}

	projectCreater, err := curd.ProjectObj.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return projectCreater.AddServiceIDs(id.ID).Save(context.Background())
}

func (curd *ServiceCURD) DeleteOneProjectByServiceIdRoutePath() string {
	return "/service/:id/project"
}

func (curd *ServiceCURD) DeleteOneProjectByServiceId(c *gin.Context) (int, error) {
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
