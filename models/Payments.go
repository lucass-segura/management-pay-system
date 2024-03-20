package models

import (
	"gorm.io/gorm"
)

type Payment struct {
	gorm.Model

	ID          uint      `gorm:"primaryKey;column:pay_id"`
	PayDate     string    `gorm:"size:50;column:pay_date;not null"`
	Client      string    `gorm:"size:100;column:client;not null"`
	Contact     string    `gorm:"size:100;column:contact;not null"`
	Product     string    `gorm:"size:150;column:product;not null"`
	Price       float64   `gorm:"column:price;not null"`
	PayMethodID PayMethod `gorm:"foreignKey:pay_method_id"`
	Description string    `gorm:"type:text;column:description"`
}
