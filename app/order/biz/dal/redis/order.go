package redis

import (
	"context"
	"errors"
	"strconv"
)

const stockPrefix = "product:stock:"

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

func DecrProductStock(ctx context.Context, productID int64, quantity int32) (bool, error) {
	key := stockPrefix + strconv.FormatInt(productID, 10)
	luaScript := `
		if redis.call("GET", KEYS[1]) >= ARGV[1] then
			return redis.call("DECRBY", KEYS[1], ARGV[1])
		else
			return -1
		end
	`
	result, err := RedisClient.Eval(ctx, luaScript, []string{key}, quantity).Int()
	if err!= nil {
		return false, err
	}

	/*
		返回值为 -1 表示未找到该缓存
		返回值为 -2 表示库存不足
		返回值为 1 表示扣减成功	
	*/
	if result < 0 {
		return false, nil		
	} 
	return true, nil
}

func IncrProductStock(ctx context.Context, productID int64, quantity int32) error {
	key := stockPrefix + strconv.FormatInt(productID, 10)

	_, err := RedisClient.IncrBy(ctx, key, int64(quantity)).Result()
	if err!= nil {
		return errors.New("回滚库存失败:" + err.Error())
	}
	return nil
}