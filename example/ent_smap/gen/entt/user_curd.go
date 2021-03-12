package entt

import (
	"context"
	"ent_samp/ent"
	"ent_samp/ent/car"
	"ent_samp/ent/user"
	"fmt"

	"github.com/gin-gonic/gin"
)

type UserCURD struct {
	Db *ent.Client

	CarObj *CarCURD
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

		r.POST(curd.CreateListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListCarsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListCarsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListCarsByUserId(c)
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

		r.POST(curd.CreateListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListCarsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListCarsByUserId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListCarsByUserIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListCarsByUserId(c)
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

	res, err := getListQueryer.All(context.Background())
	if err != nil {
		return nil, err
	}

	return &GetUserListData{count, res}, nil
}

func (curd *UserCURD) createMutation(m *ent.UserMutation, v *ent.User) {

	m.SetName(v.Name)

	m.SetAge1(v.Age1)

	m.SetEn(v.En)

}

func (curd *UserCURD) updateMutation(m *ent.UserMutation, v *ent.User) {

	m.SetName(v.Name)

	m.SetAge1(v.Age1)

	m.SetEn(v.En)

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

		user.FieldName,

		user.FieldAge1,

		user.FieldEn,
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

func (curd *UserCURD) GetListCarsByUserIdRoutePath() string {
	return "/user/:id/cars"
}

func (curd *UserCURD) GetListCarsByUserId(c *gin.Context) ([]*ent.Car, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}

	tmpQueryer := queryer.QueryCars()

	query, err := curd.CarObj.BindDefaultQuery(c)
	if err != nil {
		return nil, err
	}
	err = query.Exec(tmpQueryer)
	if err != nil {
		return nil, err
	}
	curd.CarObj.selete(tmpQueryer)
	curd.CarObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

func (curd *UserCURD) CreateListCarsByUserIdRoutePath() string {
	return "/user/:id/cars"
}

func (curd *UserCURD) CreateListCarsByUserId(c *gin.Context) ([]*ent.Car, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	bulk, err := curd.CarObj.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	bg := context.Background()
	tx, err := curd.Db.Tx(bg)
	if err != nil {
		return nil, err
	}

	cars, err := func() ([]*ent.Car, error) {
		if err != nil {
			return nil, err
		}
		cars, err := tx.Car.CreateBulk(bulk...).Save(bg)
		if err != nil {
			return nil, err
		}
		_, err = tx.User.UpdateOneID(id.ID).AddCars(cars...).Save(bg)
		if err != nil {
			return nil, err
		}

		return cars, nil
	}()
	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = fmt.Errorf("%v: %v", err, rerr)
		}
		return nil, err
	}
	return cars, tx.Commit()
}

func (curd *UserCURD) DeleteListCarsByUserIdRoutePath() string {
	return "/user/:id/cars"
}

func (curd *UserCURD) DeleteListCarsByUserId(c *gin.Context) (int, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return 0, err
	}

	query, err := curd.CarObj.BindDefaultQuery(c)

	if err != nil {
		return 0, err
	}

	ps, err := query.PredicatesExec()
	if err != nil {
		return 0, err
	}
	ids, err := queryer.QueryCars().Where(ps...).IDs(context.Background())
	if err != nil {
		return 0, err
	}

	return curd.Db.Car.Delete().Where(car.IDIn(ids...)).Exec(context.Background())
}
