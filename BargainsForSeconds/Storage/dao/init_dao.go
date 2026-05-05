package dao

import (
	"Storage/model"
	"Storage/utils"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDao() error {
	config := utils.GetMySQLConfig()
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName + "?charset=" + config.Charset
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
		return err
	}
	err1 := db.AutoMigrate(&model.ProductModel{})
	if err1 != nil {
		log.Fatal(err1.Error())
		return err1
	}
	err2 := db.AutoMigrate(&model.InvoiceModel{})
	if err2 != nil {
		log.Fatal(err2.Error())
		return err2
	}
	return nil
}
