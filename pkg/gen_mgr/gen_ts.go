package gen_mgr

import (
	"bytes"
	_ "embed"
	"entgo.io/ent/entc/gen"
	gen_apiV2 "github.com/fitan/genapi/pkg/gin_api"
	"github.com/fitan/genapi/public"
	"log"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

//go:embed ts_template/client.tmpl
var gen_ts_client_tmpl string

func genTs(apiMap map[string]*gen_apiV2.FileContext, baseConf public.BaseConf, dest string) {
	parse, err := template.New("gen_ts_client").Funcs(gen.Funcs).Funcs(FM).Parse("")
	if err != nil {
		log.Panicln(err.Error())
	}
	if err != nil {
		log.Fatalln(err.Error())
	}
	assets := assets{
		dirs: []string{
			filepath.Join(dest),
		},
	}

	for fileName, fileContext := range apiMap {
		tpl, err := parse.Parse(gen_ts_client_tmpl)
		if err != nil {
			log.Fatalln(err.Error())
		}
		b := bytes.NewBuffer(nil)
		err = tpl.Execute(b, struct {
			PkgName string
			Funcs   []gen_apiV2.Func
			BaseConf public.BaseConf

		}{
			PkgName: path.Base(dest),
			Funcs:   fileContext.Funcs,
			BaseConf: baseConf,
		})

		if err != nil {
			log.Fatalln(err.Error())
		}
		assets.files = append(assets.files, file{
			path:    filepath.Join(dest, strings.Replace(path.Base(fileName),".go", ".ts", -1)),
			content: b.Bytes(),
		})
	}

	//tpl, err := parse.New("register").Parse(register_tmplV2)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//b := bytes.NewBuffer(nil)
	//err = tpl.Execute(b, struct {
	//	PkgName      string
	//	ApiMap       map[string]*gen_apiV2.FileContext
	//	ReginsterMap map[string][]gen_apiV2.Func
	//	BaseConf public.BaseConf
	//}{
	//	PkgName:      path.Base(dest),
	//	ApiMap:       apiMap,
	//	ReginsterMap: ReginsterMap,
	//	BaseConf: baseConf,
	//})
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}

	if err := assets.write(); err != nil {
		log.Fatalln(err.Error())
	}

	//err = assets.formatGo()
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
}

func GenTs(src, dest string)  {
	context := gen_apiV2.NewApiContext()
	context.Load(src)
	context.Parse()
	for _, file := range context.Files {
		if len(file.Funcs) != 0 {
			genTs(context.Files, public.GetGenConf().BaseConf,dest)
			break
		}
	}
}

