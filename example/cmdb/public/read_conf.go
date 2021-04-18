package public

import (
	"gopkg.in/ini.v1"
)

type Conf struct {
	App   App   `ini:"app"`
	Mysql Mysql `ini:"mysql"`
	Resty Resty `ini:"resty"`
	Log   Log   `ini:"log"`
	Jwt   Jwt   `ini:"jwt"`
}

type App struct {
	Name string `ini:"name"`
	Host string `ini:"host"`
	Port string `ini:"port"`
}

type Mysql struct {
	Addr  string `ini:"addr"`
	Debug bool   `ini:"debug"`
}

type Resty struct {
	Debug bool `ini:"debug"`
}

type Log struct {
	Dir        string `ini:"dir"`
	MaxSize    int    `ini:"max_size"`
	MaxBackups int    `ini:"max_backups"`
	MaxAge     int    `ini:"max_age"`
	Compress   bool   `ini:"compress"`
}

type Jwt struct {
	Realm       string `ini:"realm"`
	IdentityKey string `ini:"identity_key"`
	SecretKey   string `ini:"secret_key"`
	Timeout     string `ini:"timeout"`
	MaxRefresh  string `ini:"max_refresh"`
}

var readConf *Conf

func NewConf() (*Conf, error) {
	conf := new(Conf)
	err := ini.MapTo(conf, "./conf/config.ini")
	return conf, err
}

func init() {
	conf, err := NewConf()
	if err != nil {
		XLog.Fatal().Err(err).Msg("")
		return
	}
	readConf = conf
}

func GetConf() *Conf {
	if readConf == nil {
		conf, err := NewConf()
		if err != nil {
			XLog.Fatal().Err(err).Msg("")
			return nil
		}

		readConf = conf
	}
	return readConf
}
