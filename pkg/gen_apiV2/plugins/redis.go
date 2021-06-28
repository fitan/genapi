package plugins

import (
	"go/types"
	"log"
)

var s []string

const (
	RedisGetKeyName  = "RedisGetKey"
	RedisGetKeysName = "RedisGetKeys"
)

func GetRedisPluginTemplate(DocFields []string, inFieldType types.Type, outFieldType types.Type) {
	if len(DocFields) < 3 {
		log.Fatalln("@Redis Mark need 1 parse")
	}

	pt := PluginTemplate{Has: true}
	pt.Keys = map[string]string{"action": DocFields[2]}

}
