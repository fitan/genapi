package entt

import (
	"cmdb/ent"
	"cmdb/ent/rolebinding"
)

func RoleBindingSelete(queryer *ent.RoleBindingQuery) {
	queryer.Select(

		rolebinding.FieldCreateTime,

		rolebinding.FieldUpdateTime,

		rolebinding.FieldRole,
	)
}

func RoleBindingCreateMutation(m *ent.RoleBindingMutation, v *ent.RoleBinding) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetRole(v.Role)

	m.SetProjectID(v.Edges.Project.ID)

	m.SetServiceID(v.Edges.Service.ID)

	m.SetUserID(v.Edges.User.ID)

}

func RoleBindingUpdateMutation(m *ent.RoleBindingMutation, v *ent.RoleBinding) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetRole(v.Role)

	m.SetProjectID(v.Edges.Project.ID)

	m.SetServiceID(v.Edges.Service.ID)

	m.SetUserID(v.Edges.User.ID)

}
