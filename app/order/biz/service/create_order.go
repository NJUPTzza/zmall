package service

import (
	"context"
	"errors"

	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
)

type CreateOrderService struct {
	ctx context.Context
	productClient productservice.Client
} 

// NewCreateOrderService new CreateOrderService
func NewCreateOrderService(ctx context.Context, productClient productservice.Client) *CreateOrderService {
	return &CreateOrderService{
		ctx: ctx,
		productClient: productClient,
	}
}

// Run create note info
func (s *CreateOrderService) Run(req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	// 1. 获取商品信息
	pResp, err := s.productClient.GetProduct(s.ctx, &product.GetProductRequest{
		Id: req.ProductId,
	})
	if err != nil {
		return nil, err
	}
	if pResp.CommonResponse.Code != 200 {
		return nil, errors.New("获取商品信息失败")
	}

	// 2. 检查库存是否足够
	if pResp.Product.Stock < req.Quantity {
		return &order.CreateOrderResponse{
			CommonResponse: &order.CommonResponse{Code: 400, Message: "库存不足"},
		}, nil
	}
	return
}
