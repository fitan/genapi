package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/rolebinding"
)

type RoleBindingIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=service.project user user.alert project project.service project.service.server service service.server"`
}

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

func RoleBindingGetIDs(role_bindings ent.RoleBindings) []int {
	IDs := make([]int, 0, len(role_bindings))
	for _, role_binding := range role_bindings {
		IDs = append(IDs, role_binding.ID)
	}
	return IDs
}

type RoleBindingDefaultQuery struct {
	RoleBindingIncludes
}

func (rb *RoleBindingDefaultQuery) PredicatesExec() ([]predicate.RoleBinding, error) {
	return RoleBindingPredicatesExec()
}

func (rb *RoleBindingDefaultQuery) Exec(queryer *ent.RoleBindingQuery) error {
	ps, err := rb.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, rb.Includes)

	queryer.Where(rolebinding.And(ps...))

	return nil
}
