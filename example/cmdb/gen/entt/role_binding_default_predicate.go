package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
)

type RoleBindingDefaultQuery struct {
}

func (rb *RoleBindingDefaultQuery) PredicatesExec() ([]predicate.RoleBinding, error) {
	return RoleBindingPredicatesExec()
}

func (rb *RoleBindingDefaultQuery) Exec(queryer *ent.RoleBindingQuery) error {
	ps, err := rb.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(role_binding.And(ps...))

	return nil
}
