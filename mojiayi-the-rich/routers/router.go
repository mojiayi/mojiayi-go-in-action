package routers

import (
	"mojiayi-the-rich/middlewire"
	v1 "mojiayi-the-rich/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {
	router := gin.New()

	router.Use(middlewire.PutTraceIdAsHeader())
	router.Use(middlewire.CostTime())

	currencyV1 := router.Group("/api/v1/currency")
	{
		currencyV1.GET("/weight", v1.CalculateWeight)
		currencyV1.GET("/goods", v1.CalculatePurchaseAmount)
	}

	return router
}
