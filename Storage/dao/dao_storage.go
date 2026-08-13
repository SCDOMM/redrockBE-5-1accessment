package dao

import (
	"Storage/kitex_gen/checkserver/service"
	"Storage/model"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func StorageHandler(invoiceModel model.InvoiceModel) error {
	err := ReduceStock(int64(invoiceModel.OrderData.ProductId))
	if err != nil {
		log.Println(err)
		return err
	}

	err = SearchInvoice(strconv.FormatInt(invoiceModel.Id, 10))
	if err != nil {
		log.Println(err)
		return err
	}
	if result := db.Create(&invoiceModel); result.Error != nil {
		log.Println(result.Error)
		return result.Error
	}
	return nil
}
func CheckProduct() (*[]model.ProductModel, error) {
	var products []model.ProductModel
	if err := db.Model(&model.ProductModel{}).Find(&products).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return &products, nil
}

func CheckHandler(orderData *service.OrderData) error {
	invoice := model.InvoiceModel{}
	err := db.Model(&model.InvoiceModel{}).Where("user_id LIKE ?", "%"+strconv.Itoa(int(orderData.UserId))+"%").Or("product_id LIKE ?", "%"+strconv.Itoa(int(orderData.ProductId))+"%").Find(&invoice)
	if err.Error != nil {
		log.Println(err.Error)
		return err.Error
	}
	return nil
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

func ReduceStock(productId int64) error {
	productModel := model.ProductModel{}
	if result := db.Model(&model.ProductModel{}).Where("id = ?", productId).Find(&productModel); result.Error != nil {
		return result.Error
	}
	if productModel.Stock == 0 {
		return fmt.Errorf("product stock is zero")
	}
	if result := db.Model(&model.ProductModel{}).Where("id = ?", productId).Update("stock", productModel.Stock-1); result.Error != nil {
		return result.Error
	}
	return nil
}
