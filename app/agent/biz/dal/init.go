package dal

import (
	"github.com/NJUPTzza/zmall/app/agent/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/agent/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
