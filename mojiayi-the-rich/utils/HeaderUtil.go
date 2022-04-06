package utils

import (
	"mojiayi-the-rich/middlewire"

	"github.com/gin-gonic/gin"
)

func GetHeaderValue(headerKey string, context *gin.Context) (headerValue string) {
	headerValue = context.Request.Header.Get(headerKey)
	if len(headerValue) == 0 {
		middlewire.MyLogger.Info("header中没有headerKey=", headerKey)
	} else {
		middlewire.MyLogger.Info("从header中取得headerKey=", headerKey, ",headerValue=", headerValue)
	}
	return headerValue
}
