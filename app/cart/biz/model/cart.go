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
