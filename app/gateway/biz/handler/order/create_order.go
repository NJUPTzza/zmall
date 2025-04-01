package handler

import (
	"context"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

type CreateOrderReq struct {
	ProductId int64 `json:"product_id"`
	Quantity  int32 `json:"quantity"`
}

func CreateOrder(orderClient orderservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// 获取 JWT 中信息
		claims := jwt.ExtractClaims(ctx, c)
		userId, ok := claims["user_id"].(int64)
		if !ok {
			c.JSON(400, map[string]string{"error": "无法获取用户ID"})
			return
		}

		var req CreateOrderReq
		if err := c.Bind(&req); err != nil {
			c.JSON(400, map[string]string{"error": "请求数据无效"})
			return
		}

		resp, err := orderClient.CreateOrder(ctx, &orderservice.CreateOrderRequest{
			UserId: userId,
			ProductId: req.ProductId,
			Quantity: req.Quantity,
		})
		if err != nil {
			c.JSON(500, map[string]string{"error": "创建订单失败"})
			return
		}

		c.JSON(200, resp)
	}
}