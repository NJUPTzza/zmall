package dal

import (
	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/product/biz/dal/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
