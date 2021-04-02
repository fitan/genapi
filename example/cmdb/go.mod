module cmdb

go 1.16

require (
	entgo.io/ent v0.7.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/fitan/genapi v0.0.0
	github.com/gin-contrib/logger v0.0.2 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/rs/zerolog v1.21.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/fitan/genapi v0.0.0 => ../../
