package dal

import (
	"github.com/NJUPTzza/zmall/app/payment/biz/dal/mq"
	"github.com/NJUPTzza/zmall/app/payment/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/payment/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
	mq.Init()
}
