package public

import (
	"cmdb/ent"
	"cmdb/ent/user"
	"strings"
)

type Tree struct {
	Names map[string]Tree `json:"names"`
}

func Includes2Tree(includes []string) Tree {
	tree := NewTree()
	for _, include := range includes {
		includeSplit := strings.Split(include, ",")
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
		return nil
	}

	resDepthTs := make([]interface{}, 0, 0)
	for name, t := range tree.Names {

		switch name {
		case user.Label:
			depthTs := Depth(t)
			f := func(query *ent.UserQuery) {
				for _, depthT := range depthTs {
					switch depthT.(type) {
					case func(query *ent.RoleBindingQuery):
						query.WithRoleBindings(depthT.(func(query *ent.RoleBindingQuery)))
					case func(query *ent.AlertQuery):
						query.WithAlerts(depthT.(func(query *ent.AlertQuery)))
					}
				}
			}
			resDepthTs = append(resDepthTs, f)
		}
	}
	return resDepthTs
}
