package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

type OnlineProductService struct {
	ctx context.Context
} // NewOnlineProductService new OnlineProductService
func NewOnlineProductService(ctx context.Context) *OnlineProductService {
	return &OnlineProductService{ctx: ctx}
}

// Run create note info
func (s *OnlineProductService) Run(req *product.OnlineProductReq) (resp *product.OnlineProductResp, err error) {
	// Finish your business logic.

	return
}
