package main

import (
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/service"

	"github.com/gin-gonic/gin"
)

func main() {
	initDependencyInjection()

	middlewire.SetupLogOutput()

	router := gin.Default()

	currencyV1 := router.Group("/api/v1/currency").Use(middlewire.PutTraceIdAsHeader(), middlewire.CostTime())
	{
		currencyV1.GET("/weight", service.CalculateWeight)
		currencyV1.GET("/goods", service.CalculatePurchaseAmount)
	}
	router.Run(":8080")
}

func initDependencyInjection() {
	container := config.LoadProjectConfig()

	container.Invoke(mapper.NewCurrencyInfoDao)

	container.Invoke(config.LoadLogOutputConfig)
}
