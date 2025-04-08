package common

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

var RabbitMQChannel *amqp091.Channel

func InitRabbitMQ() (*amqp091.Channel, error) {
	// 连接到RabbitMQ服务器
	conn, err := amqp091.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		log.Fatalf("无法连接到 RabbitMQ: %v", err)
		return nil, err
	}
	// 创建一个信道
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("无法创建信道: %v", err)
		return nil, err
	}
	RabbitMQChannel = ch
	return ch, nil
}

