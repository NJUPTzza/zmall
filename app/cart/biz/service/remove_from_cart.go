package service

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

type RemoveFromCartService struct {
	ctx context.Context
} // NewRemoveFromCartService new RemoveFromCartService
func NewRemoveFromCartService(ctx context.Context) *RemoveFromCartService {
	return &RemoveFromCartService{ctx: ctx}
}

// Run create note info
func (s *RemoveFromCartService) Run(req *cart.RemoveFromCartRequest) (resp *cart.RemoveFromCartResponse, err error) {
	// Finish your business logic.

	return
}
