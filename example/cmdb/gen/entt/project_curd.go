package entt

import (
	"cmdb/ent"
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

type ProjectCURD struct {
	Db *ent.Client

	RoleBindingObj *RoleBindingCURD

	ServiceObj *ServiceCURD
}

func (curd *ProjectCURD) RegisterRouter(router interface{}) {
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

		r.POST(curd.CreateListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServicesByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServicesByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServicesByProjectId(c)
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

		r.POST(curd.CreateListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListRoleBindingsByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListRoleBindingsByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.POST(curd.CreateListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.CreateListServicesByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.DELETE(curd.DeleteListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.DeleteListServicesByProjectId(c)
			RestReturnFunc(c, data, err)
		})

		r.GET(curd.GetListServicesByProjectIdRoutePath(), func(c *gin.Context) {
			data, err := curd.GetListServicesByProjectId(c)
			RestReturnFunc(c, data, err)
		})

	}
}

func (curd *ProjectCURD) BindObj(c *gin.Context) (*ent.Project, error) {
	body := new(ent.Project)
	err := c.ShouldBindJSON(body)
	return body, err
}

func (curd *ProjectCURD) BindObjs(c *gin.Context) (ent.Projects, error) {
	body := make(ent.Projects, 0, 0)
	err := c.ShouldBindJSON(&body)
	return body, err
}

func (curd *ProjectCURD) BindDefaultQuery(c *gin.Context) (*ProjectDefaultQuery, error) {
	body := new(ProjectDefaultQuery)
	err := c.ShouldBindQuery(body)
	return body, err
}

func (curd *ProjectCURD) GetIDs(projects ent.Projects) []int {
	IDs := make([]int, 0, len(projects))
	for _, project := range projects {
		IDs = append(IDs, project.ID)
	}
	return IDs
}

func (curd *ProjectCURD) BaseGetOneQueryer(id int) (*ent.ProjectQuery, error) {
	return curd.Db.Project.Query().Where(project.IDEQ(id)), nil
}

func (curd *ProjectCURD) defaultGetOneQueryer(c *gin.Context) (*ent.ProjectQuery, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseGetOneQueryer(id.ID)
}

func (curd *ProjectCURD) GetOneRoutePath() string {
	return "/project/:id"
}

func (curd *ProjectCURD) GetOne(c *gin.Context) (*ent.Project, error) {
	queryer, err := curd.defaultGetOneQueryer(c)
	if err != nil {
		return nil, err
	}
	ProjectSelete(queryer)
	return queryer.Only(context.Background())
}

func (curd *ProjectCURD) BaseGetListCount(queryer *ent.ProjectQuery, query *ProjectDefaultQuery) error {
	ps, err := query.PredicatesExec()
	if err != nil {
		return err
	}
	queryer.Where(project.And(ps...))
	return nil
}

func (curd *ProjectCURD) BaseGetListQueryer(queryer *ent.ProjectQuery, query *ProjectDefaultQuery) error {
	err := query.Exec(queryer)
	if err != nil {
		return err
	}

	ProjectSelete(queryer)
	curd.defaultOrder(queryer)

	return nil
}

func (curd *ProjectCURD) defaultGetListQueryer(c *gin.Context) (*ent.ProjectQuery, *ent.ProjectQuery, error) {
	query, err := curd.BindDefaultQuery(c)
	if err != nil {
		return nil, nil, err
	}
	countQueryer := curd.Db.Project.Query()

	err = curd.BaseGetListCount(countQueryer, query)
	if err != nil {
		return nil, nil, err
	}

	getListQueryer := curd.Db.Project.Query()
	err = curd.BaseGetListQueryer(getListQueryer, query)
	if err != nil {
		return nil, nil, err
	}
	return getListQueryer, countQueryer, nil
}

func (curd *ProjectCURD) GetListRoutePath() string {
	return "/projects"
}

type GetProjectListData struct {
	Count  int
	Result []*ent.Project
}

func (curd *ProjectCURD) GetList(c *gin.Context) (*GetProjectListData, error) {
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

	return &GetProjectListData{count, res}, nil
}

func (curd *ProjectCURD) optionalOrder(c *gin.Context, queryer *ent.UserQuery) error {
	var expect = map[string]int{}
	orderFunc, err := BindOrder(c, expect)
	if err != nil {
		return err
	}
	queryer.Order(orderFunc...)
	return nil
}

func (curd *ProjectCURD) defaultOrder(queryer *ent.ProjectQuery) {
	queryer.Order([]ent.OrderFunc{
		ent.Asc(),
		ent.Desc(),
	}...)
}

func (curd *ProjectCURD) BaseCreateOneCreater(body *ent.Project) *ent.ProjectCreate {
	creater := curd.Db.Project.Create()
	ProjectCreateMutation(creater.Mutation(), body)
	return creater
}

func (curd *ProjectCURD) defaultCreateOneCreater(c *gin.Context) (*ent.ProjectCreate, error) {
	body, err := curd.BindObj(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateOneCreater(body), nil
}

func (curd *ProjectCURD) CreateOneRoutePath() string {
	return "/project"
}

func (curd *ProjectCURD) CreateOne(c *gin.Context) (*ent.Project, error) {
	creater, err := curd.defaultCreateOneCreater(c)
	if err != nil {
		return nil, err
	}
	return creater.Save(context.Background())
}

func (curd *ProjectCURD) BaseCreateListBulk(body ent.Projects) []*ent.ProjectCreate {
	bulk := make([]*ent.ProjectCreate, 0, len(body))
	for _, v := range body {
		creater := curd.Db.Project.Create()
		ProjectCreateMutation(creater.Mutation(), v)
		bulk = append(bulk, creater)
	}
	return bulk
}

func (curd *ProjectCURD) defaultCreateListBulk(c *gin.Context) ([]*ent.ProjectCreate, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseCreateListBulk(body), nil
}

func (curd *ProjectCURD) CreateListRoutePath() string {
	return "/projects"
}

func (curd *ProjectCURD) CreateList(c *gin.Context) ([]*ent.Project, error) {
	bulk, err := curd.defaultCreateListBulk(c)
	if err != nil {
		return nil, err
	}
	return curd.Db.Project.CreateBulk(bulk...).Save(context.Background())
}

func (curd *ProjectCURD) BaseUpdateOneUpdater(id int, body *ent.Project) (*ent.ProjectUpdateOne, error) {
	updater := curd.Db.Project.UpdateOneID(id)
	ProjectUpdateMutation(updater.Mutation(), body)
	return updater, nil
}

func (curd *ProjectCURD) defaultUpdateOneUpdater(c *gin.Context) (*ent.ProjectUpdateOne, error) {
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

func (curd *ProjectCURD) UpdateOneRoutePath() string {
	return "/project/:id"
}

func (curd *ProjectCURD) UpdateOne(c *gin.Context) (*ent.Project, error) {
	updater, err := curd.defaultUpdateOneUpdater(c)
	if err != nil {
		return nil, err
	}
	return updater.Save(context.Background())
}

func (curd *ProjectCURD) BaseUpdateListUpdater(body ent.Projects) (*ent.Tx, error) {
	ctx := context.Background()
	tx, err := curd.Db.Tx(ctx)
	if err != nil {
		return nil, err
	}
	for _, v := range body {
		updater := tx.Project.UpdateOneID(v.ID)
		ProjectUpdateMutation(updater.Mutation(), v)
		_, err := updater.Save(ctx)
		if err != nil {
			if rerr := tx.Rollback(); rerr != nil {
				return tx, fmt.Errorf("$v: %v", err, rerr)
			}
		}
	}
	return tx, nil
}

func (curd *ProjectCURD) defaultUpdateListUpdater(c *gin.Context) (*ent.Tx, error) {
	body, err := curd.BindObjs(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseUpdateListUpdater(body)
}

func (curd *ProjectCURD) UpdateListRoutePath() string {
	return "/projects"
}

func (curd *ProjectCURD) UpdateList(c *gin.Context) error {
	tx, err := curd.defaultUpdateListUpdater(c)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (curd *ProjectCURD) BaseDeleteOneDeleter(id int) *ent.ProjectDelete {
	return curd.Db.Project.Delete().Where(project.IDEQ(id))
}

func (curd *ProjectCURD) defaultDeleteOneDeleter(c *gin.Context) (*ent.ProjectDelete, error) {
	id, err := BindId(c)
	if err != nil {
		return nil, err
	}
	return curd.BaseDeleteOneDeleter(id.ID), nil
}

func (curd *ProjectCURD) DeleteOneRoutePath() string {
	return "/project/:id"
}

func (curd *ProjectCURD) DeleteOne(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteOneDeleter(c)
	if err != nil {
		return 0, err
	}
	return deleter.Exec(context.Background())
}

func (curd *ProjectCURD) BaseDeleteListDeleter(ids []int) *ent.ProjectDelete {
	return curd.Db.Project.Delete().Where(project.IDIn(ids...))
}

func (curd *ProjectCURD) defaultDeleteListDeleter(c *gin.Context) (*ent.ProjectDelete, error) {
	ids, err := BindIds(c)
	if err != nil {
		return nil, err
	}

	return curd.BaseDeleteListDeleter(ids.Ids), nil
}

func (curd *ProjectCURD) DeleteListRoutePath() string {
	return "/projects"
}

func (curd *ProjectCURD) DeleteList(c *gin.Context) (int, error) {
	deleter, err := curd.defaultDeleteListDeleter(c)
	if err != nil {
		return 0, nil
	}
	return deleter.Exec(context.Background())
}

func (curd *ProjectCURD) GetListRoleBindingsByProjectIdRoutePath() string {
	return "/project/:id/role_bindings"
}

func (curd *ProjectCURD) GetListRoleBindingsByProjectId(c *gin.Context) ([]*ent.RoleBinding, error) {
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
	RoleBindingSelete(tmpQueryer)
	curd.RoleBindingObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *ProjectCURD) CreateListRoleBindingsByProjectIdRoutePath() string {
	return "/project/:id/role_bindings"
}

func (curd *ProjectCURD) CreateListRoleBindingsByProjectId(c *gin.Context) ([]*ent.RoleBinding, error) {
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
		_, err = tx.Project.UpdateOneID(id.ID).AddRoleBindings(role_bindings...).Save(bg)
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

func (curd *ProjectCURD) DeleteListRoleBindingsByProjectIdRoutePath() string {
	return "/project/:id/role_bindings"
}

func (curd *ProjectCURD) DeleteListRoleBindingsByProjectId(c *gin.Context) (int, error) {
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

func (curd *ProjectCURD) GetListServicesByProjectIdRoutePath() string {
	return "/project/:id/services"
}

func (curd *ProjectCURD) GetListServicesByProjectId(c *gin.Context) ([]*ent.Service, error) {
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
	ServiceSelete(tmpQueryer)
	curd.ServiceObj.defaultOrder(tmpQueryer)

	return tmpQueryer.All(context.Background())

}

// O2M
func (curd *ProjectCURD) CreateListServicesByProjectIdRoutePath() string {
	return "/project/:id/services"
}

func (curd *ProjectCURD) CreateListServicesByProjectId(c *gin.Context) ([]*ent.Service, error) {
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
		_, err = tx.Project.UpdateOneID(id.ID).AddServices(services...).Save(bg)
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

func (curd *ProjectCURD) DeleteListServicesByProjectIdRoutePath() string {
	return "/project/:id/services"
}

func (curd *ProjectCURD) DeleteListServicesByProjectId(c *gin.Context) (int, error) {
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
