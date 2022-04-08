package main

import (
	"fmt"
	"mojiayi-the-rich/dao"
	"mojiayi-the-rich/routers"
	"mojiayi-the-rich/setting"
)

func main() {
	setting.Setup()

	dao.Setup()

	router := routers.InitRouters()

	err := router.Run(":8080")
	if err != nil {
		fmt.Errorf("启动失败,err=%v", err)
	}
}
