package product

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() productservice.Client
	Service() string
	ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error)
	GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error)
	UpdateStock(ctx context.Context, Req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := productservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient productservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() productservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ListProducts(ctx context.Context, Req *product.ListProductsRequest, callOptions ...callopt.Option) (r *product.ListProductsResponse, err error) {
	return c.kitexClient.ListProducts(ctx, Req, callOptions...)
}

func (c *clientImpl) GetProduct(ctx context.Context, Req *product.GetProductRequest, callOptions ...callopt.Option) (r *product.GetProductResponse, err error) {
	return c.kitexClient.GetProduct(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateStock(ctx context.Context, Req *product.UpdateStockRequest, callOptions ...callopt.Option) (r *product.UpdateStockResponse, err error) {
	return c.kitexClient.UpdateStock(ctx, Req, callOptions...)
}
