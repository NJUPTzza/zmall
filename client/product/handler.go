package main

import (
	"context"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
	"github.com/NJUPTzza/zmall/biz/service"
)

// ProductCatalogServiceImpl implements the last service interface defined in the IDL.
type ProductCatalogServiceImpl struct{}

// AddProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) AddProduct(ctx context.Context, req *product.AddProductReq) (resp *product.AddProductResp, err error) {
	resp, err = service.NewAddProductService(ctx).Run(req)

	return resp, err
}

// UpdateProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) UpdateProduct(ctx context.Context, req *product.UpdateProductReq) (resp *product.UpdateProductResp, err error) {
	resp, err = service.NewUpdateProductService(ctx).Run(req)

	return resp, err
}

// DeleteProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DeleteProduct(ctx context.Context, req *product.DeleteProductReq) (resp *product.DeleteProductResp, err error) {
	resp, err = service.NewDeleteProductService(ctx).Run(req)

	return resp, err
}

// OnlineProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) OnlineProduct(ctx context.Context, req *product.OnlineProductReq) (resp *product.OnlineProductResp, err error) {
	resp, err = service.NewOnlineProductService(ctx).Run(req)

	return resp, err
}

// OfflineProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) OfflineProduct(ctx context.Context, req *product.OfflineProductReq) (resp *product.OfflineProductResp, err error) {
	resp, err = service.NewOfflineProductService(ctx).Run(req)

	return resp, err
}

// ListProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) ListProducts(ctx context.Context, req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	resp, err = service.NewListProductsService(ctx).Run(req)

	return resp, err
}

// GetProduct implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetProduct(ctx context.Context, req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	resp, err = service.NewGetProductService(ctx).Run(req)

	return resp, err
}

// BatchGetProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) BatchGetProducts(ctx context.Context, req *product.BatchGetProductsReq) (resp *product.BatchGetProductsResp, err error) {
	resp, err = service.NewBatchGetProductsService(ctx).Run(req)

	return resp, err
}

// SearchProducts implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) SearchProducts(ctx context.Context, req *product.SearchProductsReq) (resp *product.SearchProductsResp, err error) {
	resp, err = service.NewSearchProductsService(ctx).Run(req)

	return resp, err
}

// GetCategories implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetCategories(ctx context.Context, req *product.GetCategoriesReq) (resp *product.GetCategoriesResp, err error) {
	resp, err = service.NewGetCategoriesService(ctx).Run(req)

	return resp, err
}

// GetCategory implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) GetCategory(ctx context.Context, req *product.GetCategoryReq) (resp *product.GetCategoryResp, err error) {
	resp, err = service.NewGetCategoryService(ctx).Run(req)

	return resp, err
}

// DecrStock implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) DecrStock(ctx context.Context, req *product.DecrStockReq) (resp *product.DecrStockResp, err error) {
	resp, err = service.NewDecrStockService(ctx).Run(req)

	return resp, err
}

// IncrStock implements the ProductCatalogServiceImpl interface.
func (s *ProductCatalogServiceImpl) IncrStock(ctx context.Context, req *product.IncrStockReq) (resp *product.IncrStockResp, err error) {
	resp, err = service.NewIncrStockService(ctx).Run(req)

	return resp, err
}
