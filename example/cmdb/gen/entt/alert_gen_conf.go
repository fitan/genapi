package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
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
