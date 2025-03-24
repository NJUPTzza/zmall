package service

import (
	"context"

	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/product/biz/model"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/klog"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsRequest) (resp *product.ListProductsResponse, err error) {
	// Finish your business logic.
	page := int(req.Page)
	pageSize := int(req.PageSize)
	
	products, err := model.GetProductsByPage(mysql.DB, page, pageSize)
	if err != nil {
		klog.Error(err)
		return nil, err
	}
	resp = &product.ListProductsResponse{
		Products: make([]*product.Product, 0, len(products)),
		CommonResponse: &product.CommonResponse{
			Code: 200,
			Message: "get product list success",
		},
	}
	for _, p := range products {
		resp.Products = append(resp.Products, &product.Product{
			Id: int64(p.ID),
			Name: p.Name,
			Price: float32(p.Price),
			Stock: int32(p.Stock),
		})
	}
	return resp, nil
}
