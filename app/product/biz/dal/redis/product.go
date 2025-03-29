package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
)

const stockPrefix = "product:stock:"

func CacheProductStock(ctx context.Context) error {
	stockMap, err := mysql.GetAllProductStock(mysql.DB)
	if err != nil {
		return errors.New("获取商品库存失败: " + err.Error())
	}
	pipe := RedisClient.Pipeline()
	for productID, stock := range stockMap {
		key := stockPrefix + strconv.FormatInt(productID, 10)
		pipe.Set(ctx, key, stock, 0)
	}
	_, err = pipe.Exec(ctx)
	if err != nil {
		return errors.New("缓存商品库存失败: " + err.Error())
	}
	fmt.Println("✅ 商品库存已成功加载到 Redis")
	return nil
}

// 查询 Redis 中的商品库存
func GetProductStock(ctx context.Context, productID int64) (int32, error) {
	key := stockPrefix + strconv.FormatInt(productID, 10)
	stockStr, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	stock, err := strconv.Atoi(stockStr)
	if err != nil {
		return 0, err
	}
	return int32(stock), nil
}