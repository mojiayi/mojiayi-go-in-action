package middlewire

import (
	"mojiayi-the-rich/constants"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
)

func PutTraceIdAsHeader() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		GetTraceId(ctx)

		ctx.Next()
	}
}

func GetTraceId(ctx *gin.Context) string {
	traceId := ctx.Request.Header.Get(constants.TRACE_ID)
	if len(traceId) == 0 {
		traceId = uuid.New()
		traceId = strings.ReplaceAll(traceId, "-", "")
		ctx.Request.Header.Set(constants.TRACE_ID, traceId)
		MyLogger.Info("生成新traceId=", traceId)
	} else {
		MyLogger.Info("从header中取得traceId=", traceId)
	}
	return traceId
}
