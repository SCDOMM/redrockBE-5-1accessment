package dao

import (
	"Check/model"
	"Check/utils"
	"Storage/kitex_gen/storage/service"
	"Storage/kitex_gen/storage/service/storageservice"
	"context"
	"log"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

var kitexClient storageservice.Client

func InitKitexClient(config utils.KitexConfig) {
	kitexClient = storageservice.MustNewClient(config.ServerName, client.WithHostPorts(config.Host+strconv.Itoa(config.Port)))
}

func CheckHandler(orderData model.OrderData) error {
	var kitexData service.OrderData
	kitexData.UserId = int32(orderData.UserId)
	kitexData.ProductId = int32(orderData.ProductId)
	err := kitexClient.CheckOrder(context.Background(), &kitexData)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
