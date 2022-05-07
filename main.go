package main

import (
	"fmt"
	"mojiayi-the-rich/routers"
	"mojiayi-the-rich/setting"
)

func main() {
	setting.Setup()

	router := routers.InitRouters()

	err := router.Run(":8080")
	if err != nil {
		fmt.Printf("启动失败,err=%v", err)
	}
}
