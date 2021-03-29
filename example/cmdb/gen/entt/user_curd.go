package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/rolebinding"
	"cmdb/ent/user"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCURD struct {
	Db *ent.Client

	RoleBindingObj *RoleBindingCURD

	AlertObj *AlertCURD
}

func (curd *UserCURD) RegisterRouter(router interface{}) {
	switch router.(type) {
	case *gin.Engine:
		r := router.(*gin.Engine)

		r.GET(curd.GetOneRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOne(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoutePath(), func(c *gin.Context) {
			data, err := curd.GetList(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

	case *gin.RouterGroup:
		r := router.(*gin.RouterGroup)

		r.GET(curd.GetOneRoutePath(), func(c *gin.Context) {
			data, err := curd.GetOne(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoutePath(), func(c *gin.Context) {
			data, err := curd.GetList(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListAlertsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListAlertsByUserId(c)
			RestReturnFunc(c, data, err)
		})

	}
}

func (curd *UserCURD) BindObj(c *gin.Context) (*ent.User, error) {
	body := new(ent.User)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *UserCURD) BindObjs(c *gin.Context) (ent.Users, error) {
	body := make(ent.Users, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *UserCURD) BindDefaultQuery(c *gin.Context) (*UserDefaultQuery, error) {
	body := new(UserDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *UserCURD) BaseGetOneQueryer(c *gin.Context) (*ent.UserQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.User.Query().Where(user.IDEQ(id.ID)), nil
}

func (curd *UserCURD) GetOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) GetOne(c *gin.Context) (*ent.User, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	UserSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *UserCURD) defaultGetListCount(queryer *ent.UserQuery, query *UserDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(user.And(ps...))
	return nil
}

func (curd *UserCURD) defaultGetListQueryer(queryer *ent.UserQuery, query *UserDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	UserSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *UserCURD) BaseGetListQueryer(c *gin.Context) (*ent.UserQuery, *ent.UserQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.User.Query()

	err = curd.defaultGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.User.Query()
	err = curd.defaultGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *UserCURD) GetListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) GetList(c *gin.Context) (*GetUserListData, error) {
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

	return &GetUserListData{count, res}, nil
}

func (curd *UserCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *UserCURD) defaultOrder(queryer *ent.UserQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *UserCURD) BaseCreateOneCreater(c *gin.Context) (*ent.UserCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	creater := curd.Db.User.Create()
	UserCreateMutation(creater.Mutation(), body)
	return creater, nil
}

func (curd *UserCURD) CreateOneRoutePath() string {
	return "/user"
}

func (curd *UserCURD) CreateOne(c *gin.Context) (*ent.User, error) {
	creater, err := curd.BaseCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *UserCURD) BaseCreateListBulk(c *gin.Context) ([]*ent.UserCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	bulk := make([]*ent.UserCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.User.Create()
		UserCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk, nil
}

func (curd *UserCURD) BaseCreateList(c *gin.Context) (*ent.UserCreateBulk, error) {
	bulk, err := curd.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.User.CreateBulk(bulk...), nil
}

func (curd *UserCURD) CreateListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) CreateList(c *gin.Context) ([]*ent.User, error) {
	creater, err := curd.BaseCreateList(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *UserCURD) BaseUpdateOneUpdater(c *gin.Context) (*ent.UserUpdateOne, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	updater := curd.Db.User.UpdateOneID(id.ID)
	UserUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *UserCURD) UpdateOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) UpdateOne(c *gin.Context) (*ent.User, error) {
	updater, err := curd.BaseUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *UserCURD) BaseUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
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
		updater := tx.User.UpdateOneID(v.ID)
		UserUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *UserCURD) UpdateListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.BaseUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *UserCURD) BaseDeleteOneDeleter(c *gin.Context) (*ent.UserDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.User.Delete().Where(user.IDEQ(id.ID)), nil
}

func (curd *UserCURD) DeleteOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *UserCURD) BaseDeleteListDeleter(c *gin.Context) (*ent.UserDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.Db.User.Delete().Where(user.IDIn(ids.Ids...)), nil
}

func (curd *UserCURD) DeleteListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *UserCURD) GetListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURD) GetListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
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
func (curd *UserCURD) CreateListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURD) CreateListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
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
		_, err = tx.User.UpdateOneID(id.ID).AddRoleBindings(role_bindings...).Save(bg)
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

func (curd *UserCURD) DeleteListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURD) DeleteListRoleBindingsByUserId(c *gin.Context) (int, error) {
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

func (curd *UserCURD) GetListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURD) GetListAlertsByUserId(c *gin.Context) ([]*ent.Alert, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	tmpQueryer := queryer.QueryAlerts()

	query, err := curd.AlertObj.BindDefaultQuery(c)
	if err != nil {
		return nil, err
	}
	err = query.Exec(tmpQueryer)
	if err != nil {
		return nil, err
	}
	AlertSelete(tmpQueryer)
	curd.AlertObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *UserCURD) CreateListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURD) CreateListAlertsByUserId(c *gin.Context) ([]*ent.Alert, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.AlertObj.BaseCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}

	alerts, err := func() ([]*ent.Alert, error) {
		if err != nil {
			return nil, err
		}
		alerts, err := tx.Alert.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}
		_, err = tx.User.UpdateOneID(id.ID).AddAlerts(alerts...).Save(bg)
		if err != nil {
			return nil, err
		}

		return alerts, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return alerts, tx.Commit()
}

func (curd *UserCURD) DeleteListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURD) DeleteListAlertsByUserId(c *gin.Context) (int, error) {
	queryer, err := curd.BaseGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	query, err := curd.AlertObj.BindDefaultQuery(c)

	if err != nil {
		return 0, err
	}

	ps, err := query.PredicatesExec()
	if err != nil {
		return 0, err
	}
	ids, err := queryer.QueryAlerts().Where(ps...).IDs(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Alert.Delete().Where(alert.IDIn(ids...)).Exec(context.Background())
}
