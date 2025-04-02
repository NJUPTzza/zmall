package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	common "github.com/NJUPTzza/zmall/app/common/mq"
	"github.com/NJUPTzza/zmall/app/payment/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/payment/biz/model"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/rabbitmq/amqp091-go"
	"google.golang.org/protobuf/proto"
)

type UpdatePaymentStatusService struct {
	ctx context.Context
} // NewUpdatePaymentStatusService new UpdatePaymentStatusService
func NewUpdatePaymentStatusService(ctx context.Context) *UpdatePaymentStatusService {
	return &UpdatePaymentStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdatePaymentStatusService) Run(req *payment.UpdatePaymentStatusRequest) (resp *payment.UpdatePaymentStatusResponse, err error) {
	// 1. 查询支付记录
	paymentRecord, err := mysql.GetPaymentById(mysql.DB, req.PaymentId)
	if err != nil {
		return nil, errors.New("支付记录不存在")
	}

	// 2. 事件 -> 状态 映射
	newStatus, err := mapEventToStatus(paymentRecord.Status, req.Event)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("无效状态变更: %v -> %v", paymentRecord.Status, newStatus))
	}

	// 3. 更新数据库
	paymentRecord.Status = newStatus
	err = mysql.UpdatePayment(mysql.DB, paymentRecord)
	if err != nil {
		return nil, errors.New("更新支付记录失败")
	}

	// 4. 如果支付成功，发布支付成功事件
	if req.Event == payment.PaymentEvent_PAY_SUCCESS {
		// TODO: 发布支付成功事件
		// 用 mq 通知 order 修改状态，通知 prodcut 修改库存，通知 notifacation 发送通知
		event := &payment.PaymentMQEvent {
			PaymentId: paymentRecord.ID,
			OrderId: paymentRecord.OrderId,
			UserId: paymentRecord.UserId,
			Amount: paymentRecord.Amount,
			Status: payment.PaymentStatus(paymentRecord.Status),
		}
		// 将 PaymentMQEvent 对象序列化为字节数组
		data, err := proto.Marshal(event)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("消息序列化失败: %v", err))
		}
		// 获取 RabbitMQ 通道
		ch := common.RabbitMQChannel
		if ch == nil {
			return nil, errors.New("MQ初始化失败")
		}
		// 声明队列，确保队列存在
		queueName := "payment_event_queue"
		_, err = ch.QueueDeclare(
			queueName, // 队列名称
			true,      // 是否持久化
			false,     // 是否自动删除
			false,     // 是否排他
			false,     // 是否阻塞
			nil,       // 其他参数
		)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("声明队列失败: %v", err))
		}
		// 发布消息到队列
		err = ch.Publish(
			"",           // 默认交换机
			queueName,    // 队列名称
			false,        // 是否强制
			false,        // 是否等待
			amqp091.Publishing{
				ContentType: "application/x-protobuf", // 消息格式
				Body:        data, // 消息内容
			},
		)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("发布消息失败: %v", err))
		}
		log.Println("支付成功事件已发布到 RabbitMQ")
	}

	// 5. 返回更新后的支付信息
	return &payment.UpdatePaymentStatusResponse {
		CommonResponse: &payment.CommonResponse {
			Code: 200,
			Message: "更新支付记录成功",
		},
		Payment: &payment.Payment {
			Id: int64(paymentRecord.ID),
			OrderId: paymentRecord.OrderId,
			Amount: paymentRecord.Amount,
		},
	}, nil
}

// mapEventToStatus 事件 -> 状态 映射
func mapEventToStatus(currentStatus model.PaymentStatus, event payment.PaymentEvent) (model.PaymentStatus, error) {
	switch currentStatus {
	case model.PaymentStatusPending:
		switch event {
		case payment.PaymentEvent_PAY:
			return model.PaymentStatusPending, nil // 发起支付，状态不变
		case payment.PaymentEvent_PAY_SUCCESS:
			return model.PaymentStatusPaid, nil // 支付成功
		case payment.PaymentEvent_PAY_FAIL:
			return model.PaymentStatusFailed, nil // 支付失败
		}
	case model.PaymentStatusPaid:
		switch event {
		case payment.PaymentEvent_REQUEST_REFUND, payment.PaymentEvent_REFUND_SUCCESS:
			return model.PaymentStatusRefunded, nil // 退款成功
		}
	}

	return currentStatus, errors.New("非法的支付事件")
}