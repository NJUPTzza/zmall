package service

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

type ClearCartService struct {
	ctx context.Context
} // NewClearCartService new ClearCartService
func NewClearCartService(ctx context.Context) *ClearCartService {
	return &ClearCartService{ctx: ctx}
}

// Run create note info
func (s *ClearCartService) Run(req *cart.ClearCartRequest) (resp *cart.ClearCartResponse, err error) {
	// Finish your business logic.

	return
}
