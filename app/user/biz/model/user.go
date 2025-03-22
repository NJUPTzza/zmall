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

func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}

func GetByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}