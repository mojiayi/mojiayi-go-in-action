package routers

import (
	"mojiayi-the-rich/middlewire"
	v1 "mojiayi-the-rich/routers/api/v1"
	"mojiayi-the-rich/setting"

	"github.com/gin-gonic/gin"
)

var (
	currencyInfoService   v1.CurrencyInfoService
	currencyWeightService v1.CurrencyWeightService
	purchaseGoodsService  v1.PurchaseGoodsService
)

func InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(setting.PutTraceIdIntoLocalStorage())
	router.Use(middlewire.RecordCostTime())

	currencyV1 := router.Group("/api/v1/currency")
	{
		currencyV1.GET("/weight", currencyWeightService.CalculateWeight)
		currencyV1.GET("/list", currencyInfoService.QueryAvailableCurrency)
		currencyV1.GET("/goods", purchaseGoodsService.CalculatePurchaseAmount)
	}

	return router
}
