// Code generated by Kitex v0.9.1. DO NOT EDIT.

package orderservice

import (
	"context"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateOrder(ctx context.Context, Req *order.CreateOrderRequest, callOptions ...callopt.Option) (r *order.CreateOrderResponse, err error)
	GetOrder(ctx context.Context, Req *order.GetOrderRequest, callOptions ...callopt.Option) (r *order.GetOrderResponse, err error)
	UpdateOrderStatus(ctx context.Context, Req *order.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order.UpdateOrderStatusResponse, err error)
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
	return &kOrderServiceClient{
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

type kOrderServiceClient struct {
	*kClient
}

func (p *kOrderServiceClient) CreateOrder(ctx context.Context, Req *order.CreateOrderRequest, callOptions ...callopt.Option) (r *order.CreateOrderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateOrder(ctx, Req)
}

func (p *kOrderServiceClient) GetOrder(ctx context.Context, Req *order.GetOrderRequest, callOptions ...callopt.Option) (r *order.GetOrderResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetOrder(ctx, Req)
}

func (p *kOrderServiceClient) UpdateOrderStatus(ctx context.Context, Req *order.UpdateOrderStatusRequest, callOptions ...callopt.Option) (r *order.UpdateOrderStatusResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateOrderStatus(ctx, Req)
}
