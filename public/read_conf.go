package public

import (
	"github.com/spf13/viper"
	"log"
)

type GenConf struct {
	BaseConf BaseConf `json:"BaseConf"`
	Plugin   Plugin   `json:"Plugin"`
	Gen      Gen      `json:"Gen"`
}
type WrapResult struct {
	ImportPath     string `json:"ImportPath"`
	WrapFunc       string `json:"WrapFunc"`
	WrapResultType string `json:"WrapResultType"`
}
type BaseConf struct {
	WrapResult WrapResult `json:"WrapResult"`
}
type Match struct {
	Param            []string `json:"Param"`
	OutInterfaceName []string `json:"OutInterfaceName"`
	InInterfaceName  []string `json:"InInterfaceName"`
}
type Cover struct {
	Match      Match  `json:"Match"`
	ImportPath string `json:"ImportPath"`
	Template   string `json:"Template"`
}
type CallBack struct {
	TagName string  `json:"TagName"`
	Cover   []Cover `json:"Cover"`
}
type MountBindBefor struct {
	ImportPath string `json:"ImportPath"`
	Template   string `json:"Template"`
}
type MountBindAfter struct {
	ImportPath string `json:"ImportPath"`
	Template   string `json:"Template"`
}
type Mount struct {
	Match          Match          `json:"Match"`
	MountBindBefor MountBindBefor `json:"MountBindBefor"`
	MountBindAfter MountBindAfter `json:"MountBindAfter"`
}
type Point struct {
	TagName string  `json:"TagName"`
	Mount   []Mount `json:"Mount"`
}
type Plugin struct {
	CallBack []CallBack `json:"CallBack"`
	Point    []Point    `json:"Point"`
}
type Ent struct {
	Name string `json:"Name"`
	Src  string `json:"Src"`
	Dest string `json:"Dest"`
}

type Ts struct {
	Name   string `json:"Name"`
	Src    string `json:"Src"`
	Dest   string `json:"Dest"`
	Prefix string `json:"Prefix"`
}
type API struct {
	Name string `json:"Name"`
	Src  string `json:"Src"`
	Dest string `json:"Dest"`
}
type Gen struct {
	Ent []Ent `json:"Ent"`
	API []API `json:"Api"`
	Ts  []Ts  `json:"Ts"`
}

var genConf *GenConf
var viperConf *viper.Viper

func init() {
	viperConf, genConf = ReadConf()
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
