package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

type GetCategoryService struct {
	ctx context.Context
} // NewGetCategoryService new GetCategoryService
func NewGetCategoryService(ctx context.Context) *GetCategoryService {
	return &GetCategoryService{ctx: ctx}
}

// Run create note info
func (s *GetCategoryService) Run(req *product.GetCategoryReq) (resp *product.GetCategoryResp, err error) {
	// Finish your business logic.

	return
}
