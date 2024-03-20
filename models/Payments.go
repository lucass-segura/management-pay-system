package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	ID          uint      `gorm:"primaryKey;autoIncrement;column:pay_id"`
	PayDate     string    `gorm:"size50;column:pay_date;not null"`
	Client      string    `gorm:"size100;column:client;not null"`
	Contact     string    `gorm:"size100;column:contact;not null"`
	Product     string    `gorm:"size150;column:product;not null"`
	Price       float64   `gorm:"column:price;not null"`
	PayMethodID PayMethod `gorm:"foreignKey:pay_method_id"`
	Description string    `gorm:"type:text;column:description"`
}
