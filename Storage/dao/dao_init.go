package dao

import (
	"GeneralConfig"
	"Storage/model"
	"log"
	"strconv"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func init() {
	once.Do(func() {
		config := GeneralConfig.GetMySQLConfig()
		dsn := config.UserName + ":" + config.Password + "@tcp(" + config.Host + ":" + strconv.Itoa(config.Port) + ")/" + config.DbName + "?charset=" + config.Charset
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Println(err.Error())
		}
		err1 := db.AutoMigrate(&model.ProductModel{})
		if err1 != nil {
			log.Println(err1.Error())
		}
		err2 := db.AutoMigrate(&model.InvoiceModel{})
		if err2 != nil {
			log.Println(err2.Error())
		}
	})
}
