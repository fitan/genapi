package public

import (
	"github.com/spf13/viper"
	"log"
)

type GenConf struct {
	Plugin []Plugin `json:"plugin"`
	Gen    Gen      `json:"gen"`
	BaseConf BaseConf `json:"baseConf"`
}
type BindBefor struct {
	ImportPath string `json:"importPath"`
	Template   string `json:"template"`
}
type BindAfter struct {
	ImportPath string `json:"importPath"`
	Template   string `json:"template"`
}
type InterfaceName struct {
	Name      string    `json:"name"`
	BindBefor BindBefor `json:"bindBefor"`
	BindAfter BindAfter `json:"bindAfter"`
}
type Plugin struct {
	Name          string          `json:"name"`
	InterfaceName []InterfaceName `json:"interfaceName"`
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

type BaseConf struct {
	Wrap struct {
		ImportPath string `json:"importPath"`
		WrapFunc   string `json:"wrapFunc"`
		WrapRes string `json:"wrapRes"`
	} `json:"wrap"`
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


