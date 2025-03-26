package main

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	"github.com/NJUPTzza/zmall/app/cart/biz/service"
)

// CartServiceImpl implements the last service interface defined in the IDL.
type CartServiceImpl struct{}

// AddToCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) AddToCart(ctx context.Context, req *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	resp, err = service.NewAddToCartService(ctx).Run(req)

	return resp, err
}

// GetCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) GetCart(ctx context.Context, req *cart.GetCartRequest) (resp *cart.GetCartResponse, err error) {
	resp, err = service.NewGetCartService(ctx).Run(req)

	return resp, err
}

// RemoveFromCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) RemoveFromCart(ctx context.Context, req *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {
	resp, err = service.NewRemoveFromCartService(ctx).Run(req)

	return resp, err
}

// UpdateItemQuantity implements the CartServiceImpl interface.
func (s *CartServiceImpl) UpdateItemQuantity(ctx context.Context, req *cart.UpdateQuantityRequest) (resp *cart.UpdateQuantityResponse, err error) {
	resp, err = service.NewUpdateItemQuantityService(ctx).Run(req)

	return resp, err
}

// ClearCart implements the CartServiceImpl interface.
func (s *CartServiceImpl) ClearCart(ctx context.Context, req *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {
	resp, err = service.NewClearCartService(ctx).Run(req)

	return resp, err
}
