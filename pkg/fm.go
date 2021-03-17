package pkg

import (
	"encoding/json"
	"entgo.io/ent/entc/gen"
	"fmt"
	"log"
	"strings"
	"text/template"
)

var FM = template.FuncMap{
	"opsString":                OpsString,
	"PaseRestNodeOrderOp":      PaseRestNodeOrderOp,
	"PaseRestFieldQueryOp":     PaseRestFieldQueryOp,
	"PaseRestFieldOperability": PaseRestFieldOperability,
	"PaseRestNodePaging":       PaseRestNodePaging,
	"PaseRestNodeMethod":       PaseRestNodeMethod,
	"PaseRestEdgeMethod":       PaseRestEdgeMethod,
	"PaseFieldIsEnum":          PaseFieldIsEnum,
	"PaseRestEdgeInclude":      PaseRestEdgeInclude,
	"PaseGraphInclude":         PaseGraphInclude,
}

func OpsString(ops []gen.Op) []string {
	opss := make([]string, 0, len(ops))
	for _, op := range ops {
		opss = append(opss, opText[op])
	}
	return opss
}

var opText = [...]string{
	gen.EQ:           "EQ",
	gen.NEQ:          "NEQ",
	gen.GT:           "GT",
	gen.GTE:          "GTE",
	gen.LT:           "LT",
	gen.LTE:          "LTE",
	gen.IsNil:        "IsNil",
	gen.NotNil:       "NotNil",
	gen.EqualFold:    "EqualFold",
	gen.Contains:     "Contains",
	gen.ContainsFold: "ContainsFold",
	gen.HasPrefix:    "HasPrefix",
	gen.HasSuffix:    "HasSuffix",
	gen.In:           "In",
	gen.NotIn:        "NotIn",
}

var RestFieldType = "RestFieldOp"
var RestNodeType = "RestNodeOp"
var RestEdgeType = "RestEdgeOp"

type RestEdgeOp struct {
	Method  EdgeMethod
	Include GenRestSwitch
}

type EdgeMethod struct {
	Get    GenRestSwitch `json:"get"`
	Create GenRestSwitch `json:"create"`
	Delete GenRestSwitch `json:"delete"`
}

type GenRestSwitch int

const GenRestDefault GenRestSwitch = 0
const GenRestTrue GenRestSwitch = 1
const GenRestFalse GenRestSwitch = 2

type RestNodeOp struct {
	Paging Paging     `json:"paging"`
	Order  Order      `json:"order"`
	Method NodeMethod `json:"method"`
}

type NodeMethod struct {
	Get    GenRestSwitch `json:"get"`
	Create GenRestSwitch `json:"create"`
	Update GenRestSwitch `json:"update"`
	Delete GenRestSwitch `json:"delete"`
}

type Order struct {
	DefaultAcsOrder   []string `json:"default_acs_order"`
	DefaultDescOrder  []string `json:"default_desc_order"`
	OpenOptionalOrder bool     `json:"open_optional_order"`
	OptionalOrder     []string `json:"optional_order"`
}

type Paging struct {
	Open     bool    `json:"open"`
	Must     bool    `json:"must"`
	MaxLimit float64 `json:"max_limit"`
}

func (r RestNodeOp) Name() string {
	return RestNodeType
}

type RestFieldOp struct {
	FieldQueryable   FieldQueryable   `json:"field_queryable"`
	FieldOperability FieldOperability `json:"field_operability"`
}

func (r RestFieldOp) Name() string {
	return RestFieldType
}

type FieldQueryable struct {
	EQ           GenRestSwitch `json:"EQ"`
	NEQ          GenRestSwitch `json:"NEQ"`
	GT           GenRestSwitch `json:"GT"`
	GTE          GenRestSwitch `json:"GTE"`
	LT           GenRestSwitch `json:"LT"`
	LTE          GenRestSwitch `json:"LTE"`
	IsNil        GenRestSwitch `json:"IsNil"`
	NotNil       GenRestSwitch `json:"NotNil"`
	EqualFold    GenRestSwitch `json:"EqualFold"`
	Contains     GenRestSwitch `json:"Contains"`
	ContainsFold GenRestSwitch `json:"ContainsFold"`
	HasPrefix    GenRestSwitch `json:"HasPrefix"`
	HasSuffix    GenRestSwitch `json:"HasSuffix"`
	In           GenRestSwitch `json:"In"`
	NotIn        GenRestSwitch `json:"NotIn"`
}

type FieldOperability struct {
	Selete GenRestSwitch
	Create GenRestSwitch
	Update GenRestSwitch
}

func PaseRestNodeOrderOp(m map[string]interface{}) Order {
	if _, ok := m[RestNodeType]; ok {
		b, err := json.Marshal(m[RestNodeType])
		if err != nil {
			panic(err.Error())
		}
		restNodeOp := RestNodeOp{}
		err = json.Unmarshal(b, &restNodeOp)
		if err != nil {
			panic(err.Error())
		}
		return restNodeOp.Order
	}
	return Order{}
}

func PaseRestFieldQueryOp(m map[string]interface{}) []string {
	if _, ok := m[RestFieldType]; ok {
		b, err := json.Marshal(m[RestFieldType])
		if err != nil {
			panic(err.Error())
		}
		op := RestFieldOp{}
		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err)
		}

		res := make([]string, 0, 0)
		if op.FieldQueryable.EQ == GenRestTrue {
			res = append(res, "EQ")
		}

		if op.FieldQueryable.NEQ == GenRestTrue {
			res = append(res, "NEQ")
		}

		if op.FieldQueryable.GT == GenRestTrue {
			res = append(res, "GT")
		}

		if op.FieldQueryable.GTE == GenRestTrue {
			res = append(res, "GTE")
		}

		if op.FieldQueryable.LT == GenRestTrue {
			res = append(res, "LT")
		}

		if op.FieldQueryable.LTE == GenRestTrue {
			res = append(res, "LTE")
		}

		if op.FieldQueryable.IsNil == GenRestTrue {
			res = append(res, "IsNil")
		}

		if op.FieldQueryable.NotNil == GenRestTrue {
			res = append(res, "NotNil")
		}

		if op.FieldQueryable.EqualFold == GenRestTrue {
			res = append(res, "EqualFold")
		}

		if op.FieldQueryable.Contains == GenRestTrue {
			res = append(res, "Contains")
		}

		if op.FieldQueryable.ContainsFold == GenRestTrue {
			res = append(res, "ContainsFold")
		}

		if op.FieldQueryable.HasPrefix == GenRestTrue {
			res = append(res, "HasPrefix")
		}

		if op.FieldQueryable.HasSuffix == GenRestTrue {
			res = append(res, "HasSuffix")
		}

		if op.FieldQueryable.In == GenRestTrue {
			res = append(res, "In")
		}

		if op.FieldQueryable.NotIn == GenRestTrue {
			res = append(res, "NotIn")
		}
		return res
	}
	return nil
}

func PaseRestFieldOperability(m map[string]interface{}, o string) string {
	if _, ok := m[RestFieldType]; ok {
		b, err := json.Marshal(m[RestFieldType])
		if err != nil {
			panic(err.Error())
		}
		op := RestFieldOp{}
		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err)
		}

		switch o {
		case "Selete":
			if op.FieldOperability.Selete == GenRestDefault || op.FieldOperability.Selete == GenRestTrue {
				return "true"
			}
		case "Update":
			if op.FieldOperability.Update == GenRestDefault || op.FieldOperability.Update == GenRestTrue {
				return "true"
			}
		case "Create":
			if op.FieldOperability.Create == GenRestDefault || op.FieldOperability.Create == GenRestTrue {
				return "true"
			}
		}
		return "false"
	}
	return "true"
}

func PaseRestNodePaging(m map[string]interface{}) Paging {
	if _, ok := m[RestNodeType]; ok {
		b, err := json.Marshal(m[RestNodeType])
		if err != nil {
			panic(err.Error())
		}
		op := RestNodeOp{}
		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err.Error())
		}
		return op.Paging
	}
	return Paging{}
}

func PaseRestNodeMethod(m map[string]interface{}) map[string]string {
	if _, ok := m[RestNodeType]; ok {
		b, err := json.Marshal(m[RestNodeType])
		if err != nil {
			panic(err.Error())
		}

		op := RestNodeOp{}

		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err.Error())
		}
		res := map[string]string{}

		if op.Method.Get == GenRestDefault || op.Method.Get == GenRestTrue {
			res["Get"] = "GET"
		}

		if op.Method.Create == GenRestDefault || op.Method.Create == GenRestTrue {
			res["Create"] = "POST"
		}

		if op.Method.Update == GenRestDefault || op.Method.Update == GenRestTrue {
			res["Update"] = "PUT"
		}

		if op.Method.Delete == GenRestDefault || op.Method.Delete == GenRestTrue {
			res["Delete"] = "DELETE"
		}
		return res
	}
	return map[string]string{
		"Get":    "GET",
		"Create": "POST",
		"Update": "PUT",
		"Delete": "DELETE",
	}
}

func SliceContains(l []string, s string) bool {
	for _, t := range l {
		if t == s {
			return true
		}
	}
	return false
}

func GetInclude(m map[string]map[string]interface{}, node string, l []string, res *[][]string) {
	edges, ok := m[node]
	if ok {
		for edge, _ := range edges {
			if SliceContains(l, edge) {
				if len(l) != 0 {
					*res = append(*res, l)
				}
				continue
			}
			tmpL := append(l, edge)
			GetInclude(m, edge, tmpL[:], res)
		}
	} else {
		if len(l) != 0 {
			*res = append(*res, l)
		}
	}
}

func PaseGraphInclude(g gen.Graph) map[string][]string {
	includeMap := make(map[string]map[string]interface{})
	for _, node := range g.Nodes {
		for _, edge := range node.Edges {
			has := PaseRestEdgeInclude(edge.Annotations)
			if has {
				if _, ok := includeMap[node.Name]; ok {
					includeMap[node.Name][edge.Type.Name] = struct{}{}
				} else {
					includeMap[node.Name] = map[string]interface{}{edge.Type.Name: struct{}{}}
				}
			}
		}
	}

	includes := make(map[string][][]string, 0)
	for node, _ := range includeMap {
		res := make([][]string, 0, 0)
		GetInclude(includeMap, node, []string{node}, &res)
		includes[node] = res
	}

	includeK := make(map[string][]string, 0)
	for node, include := range includes {
		tmpM := make(map[string]interface{}, 0)
		for _, i := range include {
			s := strings.Join(i, ".")
			tmpM[s] = struct{}{}
		}
		tmpK := make([]string, 0, 0)
		for k, _ := range tmpM {
			tmpK = append(tmpK, k)
		}

		fmt.Println(tmpK)
		includeK[node] = tmpK
	}
	return includeK
}

func PaseRestEdgeInclude(m map[string]interface{}) bool {
	if _, ok := m[RestEdgeType]; ok {
		b, err := json.Marshal(m[RestEdgeType])
		if err != nil {
			log.Fatalln(err.Error())
		}
		op := RestEdgeOp{}

		err = json.Unmarshal(b, &op)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if op.Include == GenRestFalse {
			return false
		}
	}
	return true
}

func PaseRestEdgeMethod(m map[string]interface{}) map[string]string {
	if _, ok := m[RestEdgeType]; ok {
		b, err := json.Marshal(m[RestEdgeType])
		if err != nil {
			panic(err.Error())
		}

		op := RestEdgeOp{}

		err = json.Unmarshal(b, &op)
		if err != nil {
			panic(err.Error())
		}
		res := map[string]string{}

		if op.Method.Get == GenRestDefault || op.Method.Get == GenRestTrue {
			res["Get"] = "GET"
		}

		if op.Method.Create == GenRestDefault || op.Method.Create == GenRestTrue {
			res["Create"] = "POST"
		}

		if op.Method.Delete == GenRestDefault || op.Method.Delete == GenRestTrue {
			res["Delete"] = "DELETE"
		}
		return res
	}
	return map[string]string{
		"Get":    "GET",
		"Create": "POST",
		"Delete": "DELETE",
	}
}

type EnumData struct {
	Has    bool
	Values string
}

func PaseFieldIsEnum(f *gen.Field) EnumData {
	return EnumData{
		Has:    f.IsEnum(),
		Values: strings.Join(f.EnumValues(), ","),
	}
}
