package models

import "gorm.io/gorm"

type PayMethod struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;column:pay_method_id"`
	PayMethod string `gorm:"size:100;column:pay_method;not null" json:"pay_method"`
}
