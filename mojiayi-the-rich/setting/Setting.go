package setting

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var MySQLSetting = &MySQLConfig{}
var LogOutSetting = &LogOutputConfig{}
var MyLogger *logrus.Logger
var MetadataLogger *logrus.Logger
var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("setting/my.ini")
	if err != nil {
		fmt.Println("failed while load setting file setting/my.ini,err: ", err)
	}

	mapTo("mysql", MySQLSetting)
	mapTo("log", LogOutSetting)

	setupLogOutput()
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

func setupLogOutput() {
	// 打印请求中业务日志
	MyLogger = initLog(LogOutSetting.Dir, "access.log")
	// 打印请求的元数据信息
	MetadataLogger = initLog(LogOutSetting.Dir, "metadata.log")
}

func initLog(path string, filename string) *logrus.Logger {
	log := logrus.New()
	log.Formatter = &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05.999999999",
	}
	filepath := path + filename
	var file *os.File
	var err error
	if _, err = os.Stat(filepath); os.IsNotExist(err) {
		file, err = os.Create(filepath)
	} else {
		file, err = os.OpenFile(filepath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	}

	if err != nil {
		fmt.Println("fail to open log file " + filepath)
	}

	log.Out = file
	log.Level = logrus.InfoLevel

	return log
}
