package mq

import (
	common "github.com/NJUPTzza/zmall/app/common/mq"
)

func Init() {
	common.InitRabbitMQ()
}