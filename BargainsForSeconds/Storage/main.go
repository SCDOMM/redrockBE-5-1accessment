package main

import (
	"Storage/dao"
	"Storage/mq"
	"Storage/rd"
	"checkserver"
	"checkserver/kitex_gen/checkserver/service/checkservice"
	"fmt"
)

func main() {

	err := dao.InitDao()
	if err != nil {
		fmt.Println("dao error!" + err.Error())
		return
	}

	err = rd.InitRedis()
	if err != nil {
		fmt.Println("redis error!" + err.Error())
		return
	}

	rabbitMQ, err := mq.NewRabbitMQSample("test")
	if err != nil {
		fmt.Println("mq error!" + err.Error())
		return
	}
	rabbitMQ.ConsumeSample()

	svr := checkservice.NewServer(new(checkserver.CheckServiceImpl))
	err = svr.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
