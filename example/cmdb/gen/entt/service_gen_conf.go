package entt

import (
	"cmdb/ent"
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
