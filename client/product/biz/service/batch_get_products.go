package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

type BatchGetProductsService struct {
	ctx context.Context
} // NewBatchGetProductsService new BatchGetProductsService
func NewBatchGetProductsService(ctx context.Context) *BatchGetProductsService {
	return &BatchGetProductsService{ctx: ctx}
}

// Run create note info
func (s *BatchGetProductsService) Run(req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	// Finish your business logic.

	return
}
