package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
)

type UpdateStockService struct {
	ctx context.Context
} // NewUpdateStockService new UpdateStockService
func NewUpdateStockService(ctx context.Context) *UpdateStockService {
	return &UpdateStockService{ctx: ctx}
}

// Run create note info
func (s *UpdateStockService) Run(req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	// Finish your business logic.
	resp = &product.UpdateStockResponse{}
	return
}
