package casbin

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/rolebinding"
	"cmdb/ent/user"
	"cmdb/gen/entt"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCURDV2 struct {
	Db *ent.Client

	RoleBindingObj *entt.RoleBindingCURD

	AlertObj *entt.AlertCURD
}

func (curd *UserCURDV2) GetQueryer() *ent.UserQuery {
	return curd.Db.User.Query()
}

func (curd *UserCURDV2) GetCreater() *ent.UserCreate {
	return curd.Db.User.Create()
}

func (curd *UserCURDV2) GetDeleter() *ent.UserDelete {
	return curd.Db.User.Delete()
}

func (curd *UserCURDV2) GetUpdater(id int) *ent.UserUpdateOne {
	return curd.Db.User.UpdateOneID(id)
}

func (curd *UserCURDV2) GetOneQueryer(id int) *ent.UserQuery {
	return curd.GetQueryer().Where(user.IDEQ(id))
}

func (curd *UserCURDV2) GetOne(id int) (*ent.User, error) {
	return curd.GetOneQueryer(id).Only(context.Background())
}

func (curd *UserCURDV2) GetListQuery(query *entt.UserDefaultQuery) *ent.UserQuery {
	queryer := curd.GetQueryer()

	query.Exec(queryer)
	entt.UserSelete(queryer)
	curd.DefaultOrder(queryer)
	return queryer
}

func (curd *UserCURDV2) GetListCountQuery(query *entt.UserDefaultQuery) (*ent.UserQuery, error) {
	queryer := curd.GetQueryer()
	ps, err := query.PredicatesExec()
	if err != nil {
		return nil, err
	}
	return queryer.Where(user.And(ps...)),err
}

func (curd *UserCURDV2) GetList(query *entt.UserDefaultQuery) (*entt.GetUserListData, error) {
	bg := context.Background()
	list,err := curd.GetListQuery(query).All(bg)
	if err != nil {
		return nil, err
	}
	countQuery, err := curd.GetListCountQuery(query)
	if err != nil {
		return nil, err
	}
	count,err := countQuery.Count(bg)
	if err != nil {
		return nil, err
	}
	return &entt.GetUserListData{count, list}, nil
}

func (curd *UserCURDV2) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := entt.BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *UserCURDV2) DefaultOrder(queryer *ent.UserQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}



func (curd *UserCURDV2) CreateOne(user *ent.User) (*ent.User, error) {
	creater := curd.GetCreater()
	entt.UserCreateMutation(creater.Mutation(), user)
	return creater.Save(context.Background())
}


func (curd *UserCURDV2) GetBulk(users []*ent.User) []*ent.UserCreate {
	bulk := make([]*ent.UserCreate, 0, len(users))
	for _, v := range users {
		creater := curd.GetCreater()
		entt.UserCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *UserCURDV2) CreateList(users []*ent.User) ([]*ent.User, error) {
	return curd.Db.User.CreateBulk(curd.GetBulk(users)...).Save(context.Background())
}

func (curd *UserCURDV2) UpdateOne(user *ent.User) (*ent.User, error) {
	updater := curd.GetUpdater(user.ID)
	entt.UserUpdateMutation(updater.Mutation(), user)
	return updater.Save(context.Background())
}

func (curd *UserCURDV2) BaseUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
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
		entt.UserUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *UserCURDV2) UpdateList(users []*ent.User) (string, error) {
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return "", err
	}
	for _, v := range users {
		updater := tx.User.UpdateOneID(v.ID)
		entt.UserUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(bg)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return "", fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return "ok", tx.Commit()
}


func (curd *UserCURDV2) DeleteOne(id int) (int, error) {
	return curd.GetDeleter().Where(user.IDEQ(id)).Exec(context.Background())
}

func (curd *UserCURDV2) BaseDeleteListDeleter(c *gin.Context) (*ent.UserDelete, error) {
	ids, err := entt.BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.Db.User.Delete().Where(user.IDIn(ids.Ids...)), nil
}

func (curd *UserCURDV2) DeleteListRoutePath() string {
	return "/users"
}

func (curd *UserCURDV2) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.BaseDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *UserCURDV2) GetListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURDV2) GetListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
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
	entt.RoleBindingSelete(tmpQueryer)
	curd.RoleBindingObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *UserCURDV2) CreateListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURDV2) CreateListRoleBindingsByUserId(c *gin.Context) ([]*ent.RoleBinding, error) {
	id, err := entt.BindId(c)
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

func (curd *UserCURDV2) DeleteListRoleBindingsByUserIdRoutePath() string {
	return "/user/:id/role_bindings"
}

func (curd *UserCURDV2) DeleteListRoleBindingsByUserId(c *gin.Context) (int, error) {
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

func (curd *UserCURDV2) GetListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURDV2) GetListAlertsByUserId(c *gin.Context) ([]*ent.Alert, error) {
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
	entt.AlertSelete(tmpQueryer)
	curd.AlertObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *UserCURDV2) CreateListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURDV2) CreateListAlertsByUserId(c *gin.Context) ([]*ent.Alert, error) {
	id, err := entt.BindId(c)
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

func (curd *UserCURDV2) DeleteListAlertsByUserIdRoutePath() string {
	return "/user/:id/alerts"
}

func (curd *UserCURDV2) DeleteListAlertsByUserId(c *gin.Context) (int, error) {
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
