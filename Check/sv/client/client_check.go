package client

import (
	"Check/model"
	"checkserver/kitex_gen/checkserver/service"
	"checkserver/kitex_gen/checkserver/service/checkservice"
	"context"
	"log"
)

var (
	kitexClient checkservice.Client
)

func CheckHandler(orderData model.OrderData) error {
	var kitexData service.OrderData
	kitexData.UserId = int32(orderData.UserId)
	kitexData.ProductId = int32(orderData.ProductId)
	err := kitexClient.CheckOrder(context.Background(), &kitexData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
