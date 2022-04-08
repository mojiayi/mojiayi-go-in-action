package utils

import (
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/vo"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func IllegalArgumentErrorResp(msg string, context *gin.Context) {
	var resp vo.BaseVO = *new(vo.BaseVO)
	resp.SetCode(http.StatusBadRequest)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(middlewire.GetTraceId())
	context.JSON(http.StatusOK, resp)
}

func ErrorResp(code int32, msg string, context *gin.Context) {
	var resp vo.BaseVO = *new(vo.BaseVO)
	resp.SetCode(code)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(middlewire.GetTraceId())
	context.JSON(http.StatusOK, resp)
}

func SuccessResp(data interface{}, context *gin.Context) {
	var resp vo.BaseVO = vo.BaseVO{}
	resp.SetCode(http.StatusOK)
	resp.SetMsg(http.StatusText(http.StatusOK))
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(middlewire.GetTraceId())
	resp.SetData(data)
	context.JSON(http.StatusOK, resp)
}
