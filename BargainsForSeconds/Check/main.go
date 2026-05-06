package main

import (
	"Check/dao"
	"Check/router"
	"Check/utils"
	"fmt"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	config := utils.GetKitexConfig()
	dao.InitKitexClient(config)

	h := router.InitHertz()
	router.InitRouter(h)
	h.Spin()

}
