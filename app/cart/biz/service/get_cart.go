package service

import (
	"context"

	"github.com/NJUPTzza/zmall/app/cart/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/cart/biz/model"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/klog"
)

type GetCartService struct {
	ctx context.Context
} 

// NewGetCartService new GetCartService
func NewGetCartService(ctx context.Context) *GetCartService {
	return &GetCartService{ctx: ctx}
}

// Run create note info
func (s *GetCartService) Run(req *cart.GetCartRequest) (resp *cart.GetCartResponse, err error) {
	cartItems, err := model.GetCartItems(mysql.DB, req.UserId)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	resp = &cart.GetCartResponse{
		CommonResponse: &cart.CommonResponse{
			Code: 200,
			Message: "获取购物车成功",
		},
		Items: make([]*cart.CartItem, len(cartItems)),
	}
	for i, p := range cartItems {
		resp.Items[i] = &cart.CartItem{
			Id:        int64(p.ID),
			ProductId: p.ProductId,
			Quantity:  p.Quantity,
		}
	}
	return resp, nil
}
