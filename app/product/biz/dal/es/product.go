package es

import (
	"context"
	"fmt"

	"github.com/NJUPTzza/zmall/app/product/biz/model"
	"github.com/olivere/elastic"
	"gorm.io/gorm"
)

// 同步单个商品到 ES
func SyncProductToES(product *model.Product) error {
	doc := product.ToESDoc()
	_, err := Client.Index().
		Index("products").
		Id(fmt.Sprint(product.ID)).
		BodyJson(doc).
		Refresh("true").
		Do(context.Background())

	if err != nil {
		fmt.Printf("❌ 商品 ID: %d 同步失败: %v\n", product.ID, err)
		return err
	}
	fmt.Printf("✅ 商品 ID: %d 已成功同步到 Elasticsearch\n", product.ID)
	return nil
}

// 批量同步所有商品到 ES
func SyncAllProductsToES(db *gorm.DB) error {
	var products []model.Product
	if err := db.Find(&products).Error; err != nil {
		fmt.Println("❌ 从数据库获取商品失败:", err)
		return err
	}

	bulkRequest := Client.Bulk().Index("products")
	for _, p := range products {
		doc := p.ToESDoc()
		req := elastic.NewBulkIndexRequest().Id(fmt.Sprint(doc.ID)).Doc(doc)
		bulkRequest = bulkRequest.Add(req)
	}

	_, err := bulkRequest.Do(context.Background())
	if err != nil {
		fmt.Println("❌ 批量同步商品到 Elasticsearch 失败:", err)
		return err
	}

	fmt.Printf("✅ 成功将 %d 个商品批量同步到 Elasticsearch\n", len(products))
	return nil
}
