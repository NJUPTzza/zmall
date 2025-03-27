package model

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    int64 `gorm:"type:int;not null"`
	ProductId int64 `gorm:"type:int;not null"`
	Quantity  int32 `gorm:"type:int;not null;default:1"`
}

func (Cart) TableName() string {
	return "cart"
}

func Create(db *gorm.DB, cartItem *Cart) error {
	return db.Create(cartItem).Error
}

func GetCartItems(db *gorm.DB, userId int64) ([]Cart, error) {
	var cartItems []Cart
	if err := db.Where("user_id = ?", userId).Find(&cartItems).Error; err != nil {
		return nil, err
	}
	return cartItems, nil
}