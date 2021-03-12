package pkg

import (
	"bytes"
	_ "embed"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"log"
	"path"
	"path/filepath"
	"text/template"
)

//go:embed internal/templateV2/predicateV2.tmpl
var predicateV2_tmpl string

//go:embed internal/templateV2/curd.tmpl
var curd_tmpl string

//go:embed internal/templateV2/router.tmpl
var router_tmpl string

//go:embed internal/templateV2/new_obj.tmpl
var new_obj_tmpl string

//go:embed internal/templateV2/swagger_obj.tmpl
var swagger_obj_tmpl string

//go:embed internal/templateV2/swagger.tmpl
var swagger_tmpl string

//go:embed internal/templateV2/default_predicate.tmpl
var default_predicate_tmpl string

//go:embed internal/templateV2/tools.tmpl
var tools_tmpl string

//var Templates []*gen.Template
//
//func InitStart() {
//
//	Templates = []*gen.Template{
//		gen.MustParse(gen.NewTemplate("predicateV2").Funcs(gen.Funcs).Funcs(FM).Parse(predicateV2_tmpl)),
//		gen.MustParse(gen.NewTemplate("curd").Funcs(gen.Funcs).Funcs(FM).Parse(curd_tmpl)),
//		gen.MustParse(gen.NewTemplate("new_obj").Funcs(gen.Funcs).Funcs(FM).Parse(new_obj_tmpl)),
//		gen.MustParse(gen.NewTemplate("swag").Funcs(gen.Funcs).Funcs(FM).Parse(swagger_obj_tmpl)),
//		gen.MustParse(gen.NewTemplate("router_swagger").Funcs(gen.Funcs).Funcs(FM).Parse(swagger_tmpl)),
//		gen.MustParse(gen.NewTemplate("default_predicate").Funcs(gen.Funcs).Funcs(FM).Parse(default_predicate_tmpl)),
//		gen.MustParse(gen.NewTemplate("tools").Funcs(gen.Funcs).Funcs(FM).Parse(tools_tmpl)),
//
//	}
//}

type tmplMsg struct {
	Name string
	Text string
}

type GPacking struct {
	gen.Graph
	PkgName string
}

func (t *tmplMsg) NameFormat(s string) string {
	return fmt.Sprintf("%s_%s.go", s, t.Name)
}

func Load(schemaPath string, dest string) {
	nodeTmps := []tmplMsg{
		{
			Name: "predicate",
			Text: predicateV2_tmpl,
		},
		{
			Name: "curd",
			Text: curd_tmpl,
		},
		{
			Name: "swagger",
			Text: swagger_tmpl,
		},
		{
			"default_predicate",
			default_predicate_tmpl,
		},
	}
	gTmps := []tmplMsg{
		{
			Name: "new",
			Text: new_obj_tmpl,
		},
		{
			Name: "tools",
			Text: tools_tmpl,
		},
	}

	tpl, err := template.New("gen_ent").Funcs(gen.Funcs).Funcs(FM).Parse(router_tmpl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	tpl, err = tpl.Parse(pkg_name_tmpl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	tpl, err = tpl.Parse(swagger_obj_tmpl)
	if err != nil {
		log.Fatalln(err.Error())
	}

	g, err := entc.LoadGraph(schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}
	gPacking := GPacking{
		Graph:   *g,
		PkgName: path.Base(dest),
	}

	assets := assets{
		dirs: []string{
			filepath.Join(dest),
		},
	}

	for _, gTmp := range gTmps {
		parse, err := tpl.Parse(gTmp.Text)
		if err != nil {
			log.Fatalln(err.Error())
		}
		b := bytes.NewBuffer(nil)
		err = parse.Execute(b, gPacking)
		if err != nil {
			log.Fatalln(err.Error())
		}
		assets.files = append(assets.files, file{
			path:    filepath.Join(gPacking.Config.Target, dest, gTmp.Name+".go"),
			content: b.Bytes(),
		})

	}

	for _, node := range g.Nodes {

		tmpG := gen.Graph{
			Config: g.Config,
			Nodes:  []*gen.Type{node},
		}
		gPacking.Graph = tmpG

		for _, nodeTmp := range nodeTmps {
			parse, err := tpl.Parse(nodeTmp.Text)
			if err != nil {
				log.Fatalln(err.Error())
			}
			b := bytes.NewBuffer(nil)
			err = parse.Execute(b, gPacking)
			if err != nil {
				log.Fatalln(err.Error())
			}

			assets.files = append(assets.files, file{
				path:    filepath.Join(g.Config.Target, dest, nodeTmp.NameFormat(gen.Funcs["snake"].(func(string) string)(node.Name))),
				content: b.Bytes(),
			})
		}
	}

	if err := assets.write(); err != nil {
		log.Fatalln(err.Error())
	}
	err = assets.formatGo()
	if err != nil {
		log.Fatalln(err.Error())
	}
}
