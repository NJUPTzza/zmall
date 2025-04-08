package product

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ListProducts(ctx context.Context, req *product.ListProductsRequest, callOptions ...callopt.Option) (resp *product.ListProductsResponse, err error) {
	resp, err = defaultClient.ListProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetProduct(ctx context.Context, req *product.GetProductRequest, callOptions ...callopt.Option) (resp *product.GetProductResponse, err error) {
	resp, err = defaultClient.GetProduct(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetProduct call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateStock(ctx context.Context, req *product.UpdateStockRequest, callOptions ...callopt.Option) (resp *product.UpdateStockResponse, err error) {
	resp, err = defaultClient.UpdateStock(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateStock call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func SearchProducts(ctx context.Context, req *product.SearchProductsRequest, callOptions ...callopt.Option) (resp *product.SearchProductsResponse, err error) {
	resp, err = defaultClient.SearchProducts(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SearchProducts call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
