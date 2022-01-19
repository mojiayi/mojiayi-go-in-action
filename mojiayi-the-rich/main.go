package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-basic/uuid"

	"mojiayi-the-rich/param"
	"mojiayi-the-rich/service"
	"mojiayi-the-rich/vo"
)

func main() {
	router := gin.Default()
	currency := router.Group("/v1/currency")
	{
		currency.GET("/weight", func(context *gin.Context) {
			traceId := uuid.New()
			name := context.Query("name")
			if len(name) == 0 {
				wrapperErrorResp(http.StatusForbidden, "货币名称不能为空", traceId, context)
				return
			}
			amountStr := context.Query("amount")
			if len(amountStr) == 0 {
				wrapperErrorResp(http.StatusForbidden, "货币金额不能为空", traceId, context)
				return
			}

			amount, err := strconv.ParseInt(amountStr, 10, 64)
			if err != nil {
				wrapperErrorResp(http.StatusForbidden, "货币金额只能是数字", traceId, context)
				return
			}
			if amount <= 0 {
				wrapperErrorResp(http.StatusForbidden, "货币金额必须大于0", traceId, context)
				return
			}

			unitStr := context.Query("unit")
			if len(unitStr) == 0 {
				wrapperErrorResp(http.StatusForbidden, "货币单位不能为空", traceId, context)
				return
			}

			unit, err := strconv.ParseInt(unitStr, 10, 64)
			if err != nil {
				wrapperErrorResp(http.StatusForbidden, "货币单位只能是数字", traceId, context)
				return
			}
			if unit <= 0 {
				wrapperErrorResp(http.StatusForbidden, "货币单位必须大于0", traceId, context)
				return
			}

			var param param.CurrencyParam = *new(param.CurrencyParam)
			param.SetName(name)
			param.SetAmount(amount)
			param.SetUnit(unit)
			param.SetClientAgent(context.Request.UserAgent())
			param.SetClientIP(context.ClientIP())
			param.SetTimestamp(int64(time.Millisecond))
			param.SetTraceId(traceId)

			respData := service.CalculateWeight(param)
			wrapperSuccessResp(respData, traceId, context)
		})
	}
	router.Run(":8080")
}

func wrapperErrorResp(code int32, msg string, traceId string, context *gin.Context) {
	var resp vo.BaseVO = *new(vo.BaseVO)
	resp.SetCode(code)
	resp.SetMsg(msg)
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(traceId)
	context.JSON(int(code), resp)
}

func wrapperSuccessResp(data interface{}, traceId string, context *gin.Context) {
	var resp vo.BaseVO = vo.BaseVO{}
	resp.SetCode(http.StatusOK)
	resp.SetMsg("Success")
	resp.SetTimestamp(time.Now().UnixMilli())
	resp.SetTraceId(traceId)
	resp.SetData(data)

	context.JSON(http.StatusOK, resp)
}
