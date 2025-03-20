package dal

import (
	"github.com/NJUPTzza/zmall/app/cart/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/cart/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
