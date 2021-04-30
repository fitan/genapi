package routers

import (
	"cmdb/gen/handler"
)

func init()  {
	handler.RegisterAll(GetDefaultRouter())
}
