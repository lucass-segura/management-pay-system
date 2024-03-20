package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID        uint   `gorm:"primaryKey;autoIncrement;column:user_id"`
	FirstName string `gorm:"size100;column:first_name;not null"`
	LastName  string `gorm:"size100;column:last_name;not null"`
	RolID     Role   `gorm:"foreignKey:rol_id;column:rol_id;not null"`
}
