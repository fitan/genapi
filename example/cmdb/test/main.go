package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	l := []string{
		"a,b,c",
		"a,b,d",
		"d", "e", "f",
	}
	tree := Includes2Tree(l)
	marshal, err := json.Marshal(tree)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(marshal))
}
