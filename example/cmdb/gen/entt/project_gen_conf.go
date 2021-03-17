package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/project"
	"github.com/gin-gonic/gin"
)

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
	for _, project := range projects {
		IDs = append(IDs, project.ID)
	}
	return IDs
}

type ProjectDefaultQuery struct {
}

func (p *ProjectDefaultQuery) PredicatesExec() ([]predicate.Project, error) {
	return ProjectPredicatesExec()
}

func (p *ProjectDefaultQuery) Exec(queryer *ent.ProjectQuery) error {
	ps, err := p.PredicatesExec()
	if err != nil {
		return err
	}

	queryer.Where(project.And(ps...))

	return nil
}
