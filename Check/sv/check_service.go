package sv

import (
	"Check/dao"
	"Check/model"
	"log"
)

func CheckHandler(orderData model.OrderData) error {
	err := dao.CheckHandler(orderData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
