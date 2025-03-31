package mysql

import (
	"github.com/NJUPTzza/zmall/app/payment/biz/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, payment *model.Payment) error {
	return db.Create(payment).Error
}