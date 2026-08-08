package sv

import (
	"Order/model"
	"Order/sv/mq"
	"Order/sv/redis"
	"Order/utils"
	"context"
	"encoding/json"
	"log"
	"time"
)

func CreateInvoice(order model.OrderData) model.InvoiceModel {
	invoice := model.InvoiceModel{Id: utils.SnowflakeSample.GenerateID(), OrderData: order, CreatedAt: time.Now()}
	return invoice
}

func OrderHandler(ctx context.Context, orderData model.OrderData) error {
	err := redis.ReduceStock(ctx, orderData)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	invoiceModel := CreateInvoice(orderData)
	jsonData, err := json.Marshal(invoiceModel)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = mq.RabbitSample.PublishSample(jsonData)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
