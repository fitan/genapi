package gen_mgr

import (
	"encoding/json"
	"entgo.io/ent/entc/gen"
	"github.com/davecgh/go-spew/spew"
	gen_apiV2 "github.com/fitan/genapi/pkg/gin_api"
	"log"
	"path"
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
	"CheckMethodHasSwitch":     CheckMethodHasSwitch,
	"PaseRestEdgeMethod":       PaseRestEdgeMethod,
	"PaseFieldIsEnum":          PaseFieldIsEnum,
	"PaseRestEdgeInclude":      PaseRestEdgeInclude,
	"PaseGraphInclude":         PaseGraphInclude,
	"IncludesTo":               IncludesTo,
	"PaseRelType":              PaseRelType,
	"SliceHasKey":              SliceHasKey,
	"PaseFieldsOrderOp":        PaseFieldsOrderOp,
	"Join":                     strings.Join,
	"ForMat":                   GenForMat,
	"FuncImportUnique":         FuncImportUnique,
	"Dump":                     Dump,
	"PathJoin":                 path.Join,
}

func FuncImportUnique(fs []gen_apiV2.Func) string {
	importList := make([]string, 0, 0)
	for _, f := range fs {
		if f.Plugins.CallBack.Has {
			if !SliceContains(importList, f.Plugins.CallBack.Template.ImportPath) {
				importList = append(importList, f.Plugins.CallBack.Template.ImportPath)
			}
		}
		for _, p := range f.Plugins.Point {
			if p.Has {
				if !SliceContains(importList, p.BindAfter.ImportPath) {
					importList = append(importList, p.BindAfter.ImportPath)
				}

				if !SliceContains(importList, p.BindBefor.ImportPath) {
					importList = append(importList, p.BindBefor.ImportPath)
				}
			}
		}
	}
	return strings.Join(importList, "\n")
}

func SliceHasKey(l []string, k string) bool {
	for _, v := range l {
		if v == k {
			return true
		}
	}
	return false
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
	Get    EdgeMethodOp
	Create EdgeMethodOp
	Delete EdgeMethodOp
}
type EdgeMethodOp struct {
	Has       GenRestSwitch `json:"has"`
	RouterTag string        `json:"router_tag"`
	Comments  []string      `json:"comments"`
}

type GenRestSwitch int

const GenRestDefault GenRestSwitch = 0
const GenRestTrue GenRestSwitch = 1
const GenRestFalse GenRestSwitch = 2

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
	Order        GenRestSwitch `json:"Order"`
}

type FieldOperability struct {
	Selete GenRestSwitch
	Create GenRestSwitch
	Update GenRestSwitch
}

type OrderOp struct {
	OrderField []string
	Has        bool
}

func PaseFieldsOrderOp(fs []*gen.Field) OrderOp {
	orderField := make([]string, 0, 0)
	has := false
	for _, f := range fs {
		a := PaseRestFieldQueryOp(f.Annotations)

		if SliceHasKey(a, "Order") {
			orderField = append(orderField, f.Name)
			has = true
		}
	}
	return OrderOp{
		OrderField: orderField,
		Has:        has,
	}
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

		if op.FieldQueryable.Order == GenRestTrue {
			res = append(res, "Order")
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

func IncludesTo(include []string, symbol string) string {
	tmp := make([]string, 0, len(include))
	for _, v := range include {
		tmpInclude := strings.Split(v, ".")[1:]
		if len(tmpInclude) != 0 {
			tmp = append(tmp, strings.Join(tmpInclude, "."))
		}
	}
	return strings.Join(tmp, symbol)
}

func PaseGraphInclude(g gen.Graph) map[string][]string {
	includeMap := make(map[string]map[string]interface{})
	for _, node := range g.Nodes {
		for _, edge := range node.Edges {
			has := PaseRestEdgeInclude(edge.Annotations)
			if has {
				if _, ok := includeMap[Snake(node.Name)]; ok {
					includeMap[Snake(node.Name)][Snake(edge.Type.Name)] = struct{}{}
				} else {
					includeMap[Snake(node.Name)] = map[string]interface{}{Snake(edge.Type.Name): struct{}{}}
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

func PaseRestEdgeMethod(m map[string]interface{}) EdgeMethod {
	op := RestEdgeOp{}
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
		return op.Method
	}
	return op.Method
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

func Snake(s string) string {
	return gen.Funcs["snake"].(func(string) string)(s)
}

func PaseRelType(e *gen.Edge) struct {
	Src  string
	Dest string
} {
	if len(e.Rel.Type.String()) != 3 {
		log.Fatalf("Edge Name %v Not Find X2X", e.Name)
	}
	var res struct {
		Src  string
		Dest string
	}
	res.Dest = e.Rel.Type.String()[len(e.Rel.Type.String())-1:]
	res.Src = e.Rel.Type.String()[0:1]
	return res
}

type RestNodeOp struct {
	Paging Paging     `json:"paging"`
	Order  Order      `json:"order"`
	Method NodeMethod `json:"method"`
}

func (r RestNodeOp) Name() string {
	return RestNodeType
}

type NodeMethod struct {
	GetOne     NodeMethodOp
	GetList    NodeMethodOp
	CreateOne  NodeMethodOp
	CreateList NodeMethodOp
	UpdateOne  NodeMethodOp
	UpdateList NodeMethodOp
	DeleteOne  NodeMethodOp
	DeleteList NodeMethodOp
}

type NodeMethodOp struct {
	Has       GenRestSwitch `json:"has"`
	RouterTag string        `json:"router_tag"`
	Comments  []string      `json:"comments"`
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

func CheckMethodHasSwitch(s GenRestSwitch) bool {
	if s == GenRestDefault || s == GenRestTrue {
		return true
	}
	return false
}

func PaseRestNodeMethod(m map[string]interface{}) NodeMethod {
	op := RestNodeOp{}
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
		return op.Method

	}

	return op.Method
}

type TempString string

func (s TempString) Format(data interface{}) (out string, err error) {
	t := template.Must(template.New("").Parse(string(s)))
	builder := &strings.Builder{}
	if err = t.Execute(builder, data); err != nil {
		return
	}
	out = builder.String()
	return
}

func GenForMat(s string, data interface{}) string {
	res, _ := TempString(s).Format(data)
	return res
}

func Dump(i interface{}) string {
	spew.Dump(i)
	return ""
}
