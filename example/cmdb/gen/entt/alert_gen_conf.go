package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/predicate"
)

func AlertSelete(queryer *ent.AlertQuery) {
	queryer.Select(

		alert.FieldName,
	)
}

func AlertCreateMutation(m *ent.AlertMutation, v *ent.Alert) {

	m.SetName(v.Name)

}

func AlertUpdateMutation(m *ent.AlertMutation, v *ent.Alert) {

	m.SetName(v.Name)

}

func AlertGetIDs(alerts ent.Alerts) []int {
	IDs := make([]int, 0, len(alerts))
	for _, alert := range alerts {
		IDs = append(IDs, alert.ID)
	}
	return IDs
}

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
