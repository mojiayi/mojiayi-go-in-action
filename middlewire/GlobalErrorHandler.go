package middlewire

import (
	"mojiayi-the-rich/setting"
	"mojiayi-the-rich/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recover(ctx *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			setting.MyLogger.Info(r)

			utils.ErrorResp(http.StatusForbidden, "哦豁，系统开小差了！", ctx)

			ctx.Abort()
		}
	}()
	ctx.Next()
}
