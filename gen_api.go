package genapi

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"log"
)

func EntApi() {
	err := entc.Generate("./schema", &gen.Config{
		Header: ``,
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("predicateV2").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/predicateV2.tmpl")),
			gen.MustParse(gen.NewTemplate("curd").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/curd.tmpl")),
			gen.MustParse(gen.NewTemplate("new_obj").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/new_obj.tmpl")),
			gen.MustParse(gen.NewTemplate("swag").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/swag.tmpl")),
			gen.MustParse(gen.NewTemplate("router_swagger").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/router_swagger.tmpl")),
			gen.MustParse(gen.NewTemplate("default_predicate").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/default_predicate.tmpl")),
			gen.MustParse(gen.NewTemplate("tools").Funcs(gen.Funcs).Funcs(FM).ParseFiles("template/tools.tmpl")),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
