package order

import (
	"context"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order/orderservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() orderservice.Client
	Service() string
	CreateOrder(ctx context.Context, Req *order.CreateOrderRequest, callOptions ...callopt.Option) (r *order.CreateOrderResponse, err error)
	GetOrder(ctx context.Context, Req *order.GetOrderRequest, callOptions ...callopt.Option) (r *order.GetOrderResponse, err error)
	UpdateOrderStatus(ctx context.Context, Req *order.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order.UpdateOrderStatusResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := orderservice.NewClient(dstService, opts...)
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
	kitexClient orderservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() orderservice.Client {
	return c.kitexClient
}

func (c *clientImpl) CreateOrder(ctx context.Context, Req *order.CreateOrderRequest, callOptions ...callopt.Option) (r *order.CreateOrderResponse, err error) {
	return c.kitexClient.CreateOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) GetOrder(ctx context.Context, Req *order.GetOrderRequest, callOptions ...callopt.Option) (r *order.GetOrderResponse, err error) {
	return c.kitexClient.GetOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateOrderStatus(ctx context.Context, Req *order.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order.UpdateOrderStatusResponse, err error) {
	return c.kitexClient.UpdateOrderStatus(ctx, Req, callOptions...)
}
