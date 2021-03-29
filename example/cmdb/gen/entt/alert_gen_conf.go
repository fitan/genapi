package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/predicate"
)

type AlertIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof="`
}

type GetAlertListData struct {
	Count  int
	Result []*ent.Alert
}

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
	for i, _ := range alerts {
		IDs[i] = alerts[i].ID
	}
	return IDs
}

type AlertDefaultQuery struct {
	AlertIncludes
}

func (a *AlertDefaultQuery) PredicatesExec() ([]predicate.Alert, error) {
	return AlertPredicatesExec()
}

func (a *AlertDefaultQuery) Exec(queryer *ent.AlertQuery) error {
	ps, err := a.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, a.Includes)

	queryer.Where(alert.And(ps...))

	return nil
}
