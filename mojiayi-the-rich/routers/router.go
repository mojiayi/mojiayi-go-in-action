package routers

import (
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/service"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(middlewire.PutTraceIdAsHeader())
	router.Use(middlewire.CostTime())

	currencyV1 := router.Group("/api/v1/currency")
	{
		currencyV1.GET("/weight", service.CalculateWeight)
		currencyV1.GET("/goods", service.CalculatePurchaseAmount)
	}

	return router
}
