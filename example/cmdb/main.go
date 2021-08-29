package main

import (
	"cmdb/public"
	"cmdb/routers"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	"github.com/asim/go-micro/v3/config"
	"github.com/asim/go-micro/v3/config/reader"
	"github.com/asim/go-micro/v3/config/reader/json"
	"github.com/asim/go-micro/v3/config/source"
	"log"
	"os"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"

	_ "github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	consulSource "github.com/asim/go-micro/plugins/config/source/consul/v3"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/hashicorp/consul/api"
)

// go build -ldflags "-X main.GitCommitId=`git rev-parse HEAD` -X 'main.goVersion=$(go version)' -X 'main.gitHash=$(git show -s --format=%H)' -X 'main.buildTime=$(git show -s --format=%cd)'" -o main.exe version.go
var (
	gitHash   string
	gitTag    string
	buildTime string
	goVersion string
)

const SERVER_NAME = "cmdb"

type Cf struct {
	Base struct {
		Address string `json:"address"`
		Port    int    `json:"port"`
	} `json:"base"`
}

// @title cmdbapi
// @version 1.0
// @description RESTful API 文档.
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func main() {
	args := os.Args
	if len(args) == 2 && (args[1] == "--version" || args[1] == "-v") {
		fmt.Printf("Git Tag: %s \n", gitTag)
		fmt.Printf("Git Commit hash: %s \n", gitHash)
		fmt.Printf("Build TimeStamp: %s \n", buildTime)
		fmt.Printf("GoLang Version: %s \n", goVersion)
		return
	}
	if err := public.GetDB().Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	enc := yaml.NewEncoder()

	cs := consulSource.NewSource(consulSource.WithAddress("localhost:8500"), consulSource.WithPrefix(consulSource.DefaultPrefix+SERVER_NAME+"/"), consulSource.StripPrefix(true), source.WithEncoder(enc))
	//conf, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	conf, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	if err != nil {
		panic(err)
	}
	err = conf.Load(cs)
	if err != nil {
		panic(err)
	}

	fmt.Println("data", conf.Map())

	cf := new(Cf)

	w, err := conf.Watch()
	if err != nil {
		panic(err)
	}
	go func() {
		for {
			v, err := w.Next()
			if err != nil {
				public.GetXLog().Error().Err(err).Send()
			}
			v.Scan(cf)
			fmt.Println(cf)
		}
	}()

	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Address(":"+public.GetConf().App.Port),
	)

	hd := srv.NewHandler(routers.GetDefaultRouter())
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}

	c := api.DefaultConfig()
	c.Address = "localhost:8500"
	registry := consul.NewRegistry(consul.Config(c))
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(registry),
	)
	service.Init()
	service.Run()
	//routers.GetDefaultRouter().Run(public.GetConf().App.Host + ":" + public.GetConf().App.Port)
}
