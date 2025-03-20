package dal

import (
	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
