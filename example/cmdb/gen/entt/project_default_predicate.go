package entt

import (
	"cmdb/ent"
	"cmdb/ent/predicate"
	"cmdb/ent/project"
)

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
