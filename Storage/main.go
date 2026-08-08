package main

import (
	"Storage/mq"
	"checkserver"
	"checkserver/kitex_gen/checkserver/service/checkservice"
	"fmt"
)

func main() {
	mq.RabbitSample.ConsumeSample()

	svr := checkservice.NewServer(new(checkserver.CheckServiceImpl))
	err := svr.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
