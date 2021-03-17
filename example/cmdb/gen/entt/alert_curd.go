package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/predicate"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type AlertCURD struct {
	Db *ent.Client
}

func (curd *AlertCURD) RegisterRouter(router interface{}) {
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

func (curd *AlertCURD) BindObj(c *gin.Context) (*ent.Alert, error) {
	body := new(ent.Alert)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *AlertCURD) BindObjs(c *gin.Context) (ent.Alerts, error) {
	body := make(ent.Alerts, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *AlertCURD) BindDefaultQuery(c *gin.Context) (*AlertDefaultQuery, error) {
	body := new(AlertDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *AlertCURD) GetIDs(alerts ent.Alerts) []int {
	IDs := make([]int, 0, len(alerts))
	for _, alert := range alerts {
		IDs = append(IDs, alert.ID)
	}
	return IDs
}

func (curd *AlertCURD) BaseGetOneQueryer(id int) (*ent.AlertQuery, error) {
	return curd.Db.Alert.Query().Where(alert.IDEQ(id)), nil
}

func (curd *AlertCURD) defaultGetOneQueryer(c *gin.Context) (*ent.AlertQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *AlertCURD) GetOneRoutePath() string {
	return "/alert/:id"
}

func (curd *AlertCURD) GetOne(c *gin.Context) (*ent.Alert, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	AlertSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *AlertCURD) BaseGetListCount(queryer *ent.AlertQuery, query *AlertDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(alert.And(ps...))
	return nil
}

func (curd *AlertCURD) BaseGetListQueryer(queryer *ent.AlertQuery, query *AlertDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	AlertSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *AlertCURD) defaultGetListQueryer(c *gin.Context) (*ent.AlertQuery, *ent.AlertQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.Alert.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.Alert.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *AlertCURD) GetListRoutePath() string {
	return "/alerts"
}

type GetAlertListData struct {
	Count  int
	Result []*ent.Alert
}

func (curd *AlertCURD) GetList(c *gin.Context) (*GetAlertListData, error) {
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

	return &GetAlertListData{count, res}, nil
}

func (curd *AlertCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *AlertCURD) defaultOrder(queryer *ent.AlertQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *AlertCURD) BaseCreateOneCreater(body *ent.Alert) *ent.AlertCreate {
	creater := curd.Db.Alert.Create()
	AlertCreateMutation(creater.Mutation(), body)
	return creater
}

func (curd *AlertCURD) defaultCreateOneCreater(c *gin.Context) (*ent.AlertCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *AlertCURD) CreateOneRoutePath() string {
	return "/alert"
}

func (curd *AlertCURD) CreateOne(c *gin.Context) (*ent.Alert, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *AlertCURD) BaseCreateListBulk(body ent.Alerts) []*ent.AlertCreate {
	bulk := make([]*ent.AlertCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.Alert.Create()
		AlertCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *AlertCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.AlertCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *AlertCURD) CreateListRoutePath() string {
	return "/alerts"
}

func (curd *AlertCURD) CreateList(c *gin.Context) ([]*ent.Alert, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Alert.CreateBulk(bulk...).Save(context.Background())
}

func (curd *AlertCURD) BaseUpdateOneUpdater(id int, body *ent.Alert) (*ent.AlertUpdateOne, error) {
	updater := curd.Db.Alert.UpdateOneID(id)
	AlertUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *AlertCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.AlertUpdateOne, error) {
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

func (curd *AlertCURD) UpdateOneRoutePath() string {
	return "/alert/:id"
}

func (curd *AlertCURD) UpdateOne(c *gin.Context) (*ent.Alert, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *AlertCURD) BaseUpdateListUpdater(body ent.Alerts) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.Alert.UpdateOneID(v.ID)
		AlertUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *AlertCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *AlertCURD) UpdateListRoutePath() string {
	return "/alerts"
}

func (curd *AlertCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *AlertCURD) BaseDeleteOneDeleter(id int) *ent.AlertDelete {
	return curd.Db.Alert.Delete().Where(alert.IDEQ(id))
}

func (curd *AlertCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.AlertDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *AlertCURD) DeleteOneRoutePath() string {
	return "/alert/:id"
}

func (curd *AlertCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *AlertCURD) BaseDeleteListDeleter(ids []int) *ent.AlertDelete {
	return curd.Db.Alert.Delete().Where(alert.IDIn(ids...))
}

func (curd *AlertCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.AlertDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *AlertCURD) DeleteListRoutePath() string {
	return "/alerts"
}

func (curd *AlertCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}
