package sv

import (
	"Check/model"
	"Check/sv/client"
	"log"
)

func CheckHandler(orderData model.OrderData) error {
	err := client.CheckHandler(orderData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
