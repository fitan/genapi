package main

import (
	"cmdb/pkg/httpclient"
	log2 "cmdb/pkg/log"
	"cmdb/pkg/log/zlog"
	"cmdb/pkg/trace"
	"cmdb/public"
	"cmdb/routers"
	"context"
	"fmt"
	_ "github.com/asim/go-micro/plugins/config/encoder/yaml/v3"
	"github.com/asim/go-micro/plugins/registry/memory/v3"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/server"
	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"
	"go.uber.org/zap"
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
	//s := selector.NewSelector(selector.Registry(r))


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



	z := zlog.NewZlog(zlog.WithTrace(trace.GetTp(), zerolog.InfoLevel))
	l := log2.NewXlog(log2.WithTrace(trace.GetTp(),zap.InfoLevel))

	//c := api.DefaultConfig()
	//c.Address = "consul.default.10.170.34.122.xip.io:8080"
	//registry := consul.NewRegistry(consul.Config(c))
	service := micro.NewService(
		micro.Server(srv),
	)
	service.Init()
	go func() {
		//c := http.NewClient(client.Selector(s),client.Wrap(opentracing.NewClientWrapper(jaegerTracer)))

		time.Sleep(3 * time.Second)
		for {
			ctx  := trace.GetTrCxt()

			zl := z.TraceLog(ctx, "zlog get")
			zl.Info().Str("an", "bo").Msg("zlog 第一次")
			zl.Error().Str("bo", "wei").Msg("zlog 第二次")
			zl.End()


			tl := l.TraceLog(zl.Context(),  "xlog get")
			tl.Info("第一次")
			time.Sleep(3 * time.Second)
			tl.Info("第二次")
			tl.Error("错误")

			tl.End()
			cli := httpclient.NewClient(httpclient.WithHost("http://www.baidu.com"), httpclient.WithTrace(trace.GetTp(),"call baidu",false), httpclient.WithTimeOut(3 * time.Second))
			req := cli.R()
			req.SetBody()
			req.Method = resty.MethodGet
			req.URL = "/abc"
			req.SetContext(tl.Context())
			req.Send()

			//request := c.NewRequest(SERVER_NAME, "/users", "",  client.WithContentType("application/json"))
			//response := new(map[string]interface{})
			//c.Call(context.TODO(), request, response)
			time.Sleep(5 * time.Second)
		}
	}()


	err = service.Run()
	if err != nil {
		fmt.Println(err)
	}
	//go routers.GetDefaultRouter().Run(public.GetConf().App.Host + ":" + public.GetConf().App.Port)



	//client.R().SetContext(ctx).Get("/users")

	//for {}
}
