package mysql

import (
	"github.com/NJUPTzza/zmall/app/product/biz/model"
	"gorm.io/gorm"
)

func GetProductsByPage(db *gorm.DB, page, pageSize int) ([]model.Product, error) {
	var products []model.Product
	// 计算偏移量
	offset := (page - 1) * pageSize
	// 查询数据
	err := db.Offset(offset).Limit(pageSize).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil
}

func GetProductById(db *gorm.DB, id int64) (*model.Product, error) {
	var product model.Product
	err := db.Where("id = ?", id).First(&product).Error
	return &product, err
}

func UpdateStock(db *gorm.DB, productID int64, quantity int32) error {
	// 减少库存
	return db.Model(&model.Product{}).
		Where("id = ? AND stock >= ?", productID, quantity).
		Update("stock", gorm.Expr("stock - ?", quantity)).Error
}

func GetAllProductStock(db *gorm.DB) (map[int64]int32, error) {
	var products []model.Product
	stockMap := make(map[int64]int32)
	err := db.Select("id", "stock").Find(&products).Error
	if err != nil {
		return nil, err
	}
	for _, p := range products {
		stockMap[int64(p.ID)] = p.Stock
	}
	return stockMap, nil
}