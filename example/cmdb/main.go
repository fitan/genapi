package main

import (
	"cmdb/pkg/trace"
	"cmdb/public"
	"cmdb/routers"
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/client/http/v3"
	_ "github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	"github.com/asim/go-micro/plugins/registry/memory/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/client"
	"github.com/asim/go-micro/v3/selector"
	"github.com/asim/go-micro/v3/server"
	"log"
	"os"
	"time"
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

	jaegerTracer, closer, err := trace.NewJaegerTracer("micro-go","10.170.34.122:6831")
	if err != nil {
		log.Fatalln(err)
	}
	defer closer.Close()



	//tr := trace.GetTr()
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//ctx, span := tr.Start(ctx, "call")
	//span.End()
	//
	//client := httpclient.NewClient(httpclient.WithDebug(true),httpclient.WithTraceContext("call baidu ", ctx),httpclient.WithHost("http://localhost:8080"))
	r := memory.NewRegistry()
	s := selector.NewSelector(selector.Registry(r))


	srv := httpServer.NewServer(
		server.Name(SERVER_NAME),
		server.Registry(r),
		server.Address(":"+public.GetConf().App.Port),
		server.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
	)

	hd := srv.NewHandler(routers.GetDefaultRouter())
	if err := srv.Handle(hd); err != nil {
		log.Fatalln(err)
	}


	//c := api.DefaultConfig()
	//c.Address = "consul.default.10.170.34.122.xip.io:8080"
	//registry := consul.NewRegistry(consul.Config(c))
	service := micro.NewService(
		micro.Server(srv),
	)
	service.Init()
	go func() {
		err := service.Run()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//go routers.GetDefaultRouter().Run(public.GetConf().App.Host + ":" + public.GetConf().App.Port)


	c := http.NewClient(client.Selector(s),client.Wrap(opentracing.NewClientWrapper(jaegerTracer)))

	time.Sleep(3 * time.Second)
	for {
		request := c.NewRequest(SERVER_NAME, "/users", "",  client.WithContentType("application/json"), client.wit)
		response := new(map[string]interface{})
		c.Call(context.TODO(), request, response)
		time.Sleep(5 * time.Second)
	}


	//client.R().SetContext(ctx).Get("/users")

	//for {}
}
