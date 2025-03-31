package main

import (
	"context"

	"github.com/NJUPTzza/zmall/app/order/biz/service"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
)

// OrderServiceImpl implements the last service interface defined in the IDL.
type OrderServiceImpl struct {
	productClient productservice.Client
}

// CreateOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (resp *order.CreateOrderResponse, err error) {
	resp, err = service.NewCreateOrderService(ctx, s.productClient).Run(req)

	return resp, err
}

// GetOrder implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) GetOrder(ctx context.Context, req *order.GetOrderRequest) (resp *order.GetOrderResponse, err error) {
	resp, err = service.NewGetOrderService(ctx).Run(req)

	return resp, err
}

// UpdateOrderStatus implements the OrderServiceImpl interface.
func (s *OrderServiceImpl) UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest) (resp *order.UpdateOrderStatusResponse, err error) {
	resp, err = service.NewUpdateOrderStatusService(ctx).Run(req)

	return resp, err
}
