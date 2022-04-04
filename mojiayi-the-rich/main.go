package main

import (
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/service"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadProjectConfig()

	router := gin.Default()
	currencyV1 := router.Group("/v1/currency")
	{
		currencyV1.GET("/weight", service.CalculateWeight)
		currencyV1.GET("/goods", service.CalculatePurchaseAmount)
	}
	router.Run(":8080")
}
