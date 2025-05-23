// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productservice

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error)
	GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error)
	UpdateStock(ctx context.Context, Req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error)
	SearchProducts(ctx context.Context, Req *product.SearchProductsRequest, callOptions ...callopt.Option) (r *product.SearchProductsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kProductServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kProductServiceClient struct {
	*kClient
}

func (p *kProductServiceClient) ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ListProducts(ctx, Req)
}

func (p *kProductServiceClient) GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetProduct(ctx, Req)
}

func (p *kProductServiceClient) UpdateStock(ctx context.Context, Req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateStock(ctx, Req)
}

func (p *kProductServiceClient) SearchProducts(ctx context.Context, Req *product.SearchProductsRequest, callOptions ...callopt.Option) (r *product.SearchProductsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SearchProducts(ctx, Req)
}
