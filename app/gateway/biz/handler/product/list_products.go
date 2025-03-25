package handler

import (
	"context"
	"strconv"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func ListProducts(productClient productservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		pageStr := c.Query("page")
		pageSizeStr := c.Query("pageSize")

		if pageStr == "" || pageSizeStr == "" {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}
		
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "invalid page"})
			return
		}

		pageSize, err := strconv.Atoi(pageSizeStr)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "invalid pageSize"})
			return
		}

		req := product.ListProductsRequest {
			Page: int32(page),
			PageSize: int32(pageSize),
		}

		resp, err := productClient.ListProducts(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(consts.StatusOK, resp)
	}
}