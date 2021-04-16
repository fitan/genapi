module cmdb

go 1.16

require (
	entgo.io/ent v0.7.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/appleboy/gin-jwt/v2 v2.6.4
	github.com/chenjiandongx/ginprom v0.0.0-20201217063207-fe11b7f55a35
	github.com/cosmtrek/air v1.26.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fitan/genapi v0.0.0
	github.com/gin-contrib/logger v0.0.2
	github.com/gin-gonic/gin v1.7.1
	github.com/go-playground/validator/v10 v10.5.0 // indirect
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/pelletier/go-toml v1.9.0 // indirect
	github.com/prometheus/client_golang v1.10.0
	github.com/prometheus/common v0.20.0 // indirect
	github.com/rs/zerolog v1.21.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/ugorji/go v1.2.5 // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	golang.org/x/sys v0.0.0-20210415045647-66c3f260301c // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/ini.v1 v1.51.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/fitan/genapi v0.0.0 => ../../
