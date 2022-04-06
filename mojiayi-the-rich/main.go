package main

import (
	"fmt"
	"io/fs"
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/dao/mapper"
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	initLog()
	initDependencyInjection()

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
}

func initLog() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)
	dir := "d://data//weblog//mojiayi-the-rich"
	os.Mkdir(dir, fs.ModePerm)
	file, err := os.Create(dir + "//info.log")
	if err != nil {
		fmt.Println("初始化日志输出配置失败")
		os.Exit(1)
	}
	logrus.SetOutput(file)
}
