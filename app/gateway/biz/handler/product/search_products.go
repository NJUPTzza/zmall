package handler

import (
	"context"
	"strconv"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

type SearchRequestBody struct {
    Keyword   string  `json:"keyword"`
    Brand     string  `json:"brand"`
    Category  string  `json:"category"`
    MinPrice  float32 `json:"min_price"`
    MaxPrice  float32 `json:"max_price"`
}

func SearchProducts(productClient productservice.Client) app.HandlerFunc{
	return func(ctx context.Context, c *app.RequestContext) {
		// 从 query 获取分页参数
		page, _ := strconv.Atoi(c.DefaultQuery("page", "1")) 
		pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

		// 从 JSON body 解析搜索参数
		var body SearchRequestBody
		if err := c.Bind(&body); err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		 // 构造 gRPC 请求
		 req := productClient
	}
}