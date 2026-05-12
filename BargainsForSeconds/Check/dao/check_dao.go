package dao

import (
	"Check/model"
	"GeneralConfig"
	"checkserver/kitex_gen/checkserver/service"
	"checkserver/kitex_gen/checkserver/service/checkservice"
	"context"
	"log"
	"strconv"

	"github.com/cloudwego/kitex/client"
)

var (
	kitexClient checkservice.Client
)

func init() {
	config := GeneralConfig.GetKitexConfig()
	kitexClient = checkservice.MustNewClient(config.ServerName, client.WithHostPorts(config.Host+strconv.Itoa(config.Port)))
}

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
