package casbin

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/user"
	"cmdb/gen/entt"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCURDV2 struct {
	Db *ent.Client

	RoleBindingObj *RoleCURDV2
	AlertObj       *entt.AlertCURD
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
	return queryer.Where(user.And(ps...)), err
}

func (curd *UserCURDV2) GetList(query *entt.UserDefaultQuery) (*entt.GetUserListData, error) {
	bg := context.Background()
	list, err := curd.GetListQuery(query).All(bg)
	if err != nil {
		return nil, err
	}
	countQuery, err := curd.GetListCountQuery(query)
	if err != nil {
		return nil, err
	}
	count, err := countQuery.Count(bg)
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

func (curd *UserCURDV2) DeleteList(ids []int) (int, error) {
	return curd.Db.User.Delete().Where(user.IDIn(ids...)).Exec(context.Background())
}

func (curd *UserCURDV2) GetListRoleBindingsByUserId(id int, query *entt.RoleBindingDefaultQuery) (*entt.GetRoleBindingListData, error) {
	listQueryer := curd.GetOneQueryer(id).QueryRoleBindings()
	countQueryer := curd.GetOneQueryer(id).QueryRoleBindings()
	return curd.RoleBindingObj.GetListByQueryer(listQueryer, countQueryer, query)
}

func (curd *UserCURDV2) CreateListRoleBindingsByUserId(id int, roleBindings ent.RoleBindings) (ent.RoleBindings, error) {
	bulk := curd.RoleBindingObj.GetBulk(roleBindings)

	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}
	save, err := func() (ent.RoleBindings, error) {
		save, err := tx.RoleBinding.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}

		_, err = tx.User.UpdateOneID(id).AddRoleBindings(save...).Save(bg)
		if err != nil {
			return nil, err
		}
		return save, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return save, nil

}

func (curd *UserCURDV2) DeleteListRoleBindingsByUserId(id int, query *entt.RoleBindingDefaultQuery) (int, error) {
	queryer := curd.GetOneQueryer(id)

	curd.RoleBindingObj.SetListCountQuery(queryer.QueryRoleBindings(), query)
	bg := context.Background()
	ids, err := queryer.IDs(bg)
	if err != nil {
		return 0, err
	}
	return curd.RoleBindingObj.DeleteList(ids)
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
