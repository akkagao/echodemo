package conf

import (
	"github.com/BurntSushi/toml"
)

type TomlConfig struct {
	Runmodel   string
	ApiVersion string
	Port       string
	Log        log
	Database   database
}

type database struct {
	Db_host string
	Db_port int
	Db_user string
	Db_pass string
	Db_name string
	Db_type string
}

type log struct {
	LogPath      string
	LogFileName  string
	PrintConsole bool
	MaxNumber    int32
	MaxSize      int64
	Level        int32
}

//读取环境变量
var Config TomlConfig

func InitConf(runmodel string) {
	if _, err := toml.DecodeFile("conf/conf_"+runmodel+".toml", &Config); err != nil {
		panic("program argument error,is only can use one of [dev,test,pro]")
		return
	}
}

func InitTestConf(runmodel string) {
	if _, err := toml.DecodeFile("/Users/crazywolf/go/gopath/src/echodemo/conf/conf_"+runmodel+".toml", &Config); err != nil {
		panic("program argument error,is only can use one of [dev,test,pro]")
		return
	}
}
