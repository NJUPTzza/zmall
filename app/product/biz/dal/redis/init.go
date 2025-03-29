package redis

import (
	"context"

	"github.com/NJUPTzza/zmall/app/product/conf"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var (
	RedisClient *redis.Client
)	

func Init() {
	cfg := conf.GetConf().Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     conf.GetConf().Redis.Address,
		Username: conf.GetConf().Redis.Username,
		Password: conf.GetConf().Redis.Password,
		DB:       conf.GetConf().Redis.DB,
	})

	// 测试连接
	ctx := context.Background()
	if err := RedisClient.Ping(ctx).Err(); err != nil {
		klog.Fatalf("[Redis] 连接失败: %v", err) // 连接失败时直接终止程序
	}

	klog.Infof("[Redis] 连接成功，地址: %s", cfg.Address) // 连接成功时打印日志

	CacheProductStock(ctx)
}
