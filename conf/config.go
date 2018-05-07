package conf

import (
	"testgopb/conf"
)

type Configs struct {
	conf.ConfigServiceMicro	`yaml:"microservice"`
	DB DataBase	`yaml:"database"`
}

type DataBase struct {
	Host string 		`yaml:"host"`
	Port string 		`yaml:"port"`
	Username string		`yaml:"username"`
	Password string 	`yaml:"password"`
	DBname string 		`yaml:"dbname"`
}

var Config Configs

func init() {
	conf.LoadConf(&Config,"conf/config.yaml")
}
