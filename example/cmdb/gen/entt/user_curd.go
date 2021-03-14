package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/service"
	"cmdb/ent/user"
	"context"
	"fmt"
	"strings"

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

func (curd *UserCURD) GetIDs(users ent.Users) []int {
	IDs := make([]int, 0, len(users))
	for _, user := range users {
		IDs = append(IDs, user.ID)
	}
	return IDs
}

func (curd *UserCURD) BaseGetOneQueryer(id int) (*ent.UserQuery, error) {
	return curd.Db.User.Query().Where(user.IDEQ(id)), nil
}

func (curd *UserCURD) defaultGetOneQueryer(c *gin.Context) (*ent.UserQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *UserCURD) GetOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) GetOne(c *gin.Context) (*ent.User, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	curd.selete(queryer)
	return queryer.Only(context.Background())
}

func (curd *UserCURD) BaseGetListCount(queryer *ent.UserQuery, query *UserDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(user.And(ps...))
	return nil
}

func (curd *UserCURD) BaseGetListQueryer(queryer *ent.UserQuery, query *UserDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	curd.selete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *UserCURD) defaultGetListQueryer(c *gin.Context) (*ent.UserQuery, *ent.UserQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.User.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.User.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *UserCURD) GetListRoutePath() string {
	return "/users"
}

type GetUserListData struct {
	Count  int
	Result []*ent.User
}

func (curd *UserCURD) GetList(c *gin.Context) (*GetUserListData, error) {
	getListQueryer, countQueryer, err := curd.defaultGetListQueryer(c)
	if err != nil {
		return nil, err
	}

	count, err := countQueryer.Count(context.Background())
	if err != nil {
		return nil, err
	}

	res, err := getListQueryer.WithRoleBindings(func(query *ent.RoleBindingQuery) {
		query.WithProject().WithService()
	}).All(context.Background())
	if err != nil {
		return nil, err
	}

	return &GetUserListData{count, res}, nil
}

func (curd *UserCURD) createMutation(m *ent.UserMutation, v *ent.User) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetPassword(v.Password)

	m.SetEmail(v.Email)

	m.SetPhone(v.Phone)

	m.SetRole(v.Role)

}

func (curd *UserCURD) updateMutation(m *ent.UserMutation, v *ent.User) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetPassword(v.Password)

	m.SetEmail(v.Email)

	m.SetPhone(v.Phone)

	m.SetRole(v.Role)

}

type IncludeTree struct {
	Names map[string]IncludeTree
}

func UserWith(son string) func(query *ent.UserQuery) {
	switch son {
	case rolebinding.Label:
		return func(query *ent.UserQuery) {
			query.WithRoleBindings()
		}
	case alert.Label:
		return func(query *ent.UserQuery) {
			query.WithAlerts()
		}
	}
	return nil
}

func RoleBindingWith(son string) func(query *ent.RoleBindingQuery) {
	switch son {
	case project.Label:
		return func(query *ent.RoleBindingQuery) {
			query.WithProject()
		}
	case service.Label:
		return func(query *ent.RoleBindingQuery) {
			query.WithService()
		}
	}
	return nil
}

func Depth(includes []string) interface{} {
	if len(includes) == 2 {
		switch includes[0] {
		case rolebinding.Label:
			return RoleBindingWith(includes[1])
		case user.Label:
			return UserWith(includes[1])
		}
	}

	tmp := includes[1:]
	depth := Depth(tmp)
	switch tmp[0] {
	case rolebinding.Label:
		fc := depth.(func(query *ent.RoleBindingQuery))
		switch includes[0] {
		case user.Label:
			return func(query *ent.UserQuery) {
				query.WithRoleBindings(fc)
			}
		case project.Label:
			return func(query *ent.ProjectQuery) {
				query.WithRoleBindings(fc)
			}
		case service.Label:
			return func(query *ent.ServiceQuery) {
				query.WithRoleBindings(fc)
			}
		}
	case user.Label:
		fc := depth.(func(query *ent.UserQuery))
	}
}

func (curd *UserCURD) includes(m *ent.UserQuery, includes []string) {
	m.WithRoleBindings(func(query *ent.RoleBindingQuery) {
		query.WithService(func(query *ent.ServiceQuery) {
			query.WithProject()
		})
	})
	m.WithRoleBindings(func(query *ent.RoleBindingQuery) {
	})
	tree := IncludeTree{}
	for _, include := range includes {
		includeSplit := strings.Split(include, ",")
		switch includeSplit[len(includeSplit)-1] {
		case rolebinding.Label:
			fc := curd.RoleBindingWith(includeSplit[len(includeSplit)])
			f := func(query *ent.UserQuery) {
				query.WithRoleBindings(fc)
			}
		case user.Label:
			fc := curd.UserWith(includeSplit[len(includeSplit)])

		}
	}
	m.wi
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

func (curd *UserCURD) selete(queryer *ent.UserQuery) {
	queryer.Select(

		user.FieldCreateTime,

		user.FieldUpdateTime,

		user.FieldName,

		user.FieldPassword,

		user.FieldEmail,

		user.FieldPhone,

		user.FieldRole,
	)
}

func (curd *UserCURD) BaseCreateOneCreater(body *ent.User) *ent.UserCreate {
	creater := curd.Db.User.Create()
	curd.createMutation(creater.Mutation(), body)
	return creater
}

func (curd *UserCURD) defaultCreateOneCreater(c *gin.Context) (*ent.UserCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *UserCURD) CreateOneRoutePath() string {
	return "/user"
}

func (curd *UserCURD) CreateOne(c *gin.Context) (*ent.User, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *UserCURD) BaseCreateListBulk(body ent.Users) []*ent.UserCreate {
	bulk := make([]*ent.UserCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.User.Create()
		curd.createMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *UserCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.UserCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *UserCURD) CreateListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) CreateList(c *gin.Context) ([]*ent.User, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.User.CreateBulk(bulk...).Save(context.Background())
}

func (curd *UserCURD) BaseUpdateOneUpdater(id int, body *ent.User) (*ent.UserUpdateOne, error) {
	updater := curd.Db.User.UpdateOneID(id)
	curd.updateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *UserCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.UserUpdateOne, error) {
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

func (curd *UserCURD) UpdateOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) UpdateOne(c *gin.Context) (*ent.User, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *UserCURD) BaseUpdateListUpdater(body ent.Users) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.User.UpdateOneID(v.ID)
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

func (curd *UserCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *UserCURD) UpdateListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *UserCURD) BaseDeleteOneDeleter(id int) *ent.UserDelete {
	return curd.Db.User.Delete().Where(user.IDEQ(id))
}

func (curd *UserCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.UserDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *UserCURD) DeleteOneRoutePath() string {
	return "/user/:id"
}

func (curd *UserCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *UserCURD) BaseDeleteListDeleter(ids []int) *ent.UserDelete {
	return curd.Db.User.Delete().Where(user.IDIn(ids...))
}

func (curd *UserCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.UserDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *UserCURD) DeleteListRoutePath() string {
	return "/users"
}

func (curd *UserCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *UserCURD) GetListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURD) GetListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
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
	curd.RoleBindingObj.selete(tmpQueryer)
	curd.RoleBindingObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

func (curd *UserCURD) CreateListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURD) CreateListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.RoleBindingObj.defaultCreateListBulk(c)
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
	queryer, err := curd.defaultGetOneQueryer(c)
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
	queryer, err := curd.defaultGetOneQueryer(c)
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
	curd.AlertObj.selete(tmpQueryer)
	curd.AlertObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

func (curd *UserCURD) CreateListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURD) CreateListAlertsByUserId(c *gin.Context) ([]*ent.Alert, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.AlertObj.defaultCreateListBulk(c)
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
	queryer, err := curd.defaultGetOneQueryer(c)
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
