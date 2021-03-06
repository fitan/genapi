package interfacepkg

type CasbinKeyser interface {
	GetCasbinKeys() []interface{}
}

type CasbinListKeyser interface {
	GetCasbinKeys() [][]interface{}
}

type RedisKeyer interface {
	GetRedisKey() string
}

type RedisKeyser interface {
	GetRedisKeys() []string
}