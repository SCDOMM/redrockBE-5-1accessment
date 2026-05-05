package main

import (
	"Order/cache"
	"Order/mq"
	"Order/router"
	"Order/sv"
	"Order/utils"
	"fmt"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = cache.InitRedis()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = mq.InitMqUrl()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = sv.InitRabbitMQ("test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	h := router.InitHertz()
	router.InitRouter(h)
	h.Spin()
}
