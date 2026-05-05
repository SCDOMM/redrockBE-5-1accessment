package main

import (
	"Storage/dao"
	"Storage/mq"
	"Storage/rd"
	"Storage/utils"
	"fmt"
)

func main() {
	err := utils.InitConfig()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = dao.InitDao()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = rd.InitRedis()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = mq.InitMqUrl()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	rabbitMQ, err := mq.NewRabbitMQSample("test")
	rabbitMQ.ConsumeSample()

}
