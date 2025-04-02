package service

import (
	"context"
	"errors"
	"log"

	common "github.com/NJUPTzza/zmall/app/common/mq"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type UpdateOrderStatusService struct {
	ctx context.Context
} // NewUpdateOrderStatusService new UpdateOrderStatusService
func NewUpdateOrderStatusService(ctx context.Context) *UpdateOrderStatusService {
	return &UpdateOrderStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdateOrderStatusService) Run(req *order.UpdateOrderStatusRequest) (resp *order.UpdateOrderStatusResponse, err error) {
	ch := common.RabbitMQChannel
	if ch == nil {
		return nil, errors.New("MQ 未初始化")
	}
	queueName := "payment_event_queue"
	msgs, err := ch.Consume(
		queueName, // 队列名称
		"",        // 消费者标签
		true,      // 自动确认
		false,     // 非排他
		false,     // 非阻塞
		false,     // 其他参数
		nil,
	)
	if err != nil {
		log.Printf("无法消费队列 %s: %v", queueName, err)
		return nil, err
	}
	go func() {
		for msg := range msgs {
			s.processPaymentMessage(msg)
		}
	}()
	log.Println("MQ 消费者已启动，等待消息中...")

	return &order.UpdateOrderStatusResponse{
		CommonResponse: &order.CommonResponse{
			Code:    200,
			Message: "订单状态监听已启动",
		},
	}, nil
}

// 处理 MQ 消息
func (s *UpdateOrderStatusService) processPaymentMessage(msg amqp091.Delivery) {
	var event payment.paymentMQEvent
	if err := proto.Unmarshal(msg.Body, &event); err != nil {
		log.Printf("消息解析失败: %v", err)
		return
	}

	// 日志打印
	log.Printf("收到支付事件: 订单ID=%d, 支付状态=%v", event.OrderId, event.Status)

	// 订单状态映射
	var newStatus order.OrderStatus
	switch event.Status {
	case payment.PaymentStatus_PAID:
		newStatus = order.OrderStatus_SUCCESS // 订单完成
	case payment.PaymentStatus_FAILED:
		newStatus = order.OrderStatus_FAILED // 订单失败
	case payment.PaymentStatus_REFUNDED:
		newStatus = order.OrderStatus_REFUNDED // 订单已退款
	default:
		log.Printf("未知的支付状态: %v, 忽略处理", event.Status)
		return
	}

	// 更新数据库中的订单状态（这里应调用数据库层方法）
	log.Printf("订单 %d 状态更新为 %v", event.OrderId, newStatus)

	// TODO: 调用数据库方法更新订单状态
	// err := mysql.UpdateOrderStatus(event.OrderId, newStatus)
	// if err != nil {
	// 	log.Printf("更新订单 %d 状态失败: %v", event.OrderId, err)
	// }
}
