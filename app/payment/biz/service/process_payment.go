package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/NJUPTzza/zmall/app/payment/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/payment/biz/model"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order/orderservice"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

type ProcessPaymentService struct {
	ctx context.Context
	orderClient orderservice.Client
} 

// NewProcessPaymentService new ProcessPaymentService
func NewProcessPaymentService(ctx context.Context, orderClient orderservice.Client) *ProcessPaymentService {
	return &ProcessPaymentService{
		ctx: ctx,
		orderClient: orderClient,
	}
}

// Run create note info
func (s *ProcessPaymentService) Run(req *payment.ProcessPaymentRequest) (resp *payment.ProcessPaymentResponse, err error) {
	// 1. 检查订单是否存在
	pResp, err := s.orderClient.GetOrder(s.ctx, &order.GetOrderRequest{
		Id: req.GetOrderId(),
	})
	if err != nil {
		return nil, errors.New("订单服务调用失败: " + err.Error())
	}
	if pResp.Order == nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 生成支付URL
	payURL, err := generatePaymentURL(req.OrderId, req.Amount, req.PayChannel)
	if err != nil {
		return nil, errors.New("生成支付URL失败: " + err.Error())
	}

	// 3. 将支付信息存入数据库中
	newPayment := &model.Payment {
		OrderId: req.OrderId,
		Amount: req.Amount,
		PayChannel: req.PayChannel,
		Status: model.PaymentStatusPending,
		PayURL: payURL,
	}
	err = mysql.Create(mysql.DB, newPayment)
	if err != nil {
		return nil, errors.New("支付信息存入数据库失败: " + err.Error())
	}

	// 4. 触发 `PAY` 事件，更新状态（`PENDING` -> `PENDING`，但触发了状态机）
	updateService := NewUpdatePaymentStatusService(s.ctx)
	_, err = updateService.Run(&payment.UpdatePaymentStatusRequest{
		PaymentId: int64(newPayment.ID),
		Event:     payment.PaymentEvent_PAY,
	})
	if err != nil {
		return nil, errors.New("触发支付事件失败: " + err.Error())
	}

	// 5. 返回支付信息
	return &payment.ProcessPaymentResponse {
		CommonResponse: &payment.CommonResponse{
			Code: 200,
			Message:  "success",
		},
		Payment: &payment.Payment{
			Id: int64(newPayment.ID),
			OrderId: newPayment.OrderId,
			Amount: newPayment.Amount,
		},
		PayUrl: newPayment.PayURL,
	}, nil
}

// 模拟支付 URL 生成函数
func generatePaymentURL(orderId int64, amount float32, payChannel string) (string, error) {
	// 根据支付渠道生成支付 URL
	switch payChannel {
	case "wechat":
		// 微信支付模拟 URL
		return fmt.Sprintf("https://pay.wechat.com/%d?amount=%.2f", orderId, amount), nil
	case "alipay":
		// 支付宝支付模拟 URL
		return fmt.Sprintf("https://pay.alipay.com/%d?amount=%.2f", orderId, amount), nil
	default:
		// 如果支付渠道不支持，返回错误
		return "", fmt.Errorf("unsupported payment channel: %s", payChannel)
	}
}
