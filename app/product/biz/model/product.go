package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name string `gorm:"uniqueIndex;type:varchar(100);not null"`
	Price float32 `gorm:"type:decimal(10,2);not null"`
	Stock int32 `gorm:"type:int;not null"`
}

func (Product) TableName() string {
	return "product"
}

func GetProductsByPage(db *gorm.DB, page, pageSize int) ([]Product, error) {
	var products []Product
	// 计算偏移量
	offset := (page - 1) * pageSize
	// 查询数据
	err := db.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(db *gorm.DB, id int64) (*Product, error) {
	var product Product
	err := db.Where("id = ?", id).First(&product).Error
	return &product, err
}