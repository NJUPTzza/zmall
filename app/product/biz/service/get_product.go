package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/product/biz/dal/mysql"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"gorm.io/gorm"
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
	if req.Id == 0 {
		return nil, errors.New("id is empty")
	}
	productDetail, err := mysql.GetProductById(mysql.DB, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &product.GetProductResponse{
				CommonResponse: &product.CommonResponse{
					Code:    404,
					Message: "product not found",
				},
			}, nil
		}
		return nil, err
	}

	return &product.GetProductResponse{
		Product: &product.Product{
			Id: int64(productDetail.ID),
			Name: productDetail.Name,
			Price: float32(productDetail.Price),
			Stock: int32(productDetail.Stock),
		},
		CommonResponse: &product.CommonResponse{
			Code: 200,
			Message: "get product detail success",
		},
	}, nil
}
