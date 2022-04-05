package utils

import (
	"mojiayi-the-rich/constants"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/sirupsen/logrus"
)

func GetTraceId(context *gin.Context) (traceId string) {
	traceId = context.Request.Header.Get(constants.TRACE_ID)
	if len(traceId) == 0 {
		traceId := uuid.New()
		context.Request.Header.Set(constants.TRACE_ID, traceId)
		logrus.Info("上下文中没有traceId")
	} else {
		logrus.Info("直接从上下文中取得traceId")
	}
	return traceId
}
