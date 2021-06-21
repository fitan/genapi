package interfacepkg

type CasbinKeyser interface {
	GetCasbinKeys() []interface{}
}

type CasbinListKeyser interface {
	GetCasbinKeys() [][]interface{}
}