package main

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/NJUPTzza/zmall/app/product/biz/service"
)

// ProductServiceImpl implements the last service interface defined in the IDL.
type ProductServiceImpl struct{}

// ListProducts implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsRequest) (resp *product.ListProductsResponse, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) GetProduct(ctx context.Context, req *product.GetProductRequest) (resp *product.GetProductResponse, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// UpdateStock implements the ProductServiceImpl interface.
func (s *ProductServiceImpl) UpdateStock(ctx context.Context, req *product.UpdateStockRequest) (resp *product.UpdateStockResponse, err error) {
	resp, err = service.NewUpdateStockService(ctx).Run(req)

	return resp, err
}
