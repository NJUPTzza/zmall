package cart

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart/cartservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() cartservice.Client
	Service() string
	AddToCart(ctx context.Context, Req *cart.AddToCartRequest, callOptions ...callopt.Option) (r *cart.AddToCartResponse, err error)
	GetCart(ctx context.Context, Req *cart.GetCartRequest, callOptions ...callopt.Option) (r *cart.GetCartResponse, err error)
	RemoveFromCart(ctx context.Context, Req *cart.RemoveFromCartRequest, callOptions ...callopt.Option) (r *cart.RemoveFromCartResponse, err error)
	UpdateItemQuantity(ctx context.Context, Req *cart.UpdateQuantityRequest, callOptions ...callopt.Option) (r *cart.UpdateQuantityResponse, err error)
	ClearCart(ctx context.Context, Req *cart.ClearCartRequest, callOptions ...callopt.Option) (r *cart.ClearCartResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := cartservice.NewClient(dstService, opts...)
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
	kitexClient cartservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() cartservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddToCart(ctx context.Context, Req *cart.AddToCartRequest, callOptions ...callopt.Option) (r *cart.AddToCartResponse, err error) {
	return c.kitexClient.AddToCart(ctx, Req, callOptions...)
}

func (c *clientImpl) GetCart(ctx context.Context, Req *cart.GetCartRequest, callOptions ...callopt.Option) (r *cart.GetCartResponse, err error) {
	return c.kitexClient.GetCart(ctx, Req, callOptions...)
}

func (c *clientImpl) RemoveFromCart(ctx context.Context, Req *cart.RemoveFromCartRequest, callOptions ...callopt.Option) (r *cart.RemoveFromCartResponse, err error) {
	return c.kitexClient.RemoveFromCart(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateItemQuantity(ctx context.Context, Req *cart.UpdateQuantityRequest, callOptions ...callopt.Option) (r *cart.UpdateQuantityResponse, err error) {
	return c.kitexClient.UpdateItemQuantity(ctx, Req, callOptions...)
}

func (c *clientImpl) ClearCart(ctx context.Context, Req *cart.ClearCartRequest, callOptions ...callopt.Option) (r *cart.ClearCartResponse, err error) {
	return c.kitexClient.ClearCart(ctx, Req, callOptions...)
}
