package entt

import (
	"cmdb/ent"
	"cmdb/ent/alert"
	"cmdb/ent/project"
	"cmdb/ent/rolebinding"
	"cmdb/ent/server"
	"cmdb/ent/service"
	"cmdb/ent/user"
	"strings"
)

type Tree struct {
	Names map[string]Tree `json:"names"`
}

func Includes2Tree(includes []string) Tree {
	tree := NewTree()
	for _, include := range includes {
		includeSplit := strings.Split(include, ".")
		t := Include2Tree(includeSplit)
		MergeTree(tree, t)
	}
	return tree
}

func MergeTree(tree Tree, t Tree) {
	if len(t.Names) == 0 {
		return
	}
	for name, tmpT := range t.Names {
		if _, ok := tree.Names[name]; !ok {
			tree.Names[name] = tmpT
		} else {
			MergeTree(tree.Names[name], t.Names[name])
		}
	}
}

func NewTree() Tree {
	return Tree{Names: make(map[string]Tree)}
}

func Include2Tree(ks []string) Tree {
	if len(ks) == 1 {
		tree := NewTree()
		tree.Names[ks[0]] = NewTree()
		return tree
	}
	t := Include2Tree(ks[1:])
	return Tree{
		Names: map[string]Tree{
			ks[0]: t,
		},
	}
}

func Depth(tree Tree) []interface{} {
	if len(tree.Names) == 0 {
		return []interface{}{nil}
	}

	resDepthTs := make([]interface{}, 0, 0)
	for name, t := range tree.Names {

		switch name {

		case alert.Label:
			depthTs := Depth(t)
			f := func(query *ent.AlertQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					default:
						AlertSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		case project.Label:
			depthTs := Depth(t)
			f := func(query *ent.ProjectQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					case func(query *ent.RoleBindingQuery):
						ProjectSelete(query.WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery))))

					case func(query *ent.ServiceQuery):
						ProjectSelete(query.WithServices(depthT.(func(query *ent.ServiceQuery))))

					default:
						ProjectSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		case rolebinding.Label:
			depthTs := Depth(t)
			f := func(query *ent.RoleBindingQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					case func(query *ent.ProjectQuery):
						RoleBindingSelete(query.WithProject(depthT.(func(query *ent.ProjectQuery))))

					case func(query *ent.ServiceQuery):
						RoleBindingSelete(query.WithService(depthT.(func(query *ent.ServiceQuery))))

					case func(query *ent.UserQuery):
						RoleBindingSelete(query.WithUser(depthT.(func(query *ent.UserQuery))))

					default:
						RoleBindingSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		case server.Label:
			depthTs := Depth(t)
			f := func(query *ent.ServerQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					case func(query *ent.ServiceQuery):
						ServerSelete(query.WithServices(depthT.(func(query *ent.ServiceQuery))))

					default:
						ServerSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		case service.Label:
			depthTs := Depth(t)
			f := func(query *ent.ServiceQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					case func(query *ent.RoleBindingQuery):
						ServiceSelete(query.WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery))))

					case func(query *ent.ServerQuery):
						ServiceSelete(query.WithServers(depthT.(func(query *ent.ServerQuery))))

					case func(query *ent.ProjectQuery):
						ServiceSelete(query.WithProject(depthT.(func(query *ent.ProjectQuery))))

					default:
						ServiceSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		case user.Label:
			depthTs := Depth(t)
			f := func(query *ent.UserQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {

					case func(query *ent.RoleBindingQuery):
						UserSelete(query.WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery))))

					case func(query *ent.AlertQuery):
						UserSelete(query.WithAlerts(depthT.(func(query *ent.AlertQuery))))

					default:
						UserSelete(query)
					}
				}
			}
			resDepthTs = append(resDepthTs, f)

		}
	}
	return resDepthTs
}

type Includes struct {
	Includes []string `form:"includes"`
}

func QueryerIncludes(queryer interface{}, includes []string) {
	tree := Includes2Tree(includes)
	depthTs := Depth(tree)
	switch queryer.(type) {

	case *ent.AlertQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			}
		}

	case *ent.ProjectQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			case func(query *ent.RoleBindingQuery):
				queryer.(*ent.ProjectQuery).WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery)))

			case func(query *ent.ServiceQuery):
				queryer.(*ent.ProjectQuery).WithServices(depthT.(func(query *ent.ServiceQuery)))

			}
		}

	case *ent.RoleBindingQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			case func(query *ent.ProjectQuery):
				queryer.(*ent.RoleBindingQuery).WithProject(depthT.(func(query *ent.ProjectQuery)))

			case func(query *ent.ServiceQuery):
				queryer.(*ent.RoleBindingQuery).WithService(depthT.(func(query *ent.ServiceQuery)))

			case func(query *ent.UserQuery):
				queryer.(*ent.RoleBindingQuery).WithUser(depthT.(func(query *ent.UserQuery)))

			}
		}

	case *ent.ServerQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			case func(query *ent.ServiceQuery):
				queryer.(*ent.ServerQuery).WithServices(depthT.(func(query *ent.ServiceQuery)))

			}
		}

	case *ent.ServiceQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			case func(query *ent.RoleBindingQuery):
				queryer.(*ent.ServiceQuery).WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery)))

			case func(query *ent.ServerQuery):
				queryer.(*ent.ServiceQuery).WithServers(depthT.(func(query *ent.ServerQuery)))

			case func(query *ent.ProjectQuery):
				queryer.(*ent.ServiceQuery).WithProject(depthT.(func(query *ent.ProjectQuery)))

			}
		}

	case *ent.UserQuery:
		for _, depthT := range depthTs {
			switch depthT.(type) {

			case func(query *ent.RoleBindingQuery):
				queryer.(*ent.UserQuery).WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery)))

			case func(query *ent.AlertQuery):
				queryer.(*ent.UserQuery).WithAlerts(depthT.(func(query *ent.AlertQuery)))

			}
		}

	}
}
