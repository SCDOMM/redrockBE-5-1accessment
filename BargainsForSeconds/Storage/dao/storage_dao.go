package dao

import (
	"Storage/model"
	"errors"
	"log"
	"strconv"
)

func StorageHandler(invoiceModel model.InvoiceModel) error {
	err := SearchInvoice(strconv.FormatInt(invoiceModel.Id, 10))
	if err != nil {
		log.Fatal(err)
		return err
	}
	if result := db.Create(&invoiceModel); result.Error != nil {
		log.Fatal(result.Error)
		return result.Error
	}
	return nil
}
func ExternalHandler() {

}

func SearchInvoice(invoiceId string) error {
	var countResult int64
	if result := db.Model(&model.InvoiceModel{}).Where("id = ?", invoiceId).Count(&countResult); result.Error != nil {
		log.Println("dataBase error:", result.Error.Error())
		return result.Error
	}
	if countResult != 0 {
		log.Println("this invoice is exist!")
		return errors.New("this invoice is exist")
	}
	return nil
}
