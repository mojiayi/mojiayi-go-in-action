package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetHeaderValue(headerKey string, context *gin.Context) (headerValue string) {
	headerValue = context.Request.Header.Get(headerKey)
	if len(headerValue) == 0 {
		logrus.Info("header中没有headerKey=", headerKey)
	} else {
		logrus.Info("从header中取得headerKey=", headerKey, ",headerValue=", headerValue)
	}
	return headerValue
}
