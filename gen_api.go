package genapi

import (
	_ "embed"
	"entgo.io/ent/entc/gen"
)

//go:embed template/predicateV2.tmpl
var predicateV2_tmpl string

//go:embed template/curd.tmpl
var curd_tmpl string

//go:embed template/new_obj.tmpl
var new_obj_tmpl string

//go:embed template/swag.tmpl
var swag_tmpl string

//go:embed template/router_swagger.tmpl
var router_swagger_tmpl string

//go:embed template/default_predicate.tmpl
var default_predicate_tmpl string

//go:embed template/tools.tmpl
var tools_tmpl string

var Templates []*gen.Template

func init() {

	Templates = []*gen.Template{
		gen.MustParse(gen.NewTemplate("predicateV2").Funcs(gen.Funcs).Funcs(FM).Parse(predicateV2_tmpl)),
		gen.MustParse(gen.NewTemplate("curd").Funcs(gen.Funcs).Funcs(FM).Parse(curd_tmpl)),
		gen.MustParse(gen.NewTemplate("new_obj").Funcs(gen.Funcs).Funcs(FM).Parse(new_obj_tmpl)),
		gen.MustParse(gen.NewTemplate("swag").Funcs(gen.Funcs).Funcs(FM).Parse(swag_tmpl)),
		gen.MustParse(gen.NewTemplate("router_swagger").Funcs(gen.Funcs).Funcs(FM).Parse(router_swagger_tmpl)),
		gen.MustParse(gen.NewTemplate("default_predicate").Funcs(gen.Funcs).Funcs(FM).Parse(default_predicate_tmpl)),
		gen.MustParse(gen.NewTemplate("tools").Funcs(gen.Funcs).Funcs(FM).Parse(tools_tmpl)),
		//gen.MustParse(gen.NewTemplate("predicateV2").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/predicateV2.tmpl")),
		//gen.MustParse(gen.NewTemplate("curd").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/curd.tmpl")),
		//gen.MustParse(gen.NewTemplate("new_obj").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/new_obj.tmpl")),
		//gen.MustParse(gen.NewTemplate("swag").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/swag.tmpl")),
		//gen.MustParse(gen.NewTemplate("router_swagger").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/router_swagger.tmpl")),
		//gen.MustParse(gen.NewTemplate("default_predicate").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/default_predicate.tmpl")),
		//gen.MustParse(gen.NewTemplate("tools").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/tools.tmpl")),
	}
}
