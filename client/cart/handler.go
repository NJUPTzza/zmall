package main

import (
	"context"
	cart "github.com/NJUPTzza/zmall/client/cart/kitex_gen/cart"
	"github.com/NJUPTzza/zmall/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddItem(ctx context.Context, req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	resp, err = service.NewAddItemService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartReq) (resp *cart.GetCartResp, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// EmptyCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) EmptyCart(ctx context.Context, req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	resp, err = service.NewEmptyCartService(ctx).Run(req)

	return resp, err
}

// UpdateCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateCart(ctx context.Context, req *cart.UpdateCartReq) (resp *cart.UpdateCartResp, err error) {
	resp, err = service.NewUpdateCartService(ctx).Run(req)

	return resp, err
}

// DeleteItem implements the CartServiceImpl interface.
func (s *CartServiceImpl) DeleteItem(ctx context.Context, req *cart.DeleteItemReq) (resp *cart.DeleteItemResp, err error) {
	resp, err = service.NewDeleteItemService(ctx).Run(req)

	return resp, err
}
