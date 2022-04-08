package middlewire

import (
	"mojiayi-the-rich/setting"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"
	"github.com/go-eden/routine"
)

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

		setting.MyLogger.Info("生成新traceId=", traceId)
	} else {
		traceId = localTraceId.Get().(string)
		setting.MyLogger.Info("从local storage中取得traceId=", traceId)
	}
	return traceId
}
