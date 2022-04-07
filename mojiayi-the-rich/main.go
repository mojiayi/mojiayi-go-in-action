package main

import (
	"mojiayi-the-rich/dao"
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/routers"
	"mojiayi-the-rich/setting"
)

func main() {
	middlewire.SetupLogOutput()

	setting.Setup()

	dao.Setup()

	router := routers.InitRouters()

	router.Run(":8080")
}
