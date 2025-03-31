package mysql

import (
	"github.com/NJUPTzza/zmall/app/payment/biz/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, payment *model.Payment) error {
	return db.Create(payment).Error
}

func GetPaymentById(db *gorm.DB, id int64) (*model.Payment, error) {
	var payment model.Payment
	err := db.Where("id = ?", id).First(&payment).Error
	return &payment, err
}

func UpdatePayment(db *gorm.DB, payment *model.Payment) error {
	return db.Save(payment).Error	
}