package setting

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var MySQLSetting = &MySQLConfig{}
var LogOutSetting = &LogOutputConfig{}
var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("setting/my.ini")
	if err != nil {
		fmt.Println("failed while load setting file setting/my.ini,err: ", err)
	}

	mapTo("mysql", MySQLSetting)
	mapTo("log", LogOutSetting)
}

type MySQLConfig struct {
	IP       string
	Port     int
	User     string
	Password string
	Database string
}

type LogOutputConfig struct {
	Dir string
}

func mapTo(section string, value interface{}) {
	err := cfg.Section(section).MapTo(value)
	if err != nil {
		fmt.Println("failed while cfg.MapTo "+section+",err: ", err)
	}
}
