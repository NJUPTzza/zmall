package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/order/biz/dal/mysql"
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
	// 1. 根据订单 ID 查询订单信息
	orderInfo, err := mysql.GetOrderById(mysql.DB, req.Id)
	if err != nil {
		return nil, errors.New("订单不存在")
	}

	// 2. 返回订单信息
	resp = &order.GetOrderResponse{
		CommonResponse: &order.CommonResponse{
			Code: 200,
			Message:  "success",
		},
		Order: &order.Order{
			Id: int64(orderInfo.ID),
			UserId: orderInfo.UserId,
			ProductId: orderInfo.ProductId,
			Quantity: orderInfo.Quantity,
			TotalPrice: orderInfo.TotalPrice,
			Status: order.OrderStatus(orderInfo.Status),
		},
	}
	return resp, nil
}
