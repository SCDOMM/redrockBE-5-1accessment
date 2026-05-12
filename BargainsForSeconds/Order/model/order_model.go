package model

import (
	"time"

	"gorm.io/gorm"
)

type ProductModel struct {
	Id    int     `gorm:"primary_key;auto_increment;not_null;unique"`
	Name  string  `gorm:"not null"`
	Price float64 `gorm:"not null"`
	Stock int     `gorm:"not null"`
}
type OrderData struct {
	UserId    int
	ProductId int
}
type InvoiceModel struct {
	Id        int64 `gorm:"primary_key;auto_increment;not_null;unique"`
	OrderData `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	gorm.DeletedAt
}
