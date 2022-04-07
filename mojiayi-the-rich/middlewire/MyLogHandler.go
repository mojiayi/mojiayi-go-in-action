package middlewire

import (
	"fmt"
	"mojiayi-the-rich/setting"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var MyLogger *logrus.Logger
var MetadataLogger *logrus.Logger

func SetupLogOutput() {
	// 打印请求中业务日志
	MyLogger = initLog(setting.LogOutSetting.Dir, "access.log")
	// 打印请求的元数据信息
	MetadataLogger = initLog(setting.LogOutSetting.Dir, "metadata.log")
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

func CostTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		MetadataLogger.WithFields(logrus.Fields{
			"cost":   time.Since(startTime).Milliseconds(),
			"ip":     ctx.ClientIP(),
			"method": ctx.Request.Method,
			"uri":    ctx.Request.RequestURI,
			"usage":  "metadata",
		}).Info("request metadata")
	}
}
