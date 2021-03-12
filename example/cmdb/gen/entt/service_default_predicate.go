package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/service"
)

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
