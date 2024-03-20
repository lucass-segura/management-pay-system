package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model

	ID  uint   `gorm:"primaryKey;autoIncrement;column:rol_id"`
	Rol string `gorm:"size:100;column:rol;not null"`
}
