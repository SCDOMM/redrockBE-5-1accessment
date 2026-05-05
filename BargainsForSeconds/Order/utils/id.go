package utils

import (
	"Order/model"
	"strings"
	"time"
)

var (
	machineID = GetMachineId()
	sf        = NewSnowflake(machineID)
)

func DecodeID(str string) string {
	return strings.Split(str, ":")[2]
}
func CreateInvoice(order model.OrderData) model.InvoiceModel {
	invoice := model.InvoiceModel{Id: sf.GenerateID(), OrderData: order, CreateAt: time.Now()}
	return invoice
}
