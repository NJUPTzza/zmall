package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductRequest) (resp *product.GetProductResponse, err error) {
	// Finish your business logic.

	return
}
