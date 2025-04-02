package handler

import (
	"context"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ProcessPayment(paymentClient paymentservice.Client) app.HandlerFunc {
	return func (ctx context.Context, c *app.RequestContext) {
		var req payment.ProcessPaymentRequest

		if err := c.Bind(&req); err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		resp, err := paymentClient.ProcessPayment(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(consts.StatusOK, resp)
	}
}