package mysql

import (
	"github.com/NJUPTzza/zmall/app/user/biz/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}

func GetByUsername(db *gorm.DB, username string) (*model.User, error) {
	var user model.User
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetUserById(db *gorm.DB, id int64) (*model.User, error) {
	var user model.User
	err := db.Where("id = ?", id).First(&user).Error
	return &user, err
}

