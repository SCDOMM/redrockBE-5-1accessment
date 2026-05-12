package dao

import (
	"GeneralConfig"
	"Storage/model"
	"checkserver/kitex_gen/checkserver/service"
	"errors"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDao() error {
	config := GeneralConfig.GetMySQLConfig()
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName + "?charset=" + config.Charset
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err1 := db.AutoMigrate(&model.ProductModel{})
	if err1 != nil {
		log.Println(err1.Error())
		return err1
	}
	err2 := db.AutoMigrate(&model.InvoiceModel{})
	if err2 != nil {
		log.Println(err2.Error())
		return err2
	}
	return nil
}

func StorageHandler(invoiceModel model.InvoiceModel) error {
	err := SearchInvoice(strconv.FormatInt(invoiceModel.Id, 10))
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
func CheckHandler(orderData *service.OrderData) error {
	invoice := model.InvoiceModel{}
	err0 := db.Model(&model.InvoiceModel{}).Where("user_id LIKE ?", "%"+strconv.Itoa(int(orderData.UserId))+"%").Or("product_id LIKE ?", "%"+strconv.Itoa(int(orderData.ProductId))+"%").Find(&invoice)
	if err0.Error != nil {
		log.Println(err0.Error)
		return err0.Error
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
