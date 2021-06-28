package public

import (
	"github.com/spf13/viper"
	"log"
)

type GenConf struct {
	BaseConf BaseConf `json:"baseConf"`
	Plugin   []Plugin `json:"plugin"`
	Gen      Gen      `json:"gen"`
}
type Wrap struct {
	ImportPath string `json:"importPath"`
	WrapFunc   string `json:"wrapFunc"`
	WrapRes    string `json:"wrapRes"`
}
type BaseConf struct {
	Wrap Wrap `json:"wrap"`
}
type Befor struct {
	ImportPath string `json:"importPath"`
	Template   string `json:"template"`
}
type After struct {
	ImportPath string `json:"importPath"`
	Template   string `json:"template"`
}
type InBindinterfaceName struct {
	Name  string `json:"name"`
	Befor Befor  `json:"Befor"`
	After After  `json:"After"`
}
type InBindInterfaceName struct {
	Name  string `json:"name"`
	After After  `json:"After"`
}
type OutBindInterfaceName struct {
	Name  string `json:"name"`
	After After  `json:"After"`
}
type Plugin struct {
	Name                 string                 `json:"name"`
	InBindInterfaceName  []InBindInterfaceName  `json:"inBindInterfaceName,omitempty"`
	OutBindInterfaceName []OutBindInterfaceName `json:"outBindInterfaceName,omitempty"`
}
type Ent struct {
	Name string `json:"name"`
	Src  string `json:"src"`
	Dest string `json:"dest"`
}
type API struct {
	Name string `json:"name"`
	Src  string `json:"src"`
	Dest string `json:"dest"`
}
type Gen struct {
	Ent []Ent `json:"ent"`
	API []API `json:"api"`
}

//
//type Ent struct {
//	Name string
//	Src string
//	Dest string
//}
//
//
//type Api struct {
//	Name string
//	Src string
//	Dest string
//}

var v *viper.Viper
var genConf *GenConf

func init() {
	v, genConf = ReadConf()
}

func GetViper() *viper.Viper {
	return v
}

func GetGenConf() *GenConf {
	return genConf
}

func ReadConf() (*viper.Viper, *GenConf) {
	v := viper.New()
	v.SetConfigName("gen")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	v.AddConfigPath("$HOME/.gen")
	err := v.ReadInConfig()
	if err != nil {
		log.Panicln(err)
	}

	genConf := &GenConf{}
	if err := v.Unmarshal(genConf); err != nil {
		log.Panic(err)
	}

	return v, genConf
}
