package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Price float32 `gorm:"type:decimal(10,2);not null"`
	Stock int32 `gorm:"type:int;not null;default:0;check:stock>=0"`
	Brand    string  `gorm:"type:varchar(50);default:''"`            // 品牌
	Category string  `gorm:"type:varchar(50);default:''"`            // 分类
}

func (Product) TableName() string {
	return "product"
}

func (p *Product) ToESDoc() *ProductESDoc {
	return &ProductESDoc{
		ID:       int64(p.ID),
		Name:     p.Name,
		Price:    p.Price,
		Stock:    p.Stock,
		Brand:    p.Brand,
		Category: p.Category,
	}
}