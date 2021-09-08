module cmdb

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/asim/go-micro/plugins/client/http/v3 v3.0.0-20210907061356-440aa4a1ce13
	github.com/asim/go-micro/plugins/config/encoder/yaml/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/plugins/registry/memory/v3 v3.0.0-20210630062103-c13bb07171bc
	github.com/asim/go-micro/plugins/server/http/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v3 v3.0.0-20210907061356-440aa4a1ce13
	github.com/asim/go-micro/v3 v3.6.0
	github.com/casbin/casbin/v2 v2.28.1
	github.com/casbin/gorm-adapter/v3 v3.2.6
	github.com/fitan/genapi v0.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/go-redis/redis/v8 v8.10.0
	github.com/go-resty/resty/v2 v2.6.0
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rs/zerolog v1.21.0
	github.com/spf13/viper v1.8.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/ugorji/go v1.2.5 // indirect
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.22.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC3
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC3
	go.opentelemetry.io/otel/sdk v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.0.0-RC3
	golang.org/x/sys v0.0.0-20210603081109-ebe580a85c40 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/fitan/genapi v0.0.0 => ../../
