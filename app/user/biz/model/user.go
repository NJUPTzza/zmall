package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex;type:varchar(50) not null"`
	PasswordHashed string `gorm:"type:varchar(255) not null"` 
	Email string `gorm:"uniqueIndex;type:varchar(100) not null"`
	Phone string `gorm:"uniqueIndex;type:varchar(20) not null"`
}

func (User) TableName() string {
	return "user"
}
