package handler

import (
	"context"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type UpdatePaymentStatusReq struct {
	PaymentID int64               `json:"payment_id"`
	Event     payment.PaymentEvent `json:"event"`
}

/*
const (
	PaymentEvent_PAY            PaymentEvent = 0 // 发起支付
	PaymentEvent_PAY_SUCCESS    PaymentEvent = 1 // 支付成功
	PaymentEvent_PAY_FAIL       PaymentEvent = 2 // 支付失败
	PaymentEvent_REQUEST_REFUND PaymentEvent = 3 // 申请退款
	PaymentEvent_REFUND_SUCCESS PaymentEvent = 4 // 退款成功
)
*/

func UpdatePaymentStatus(paymentClient paymentservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req UpdatePaymentStatusReq

		// 解析 JSON 请求体
		if err := c.Bind(&req); err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "请求参数错误"})
			return
		}

		// 调用 RPC 服务
		resp, err := paymentClient.UpdatePaymentStatus(ctx, &payment.UpdatePaymentStatusRequest {
			PaymentId: req.PaymentID,
			Event:     req.Event,
		})
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		// 返回响应
		c.JSON(consts.StatusOK, resp)
	}
}