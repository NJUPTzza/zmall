package service

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

type AddToCartService struct {
	ctx context.Context
} // NewAddToCartService new AddToCartService
func NewAddToCartService(ctx context.Context) *AddToCartService {
	return &AddToCartService{ctx: ctx}
}

// Run create note info
func (s *AddToCartService) Run(req *cart.AddToCartRequest) (resp *cart.AddToCartResponse, err error) {
	// Finish your business logic.

	return
}
