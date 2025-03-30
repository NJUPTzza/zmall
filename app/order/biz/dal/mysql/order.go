package mysql

import (
	"github.com/NJUPTzza/zmall/app/order/biz/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, order *model.Order) error {
	return db.Create(order).Error
}

func GetOrderById(db *gorm.DB, id int64) (*model.Order, error) {
	var order model.Order
	err := db.Where("id =?", id).First(&order).Error
	return &order, err	
}