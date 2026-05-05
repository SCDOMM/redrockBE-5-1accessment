package sv

import (
	"Order/cache"
	"Order/model"
	"Order/mq"
	"Order/utils"
	"context"
	"encoding/json"
	"log"
)

var rabbitMQ *mq.RabbitMQ

func InitRabbitMQ(queueName string) error {
	var err error
	rabbitMQ, err = mq.NewRabbitMQSample(queueName)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func OrderHandler(ctx context.Context, orderData model.OrderData) error {
	err := cache.ReduceStock(ctx, orderData)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	invoiceModel := utils.CreateInvoice(orderData)
	jsonData, err := json.Marshal(invoiceModel)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = rabbitMQ.PublishSample(jsonData)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
