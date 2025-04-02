package mq

import (
	common "github.com/NJUPTzzz/zmall/app/common/mq"
)

func Init() {
	common.InitRabbitMQ()
}