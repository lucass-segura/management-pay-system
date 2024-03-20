package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;column:user_id"`
	FirstName string `gorm:"size:100;column:first_name;not null"`
	LastName  string `gorm:"size:100;column:last_name;not null"`
	RolID     uint   `gorm:"column:rol_id;not null"`
	Rol       Role   `gorm:"foreignKey:RolID"`
}
