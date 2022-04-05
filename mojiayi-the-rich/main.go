package main

import (
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	initLog()
	initDependencyInjection()

	router := gin.Default()
	currencyV1 := router.Group("/v1/currency")
	{
		currencyV1.GET("/weight", service.CalculateWeight)
		currencyV1.GET("/goods", service.CalculatePurchaseAmount)
	}
	router.Run(":8080")
	logrus.Info("启动项目mojiayi-the-rich成功")
}

func initDependencyInjection() {
	container := config.LoadProjectConfig()

	container.Invoke(mapper.NewCurrencyInfoDao)
}

func initLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
}
