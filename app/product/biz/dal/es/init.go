package es

import (
	"log"

	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
	"github.com/olivere/elastic/v7"
)

var Client *elastic.Client

func Init() {
	client, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
		elastic.SetSniff(false), // 不用 sniff，否则 Docker 会报错
	)
	if err != nil {
		log.Fatalf("Elasticsearch 初始化失败: %v", err)
	}
	Client = client

	SyncAllProductsToES(mysql.DB)
}