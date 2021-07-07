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

//go:embed internal/templateV3/predicate.tmpl
var predicate_tmpl string

//go:embed internal/templateV3/curd.tmpl
var curd_tmplV2 string

//go:embed internal/templateV3/func.tmpl
var func_tmpl string

//go:embed internal/templateV2/includes.tmpl
var includesV2_tmpl string

//go:embed internal/templateV3/option.tmpl
var option_tmpl string

//go:embed internal/templateV3/new_obj.tmpl
var new_obj string

//go:embed internal/templateV3/router.tmpl
var router_tmpl string

//go:embed internal/templateV3/swagger_obj.tmpl
var swagger_obj_tmpl string


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

type tmplMsgV2 struct {
	Name string
	Text string
}

type GPackingV2 struct {
	gen.Graph
	PkgName  string
	Includes map[string][]string
}

func (t *tmplMsgV2) NameFormat(s string) string {
	return fmt.Sprintf("%s_%s.go", s, t.Name)
}

func LoadV2(schemaPath string, dest string) {
	nodeTmps := []tmplMsgV2{
		{
			Name: "predicate",
			Text: predicate_tmpl,
		},
		{
			Name: "curd",
			Text: curd_tmplV2,
		},
		{
			Name: "option",
			Text: option_tmpl,
		},
		{
			Name: "func",
			Text: func_tmpl,
		},
	}
	gTmps := []tmplMsgV2{
		{
			Name: "new",
			Text: new_obj,
		},
		//{
		//	Name: "tools",
		//	Text: tools_tmpl,
		//},
		{
			Name: "includes",
			Text: includesV2_tmpl,
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
	gPacking := GPackingV2{
		Graph:    *g,
		PkgName:  path.Base(dest),
		Includes: PaseGraphInclude(*g),
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
