package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/NJUPTzza/zmall/app/common"
	"github.com/NJUPTzza/zmall/app/payment/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/payment/biz/model"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
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
		ch := common
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