package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

type DecrStockService struct {
	ctx context.Context
} // NewDecrStockService new DecrStockService
func NewDecrStockService(ctx context.Context) *DecrStockService {
	return &DecrStockService{ctx: ctx}
}

// Run create note info
func (s *DecrStockService) Run(req *product.DecrStockReq) (resp *product.DecrStockResp, err error) {
	// Finish your business logic.

	return
}
