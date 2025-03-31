package dal

import (
	"github.com/NJUPTzza/zmall/app/notification/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/notification/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
