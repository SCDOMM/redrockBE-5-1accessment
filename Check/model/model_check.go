package model

import (
	"time"

	"gorm.io/gorm"
)

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
