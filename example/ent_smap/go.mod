module ent_samp

go 1.14

require (
	entgo.io/ent v0.6.0
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/fitan/genapi v0.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.1-0.20200311113236-681ffa848bae
	github.com/pyroscope-io/pyroscope v0.0.26 // indirect
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
)

replace github.com/fitan/genapi => ../../../genapi
