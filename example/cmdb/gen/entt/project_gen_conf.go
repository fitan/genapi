package entt

import (
	"cmdb/ent"
	"cmdb/ent/project"
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
