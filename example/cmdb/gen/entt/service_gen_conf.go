package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/service"
)

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
	for _, service := range services {
		IDs = append(IDs, service.ID)
	}
	return IDs
}

type ServiceDefaultQuery struct {
}

func (s *ServiceDefaultQuery) PredicatesExec() ([]predicate.Service, error) {
	return ServicePredicatesExec()
}

func (s *ServiceDefaultQuery) Exec(queryer *ent.ServiceQuery) error {
	ps, err := s.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(service.And(ps...))

	return nil
}
