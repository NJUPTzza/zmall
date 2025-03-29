package mysql

import (
	"github.com/NJUPTzza/zmall/app/cart/biz/model"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, cartItem *model.Cart) error {
	return db.Create(cartItem).Error
}

func GetCartItems(db *gorm.DB, userId int64) ([]model.Cart, error) {
	var cartItems []model.Cart
	if err := db.Where("user_id = ?", userId).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}