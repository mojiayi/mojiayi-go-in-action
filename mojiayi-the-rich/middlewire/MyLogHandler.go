package middlewire

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CostTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		costTime := time.Since(startTime)
		logrus.Info("uri=", ctx.Request.RequestURI+",costTime=", costTime.Milliseconds())
	}
}
