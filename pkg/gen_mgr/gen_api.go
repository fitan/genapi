package gen_mgr

import (
	"bytes"
	_ "embed"
	"entgo.io/ent/entc/gen"
	"github.com/fitan/genapi/pkg/gin_api"
	"github.com/fitan/genapi/public"
	"github.com/marcinwyszynski/directory_tree"
	"log"
	"path"
	"path/filepath"
	"sync"
	"text/template"
)

//go:embed gin_api_template/handler.tmpl
var gen_api_tmplV2 string

//go:embed gin_api_template/register.tmpl
var register_tmplV2 string

//go:embed gin_api_template/transfer_register.tmpl
var transfer_register_tmpl string

//go:embed gin_api_template/transfer.tmpl
var transfer_tmpl string

//go:embed ent_fn_template/pkg_name.tmpl
var pkg_name_tmpl string

func genApiV2(apiMap map[string]*gen_apiV2.FileContext, ReginsterMap map[string][]gen_apiV2.Func, baseConf public.BaseConf, dest string) {
	parse, err := template.New("gen_api").Funcs(gen.Funcs).Funcs(FM).Parse(pkg_name_tmpl)
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
		tpl, err := parse.Parse(transfer_tmpl)
		//tpl, err := parse.Parse(gen_api_tmplV2)
		if err != nil {
			log.Fatalln(err.Error())
		}
		b := bytes.NewBuffer(nil)
		err = tpl.Execute(b, struct {
			PkgName  string
			Funcs    []gen_apiV2.Func
			BaseConf public.BaseConf
		}{
			PkgName:  path.Base(dest),
			Funcs:    fileContext.Funcs,
			BaseConf: baseConf,
		})

		if err != nil {
			log.Fatalln(err.Error())
		}
		assets.files = append(assets.files, file{
			path:    filepath.Join(dest, path.Base(fileName)),
			content: b.Bytes(),
		})
	}

	tpl, err := parse.New("register").Parse(transfer_register_tmpl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	b := bytes.NewBuffer(nil)
	err = tpl.Execute(b, struct {
		PkgName      string
		ApiMap       map[string]*gen_apiV2.FileContext
		ReginsterMap map[string][]gen_apiV2.Func
		BaseConf     public.BaseConf
	}{
		PkgName:      path.Base(dest),
		ApiMap:       apiMap,
		ReginsterMap: ReginsterMap,
		BaseConf:     baseConf,
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	assets.files = append(assets.files, file{
		path:    filepath.Join(dest, path.Base("register.go")),
		content: b.Bytes(),
	})

	//tpl1, err := parse.New("role_method").Parse(role_method)
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//
	//b = bytes.NewBuffer(nil)
	//err = tpl1.Execute(b, struct {
	//	PkgName string
	//	ReginsterMap map[string][]gen_apiV2.Func
	//}{
	//	PkgName: path.Base(dest),
	//	ReginsterMap: ReginsterMap,
	//})
	//if err != nil {
	//	log.Fatalln(err.Error())
	//}
	//assets.files = append(assets.files, file{
	//	path: filepath.Join(dest, path.Base("role_method.go")),
	//	content: b.Bytes(),
	//})

	if err := assets.write(); err != nil {
		log.Fatalln(err.Error())
	}

	err = assets.formatGo()
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func GenApi(src, dest string) {
	context := gen_apiV2.NewApiContext()
	context.Load(src)
	context.Parse(gen_apiV2.ParseOption{false})
	for _, file := range context.Files {
		if len(file.Funcs) != 0 {
			genApiV2(context.Files, context.ReginsterMap, public.GetGenConf().BaseConf, dest)
			break
		}
	}
}

func DepthGen(src, dest string, fn func(src, dest string)) {
	tree, err := directory_tree.NewTree(src)
	if err != nil {
		log.Panicln(err)
	}

	g := sync.WaitGroup{}
	depthGen(&g, tree, dest, fn)
	g.Wait()
}

func depthGen(g *sync.WaitGroup, tree *directory_tree.Node, dest string, fn func(src, dest string)) {
	//context := gen_apiV2.NewApiContext()
	//context.Load(tree.FullPath)
	//context.Parse()
	//for _, file := range context.Files {
	//	if len(file.Funcs) != 0 {
	//		genApiV2(context.Files, context.ReginsterMap, public.GetGenConf().BaseConf,Dir)
	//		break
	//	}
	//}
	g.Add(1)
	go func() {
		fn(tree.FullPath, dest)
		g.Done()
	}()

	for _, node := range tree.Children {
		if node.Info.IsDir {
			depthGen(g, node, path.Join(dest, node.Info.Name), fn)
		}
	}
}
