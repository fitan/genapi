package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/project"
)

type ProjectIncludes struct {
	Includes []string `form:"includes" json:"includes" binding:"dive,oneof=role_binding.service.server service.role_binding.user service.role_binding.user.alert service.server role_binding role_binding.user role_binding.user.alert service.role_binding service role_binding.service"`
}

type GetProjectListData struct {
	Count  int
	Result []*ent.Project
}

func ProjectSelete(queryer *ent.ProjectQuery) {
	queryer.Select(

		project.FieldCreateTime,

		project.FieldUpdateTime,

		project.FieldName,
	)
}

func ProjectCreateMutation(m *ent.ProjectMutation, v *ent.Project) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

}

func ProjectUpdateMutation(m *ent.ProjectMutation, v *ent.Project) {

	m.SetCreateTime(v.CreateTime)

	m.SetUpdateTime(v.UpdateTime)

	m.SetName(v.Name)

}

func ProjectGetIDs(projects ent.Projects) []int {
	IDs := make([]int, 0, len(projects))
	for i, _ := range projects {
		IDs[i] = projects[i].ID
	}
	return IDs
}

type ProjectDefaultQuery struct {
	ProjectIncludes
}

func (p *ProjectDefaultQuery) PredicatesExec() ([]predicate.Project, error) {
	return ProjectPredicatesExec()
}

func (p *ProjectDefaultQuery) Exec(queryer *ent.ProjectQuery) error {
	ps, err := p.PredicatesExec()
	if err != nil {
		return err
	}
	QueryerIncludes(queryer, p.Includes)

	queryer.Where(project.And(ps...))

	return nil
}
