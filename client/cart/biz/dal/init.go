package dal

import (
	"github.com/NJUPTzza/zmall/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
