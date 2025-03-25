package handler

import (
	"context"
	"strconv"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetProduct(productClient productservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		idStr := c.Query("id")
		if idStr == "" {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}
		// 将 id 转换为 int64
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "invalid id"})
			return
		}
		// 创建 GetProductRequest 消息体
		req := product.GetProductRequest {
			Id: int64(id),
		}
		resp, err := productClient.GetProduct(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		c.JSON(consts.StatusOK, resp)
	}
}