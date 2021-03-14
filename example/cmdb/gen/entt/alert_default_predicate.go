package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/predicate"
)

type AlertDefaultQuery struct {
}

func (a *AlertDefaultQuery) PredicatesExec() ([]predicate.Alert, error) {
	return AlertPredicatesExec()
}

func (a *AlertDefaultQuery) Exec(queryer *ent.AlertQuery) error {
	ps, err := a.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(alert.And(ps...))

	return nil
}
