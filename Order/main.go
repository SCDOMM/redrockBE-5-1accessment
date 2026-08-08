package main

import (
	"Order/cache"
	"Order/router"
	"Order/sv"
	"fmt"
)

func main() {
	err := cache.InitRedis()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = sv.InitRabbitMQ("test")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	h := router.InitHertzServer()
	router.InitRouter(h)
	h.Spin()
}
