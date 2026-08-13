package main

import (
	"Storage/kitex_gen/checkserver/service/checkservice"
	"Storage/mq"
	"Storage/registry"
	"fmt"
)

func main() {
	registry.InitRegister()
	defer registry.CloseRegistry()
	go mq.RabbitSample.ConsumeSample()

	svr := checkservice.NewServer(new(CheckServiceImpl))
	if err := svr.Run(); err != nil {
		fmt.Println("RPC Server 启动失败:", err)
		return
	}
}
