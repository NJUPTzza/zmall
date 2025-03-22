package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"uniqueIndex"`
	PasswordHashed string `gorm:"type:varchar(255) not null"` 
	Email string `gorm:"uniqueIndex"`
}

func (User) TableName() string {
	return "user"
}

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}