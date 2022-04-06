package main

import (
	"mojiayi-the-rich/config"
	"mojiayi-the-rich/middlewire"
	"mojiayi-the-rich/routers"
)

func main() {
	config.SetupDependencyInjection()

	middlewire.SetupLogOutput()

	router := routers.InitRouters()

	router.Run(":8080")
}
