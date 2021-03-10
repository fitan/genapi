package genrest

import (
	"context"
	"ent_samp/ent"
	"ent_samp/ent/car"
	"fmt"

	"github.com/gin-gonic/gin"
)

type CarCURD struct {
	Db *ent.Client
}

func (curd *CarCURD) RegisterRouter(router interface{}) {
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

func (curd *CarCURD) BindObj(c *gin.Context) (*ent.Car, error) {
	body := new(ent.Car)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *CarCURD) BindObjs(c *gin.Context) (ent.Cars, error) {
	body := make(ent.Cars, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *CarCURD) BindDefaultQuery(c *gin.Context) (*CarDefaultQuery, error) {
	body := new(CarDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *CarCURD) BaseGetOneQueryer(id int) (*ent.CarQuery, error) {
	return curd.Db.Car.Query().Where(car.IDEQ(id)), nil
}

func (curd *CarCURD) defaultGetOneQueryer(c *gin.Context) (*ent.CarQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *CarCURD) GetOneRoutePath() string {
	return "/car/:id"
}

func (curd *CarCURD) GetOne(c *gin.Context) (*ent.Car, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	curd.selete(queryer)
	return queryer.Only(context.Background())
}

func (curd *CarCURD) BaseGetListCount(queryer *ent.CarQuery, query *CarDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(car.And(ps...))
	return nil
}

func (curd *CarCURD) BaseGetListQueryer(queryer *ent.CarQuery, query *CarDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	curd.selete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *CarCURD) defaultGetListQueryer(c *gin.Context) (*ent.CarQuery, *ent.CarQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.Car.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.Car.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *CarCURD) GetListRoutePath() string {
	return "/cars"
}

type GetCarListData struct {
	Count  int
	Result []*ent.Car
}

func (curd *CarCURD) GetList(c *gin.Context) (*GetCarListData, error) {
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

	return &GetCarListData{count, res}, nil
}

func (curd *CarCURD) createMutation(m *ent.CarMutation, v *ent.Car) {

	m.SetModel(v.Model)

	m.SetRegisteredAt(v.RegisteredAt)

}

func (curd *CarCURD) updateMutation(m *ent.CarMutation, v *ent.Car) {

	m.SetModel(v.Model)

	m.SetRegisteredAt(v.RegisteredAt)

}

func (curd *CarCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *CarCURD) defaultOrder(queryer *ent.CarQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *CarCURD) selete(queryer *ent.CarQuery) {
	queryer.Select(

		car.FieldModel,

		car.FieldRegisteredAt,
	)
}

func (curd *CarCURD) BaseCreateOneCreater(body *ent.Car) *ent.CarCreate {
	creater := curd.Db.Car.Create()
	curd.createMutation(creater.Mutation(), body)
	return creater
}

func (curd *CarCURD) defaultCreateOneCreater(c *gin.Context) (*ent.CarCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *CarCURD) CreateOneRoutePath() string {
	return "/car"
}

func (curd *CarCURD) CreateOne(c *gin.Context) (*ent.Car, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *CarCURD) BaseCreateListBulk(body ent.Cars) []*ent.CarCreate {
	bulk := make([]*ent.CarCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.Car.Create()
		curd.createMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *CarCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.CarCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *CarCURD) CreateListRoutePath() string {
	return "/cars"
}

func (curd *CarCURD) CreateList(c *gin.Context) ([]*ent.Car, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Car.CreateBulk(bulk...).Save(context.Background())
}

func (curd *CarCURD) BaseUpdateOneUpdater(id int, body *ent.Car) (*ent.CarUpdateOne, error) {
	updater := curd.Db.Car.UpdateOneID(id)
	curd.updateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *CarCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.CarUpdateOne, error) {
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

func (curd *CarCURD) UpdateOneRoutePath() string {
	return "/car/:id"
}

func (curd *CarCURD) UpdateOne(c *gin.Context) (*ent.Car, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *CarCURD) BaseUpdateListUpdater(body ent.Cars) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.Car.UpdateOneID(v.ID)
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

func (curd *CarCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *CarCURD) UpdateListRoutePath() string {
	return "/cars"
}

func (curd *CarCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *CarCURD) BaseDeleteOneDeleter(id int) *ent.CarDelete {
	return curd.Db.Car.Delete().Where(car.IDEQ(id))
}

func (curd *CarCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.CarDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *CarCURD) DeleteOneRoutePath() string {
	return "/car/:id"
}

func (curd *CarCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *CarCURD) BaseDeleteListDeleter(ids []int) *ent.CarDelete {
	return curd.Db.Car.Delete().Where(car.IDIn(ids...))
}

func (curd *CarCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.CarDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *CarCURD) DeleteListRoutePath() string {
	return "/cars"
}

func (curd *CarCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}
