package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Price float32 `gorm:"type:decimal(10,2);not null"`
	Stock int32 `gorm:"type:int;not null;default:0;check:stock>=0"`
}

func (Product) TableName() string {
	return "product"
}
