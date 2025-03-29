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

func UpdateStock(db *gorm.DB, productID int64, quantity int32, increase bool) error {
	if increase {
		// 增加库存
		return db.Model(&Product{}).
			Where("id = ?", productID).
			Update("stock", gorm.Expr("stock + ?", quantity)).Error
	} else {
		// 减少库存
		return db.Model(&Product{}).
			Where("id = ? AND stock >= ?", productID, quantity).
			Update("stock", gorm.Expr("stock - ?", quantity)).Error
	}
}