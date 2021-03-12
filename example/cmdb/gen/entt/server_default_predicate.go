package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/server"
)

type ServerDefaultQuery struct {
}

func (s *ServerDefaultQuery) PredicatesExec() ([]predicate.Server, error) {
	return ServerPredicatesExec()
}

func (s *ServerDefaultQuery) Exec(queryer *ent.ServerQuery) error {
	ps, err := s.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(server.And(ps...))

	return nil
}
