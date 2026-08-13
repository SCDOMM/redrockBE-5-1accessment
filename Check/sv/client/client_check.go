package client

import (
	"Check/kitex_gen/checkserver/service"
	"Check/model"
	"context"
	"fmt"
	"log"
)

func CheckHandler(orderData model.OrderData) error {
	var kitexData service.OrderData
	kitexData.UserId = int32(orderData.UserId)
	kitexData.ProductId = int32(orderData.ProductId)

	cli := getKitexClient()
	if cli == nil {
		return fmt.Errorf("kitex client 尚未初始化")
	}
	err := cli.CheckOrder(context.Background(), &kitexData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
