module cmdb

go 1.16

require (
	entgo.io/ent v0.9.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/asim/go-micro/plugins/config/encoder/yaml/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/plugins/config/source/consul/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/plugins/registry/consul/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/plugins/server/http/v3 v3.0.0-20210824071433-49eccbc85a0f
	github.com/asim/go-micro/v3 v3.6.0
	github.com/casbin/casbin/v2 v2.28.1
	github.com/casbin/gorm-adapter/v3 v3.2.6
	github.com/fitan/genapi v0.0.0
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/go-redis/redis/v8 v8.10.0
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/hashicorp/consul/api v1.9.0
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/reactivex/rxgo/v2 v2.5.0
	github.com/rs/zerolog v1.21.0
	github.com/spf13/viper v1.8.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.5 // indirect
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/fitan/genapi v0.0.0 => ../../
