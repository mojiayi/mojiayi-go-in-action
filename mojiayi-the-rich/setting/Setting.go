package setting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/go-eden/routine"
	"os"
	"strings"

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
	log.Formatter = &MyLogFormatter{}

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

type MyLogFormatter struct {
}

func (m *MyLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var buffer *bytes.Buffer
	if entry.Buffer != nil {
		buffer = entry.Buffer
	} else {
		buffer = &bytes.Buffer{}
	}

	var requestMetadata = make(map[string]interface{})
	for k, v := range entry.Data {
		requestMetadata[k] = v
	}
	str, _ := json.Marshal(requestMetadata)

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog = fmt.Sprintf("%s|%s|%s|%s|%s\n", timestamp, entry.Level, GetTraceId(), entry.Message, string(str))
	buffer.WriteString(newLog)
	return buffer.Bytes(), nil
}

var localTraceId = routine.NewLocalStorage()

func PutTraceIdAsHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GetTraceId()

		ctx.Next()
	}
}

func GetTraceId() string {
	var traceId = ""
	if localTraceId.Get() == nil {
		traceId = strings.ReplaceAll(uuid.New(), "-", "")
		localTraceId.Set(traceId)
	} else {
		traceId = localTraceId.Get().(string)
	}
	return traceId
}
