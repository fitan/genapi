package public

import (
	"github.com/spf13/viper"
	"gopkg.in/ini.v1"
	"sync"
)

var confLock sync.Mutex

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


func GetConf() *Conf {
	if readConf == nil {
		confLock.Lock()
		defer  confLock.Unlock()
		conf, err := NewConf()
		if err != nil {
			GetXLog().Error().Err(err).Msgf("")
		}

		readConf = conf
	}
	return readConf
}

func ReadYamlConf(confName string, obj interface{},paths ...string) error {
	v := viper.New()
	v.SetConfigName(confName)
	v.SetConfigType("yaml")
	for _, path := range paths {
		v.AddConfigPath(path)
	}
	err := v.ReadInConfig()
	if err != nil {
		GetXLog().Error().Err(err)
		return err
	}

	if err = v.Unmarshal(obj); err != nil {
		GetXLog().Error().Err(err)
		return err
	}

	return nil
}
