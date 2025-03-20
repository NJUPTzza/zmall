package service

import (
	"context"
	cart "github.com/NJUPTzza/zmall/client/cart/kitex_gen/cart"
)

type DeleteItemService struct {
	ctx context.Context
} // NewDeleteItemService new DeleteItemService
func NewDeleteItemService(ctx context.Context) *DeleteItemService {
	return &DeleteItemService{ctx: ctx}
}

// Run create note info
func (s *DeleteItemService) Run(req *cart.DeleteItemReq) (resp *cart.DeleteItemResp, err error) {
	// Finish your business logic.

	return
}
