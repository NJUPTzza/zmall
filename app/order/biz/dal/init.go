package dal

import (
	"github.com/NJUPTzza/zmall/app/order/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
