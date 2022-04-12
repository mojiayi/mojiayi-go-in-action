package utils

import (
	"mojiayi-the-rich/setting"
	"mojiayi-the-rich/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IllegalArgumentErrorResp(msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.SetCode(http.StatusBadRequest)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(setting.GetTraceId())
	context.JSON(http.StatusOK, resp)
}

func ErrorResp(code int, msg string, context *gin.Context) {
	var resp = *new(vo.BaseVO)
	resp.SetCode(code)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(setting.GetTraceId())
	context.JSON(http.StatusOK, resp)
}

func SuccessResp(data interface{}, context *gin.Context) {
	var resp = vo.BaseVO{}
	resp.SetCode(http.StatusOK)
	resp.SetMsg(http.StatusText(http.StatusOK))
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(setting.GetTraceId())
	resp.SetData(data)
	context.JSON(http.StatusOK, resp)
}
