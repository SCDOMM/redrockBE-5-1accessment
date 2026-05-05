package sv

import (
	"Storage/dao"
	"Storage/model"
	"encoding/json"
	"log"
)

func StorageMsgHandler(message []byte) error {
	invoiceModel := model.InvoiceModel{}
	err := json.Unmarshal(message, &invoiceModel)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err = dao.StorageHandler(invoiceModel)
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	return nil
}
