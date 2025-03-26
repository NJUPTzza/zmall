package service

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

type UpdateItemQuantityService struct {
	ctx context.Context
} // NewUpdateItemQuantityService new UpdateItemQuantityService
func NewUpdateItemQuantityService(ctx context.Context) *UpdateItemQuantityService {
	return &UpdateItemQuantityService{ctx: ctx}
}

// Run create note info
func (s *UpdateItemQuantityService) Run(req *cart.UpdateQuantityRequest) (resp *cart.UpdateQuantityResponse, err error) {
	// Finish your business logic.

	return
}
