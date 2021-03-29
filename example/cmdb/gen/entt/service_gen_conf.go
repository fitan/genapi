package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/service"
)

type ServiceIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=role_binding.user.alert project.role_binding.user project.role_binding.user.alert project project.role_binding role_binding role_binding.user role_binding.project server"`
}

type GetServiceListData struct {
	Count  int
	Result []*ent.Service
}

func ServiceSelete(queryer *ent.ServiceQuery) {
	queryer.Select(

		service.FieldCreateTime,

		service.FieldUpdateTime,

		service.FieldName,
	)
}

func ServiceCreateMutation(m *ent.ServiceMutation, v *ent.Service) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetProjectID(v.Edges.Project.ID)

}

func ServiceUpdateMutation(m *ent.ServiceMutation, v *ent.Service) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

	m.SetProjectID(v.Edges.Project.ID)

}

func ServiceGetIDs(services ent.Services) []int {
	IDs := make([]int, 0, len(services))
	for i, _ := range services {
		IDs[i] = services[i].ID
	}
	return IDs
}

type ServiceDefaultQuery struct {
	ServiceIncludes
}

func (s *ServiceDefaultQuery) PredicatesExec() ([]predicate.Service, error) {
	return ServicePredicatesExec()
}

func (s *ServiceDefaultQuery) Exec(queryer *ent.ServiceQuery) error {
	ps, err := s.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, s.Includes)

	queryer.Where(service.And(ps...))

	return nil
}
