package service

import (
	"context"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
)

type GetOrderService struct {
	ctx context.Context
} // NewGetOrderService new GetOrderService
func NewGetOrderService(ctx context.Context) *GetOrderService {
	return &GetOrderService{ctx: ctx}
}

// Run create note info
func (s *GetOrderService) Run(req *order.GetOrderRequest) (resp *order.GetOrderResponse, err error) {
	// Finish your business logic.

	return
}
