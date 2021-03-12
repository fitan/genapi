package entt

import (
	"cmdb/ent"
	"cmdb/ent/server"
	"cmdb/ent/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ServerCURD struct {
	Db *ent.Client

	ServiceObj *ServiceCURD
}

func (curd *ServerCURD) RegisterRouter(router interface{}) {
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

		r.POST(curd.CreateListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServicesByServerId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServicesByServerId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServicesByServerId(c)
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

		r.POST(curd.CreateListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServicesByServerId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServicesByServerId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServicesByServerIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServicesByServerId(c)
			RestReturnFunc(c, data, err)
		})

	}
}

func (curd *ServerCURD) BindObj(c *gin.Context) (*ent.Server, error) {
	body := new(ent.Server)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *ServerCURD) BindObjs(c *gin.Context) (ent.Servers, error) {
	body := make(ent.Servers, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *ServerCURD) BindDefaultQuery(c *gin.Context) (*ServerDefaultQuery, error) {
	body := new(ServerDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *ServerCURD) BaseGetOneQueryer(id int) (*ent.ServerQuery, error) {
	return curd.Db.Server.Query().Where(server.IDEQ(id)), nil
}

func (curd *ServerCURD) defaultGetOneQueryer(c *gin.Context) (*ent.ServerQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *ServerCURD) GetOneRoutePath() string {
	return "/server/:id"
}

func (curd *ServerCURD) GetOne(c *gin.Context) (*ent.Server, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	curd.selete(queryer)
	return queryer.Only(context.Background())
}

func (curd *ServerCURD) BaseGetListCount(queryer *ent.ServerQuery, query *ServerDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(server.And(ps...))
	return nil
}

func (curd *ServerCURD) BaseGetListQueryer(queryer *ent.ServerQuery, query *ServerDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	curd.selete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *ServerCURD) defaultGetListQueryer(c *gin.Context) (*ent.ServerQuery, *ent.ServerQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.Server.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.Server.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *ServerCURD) GetListRoutePath() string {
	return "/servers"
}

type GetServerListData struct {
	Count  int
	Result []*ent.Server
}

func (curd *ServerCURD) GetList(c *gin.Context) (*GetServerListData, error) {
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

	return &GetServerListData{count, res}, nil
}

func (curd *ServerCURD) createMutation(m *ent.ServerMutation, v *ent.Server) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetIP(v.IP)

	m.SetMachineType(v.MachineType)

	m.SetPlatformType(v.PlatformType)

	m.SetSystemType(v.SystemType)

}

func (curd *ServerCURD) updateMutation(m *ent.ServerMutation, v *ent.Server) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetIP(v.IP)

	m.SetMachineType(v.MachineType)

	m.SetPlatformType(v.PlatformType)

	m.SetSystemType(v.SystemType)

}

func (curd *ServerCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *ServerCURD) defaultOrder(queryer *ent.ServerQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *ServerCURD) selete(queryer *ent.ServerQuery) {
	queryer.Select(

		server.FieldCreateTime,

		server.FieldUpdateTime,

		server.FieldIP,

		server.FieldMachineType,

		server.FieldPlatformType,

		server.FieldSystemType,
	)
}

func (curd *ServerCURD) BaseCreateOneCreater(body *ent.Server) *ent.ServerCreate {
	creater := curd.Db.Server.Create()
	curd.createMutation(creater.Mutation(), body)
	return creater
}

func (curd *ServerCURD) defaultCreateOneCreater(c *gin.Context) (*ent.ServerCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *ServerCURD) CreateOneRoutePath() string {
	return "/server"
}

func (curd *ServerCURD) CreateOne(c *gin.Context) (*ent.Server, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *ServerCURD) BaseCreateListBulk(body ent.Servers) []*ent.ServerCreate {
	bulk := make([]*ent.ServerCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.Server.Create()
		curd.createMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *ServerCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.ServerCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *ServerCURD) CreateListRoutePath() string {
	return "/servers"
}

func (curd *ServerCURD) CreateList(c *gin.Context) ([]*ent.Server, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Server.CreateBulk(bulk...).Save(context.Background())
}

func (curd *ServerCURD) BaseUpdateOneUpdater(id int, body *ent.Server) (*ent.ServerUpdateOne, error) {
	updater := curd.Db.Server.UpdateOneID(id)
	curd.updateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *ServerCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.ServerUpdateOne, error) {
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

func (curd *ServerCURD) UpdateOneRoutePath() string {
	return "/server/:id"
}

func (curd *ServerCURD) UpdateOne(c *gin.Context) (*ent.Server, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *ServerCURD) BaseUpdateListUpdater(body ent.Servers) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.Server.UpdateOneID(v.ID)
		curd.updateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *ServerCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *ServerCURD) UpdateListRoutePath() string {
	return "/servers"
}

func (curd *ServerCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *ServerCURD) BaseDeleteOneDeleter(id int) *ent.ServerDelete {
	return curd.Db.Server.Delete().Where(server.IDEQ(id))
}

func (curd *ServerCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.ServerDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *ServerCURD) DeleteOneRoutePath() string {
	return "/server/:id"
}

func (curd *ServerCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *ServerCURD) BaseDeleteListDeleter(ids []int) *ent.ServerDelete {
	return curd.Db.Server.Delete().Where(server.IDIn(ids...))
}

func (curd *ServerCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.ServerDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *ServerCURD) DeleteListRoutePath() string {
	return "/servers"
}

func (curd *ServerCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *ServerCURD) GetListServicesByServerIdRoutePath() string {
	return "/server/:id/services"
}

func (curd *ServerCURD) GetListServicesByServerId(c *gin.Context) ([]*ent.Service, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	tmpQueryer := queryer.QueryServices()

	query, err := curd.ServiceObj.BindDefaultQuery(c)
	if err != nil {
		return nil, err
	}
	err = query.Exec(tmpQueryer)
	if err != nil {
		return nil, err
	}
	curd.ServiceObj.selete(tmpQueryer)
	curd.ServiceObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

func (curd *ServerCURD) CreateListServicesByServerIdRoutePath() string {
	return "/server/:id/services"
}

func (curd *ServerCURD) CreateListServicesByServerId(c *gin.Context) ([]*ent.Service, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.ServiceObj.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}

	services, err := func() ([]*ent.Service, error) {
		if err != nil {
			return nil, err
		}
		services, err := tx.Service.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}
		_, err = tx.Server.UpdateOneID(id.ID).AddServices(services...).Save(bg)
		if err != nil {
			return nil, err
		}

		return services, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return services, tx.Commit()
}

func (curd *ServerCURD) DeleteListServicesByServerIdRoutePath() string {
	return "/server/:id/services"
}

func (curd *ServerCURD) DeleteListServicesByServerId(c *gin.Context) (int, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	query, err := curd.ServiceObj.BindDefaultQuery(c)

	if err != nil {
		return 0, err
	}

	ps, err := query.PredicatesExec()
	if err != nil {
		return 0, err
	}
	ids, err := queryer.QueryServices().Where(ps...).IDs(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Service.Delete().Where(service.IDIn(ids...)).Exec(context.Background())
}
