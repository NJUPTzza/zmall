package payment

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() paymentservice.Client
	Service() string
	ProcessPayment(ctx context.Context, Req *payment.ProcessPaymentRequest, callOptions ...callopt.Option) (r *payment.ProcessPaymentResponse, err error)
	GetPaymentStatus(ctx context.Context, Req *payment.GetPaymentStatusRequest, callOptions ...callopt.Option) (r *payment.GetPaymentStatusResponse, err error)
	UpdatePaymentStatus(ctx context.Context, Req *payment.UpdatePaymentStatusRequest, callOptions ...callopt.Option) (r *payment.UpdatePaymentStatusResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := paymentservice.NewClient(dstService, opts...)
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
	kitexClient paymentservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() paymentservice.Client {
	return c.kitexClient
}

func (c *clientImpl) ProcessPayment(ctx context.Context, Req *payment.ProcessPaymentRequest, callOptions ...callopt.Option) (r *payment.ProcessPaymentResponse, err error) {
	return c.kitexClient.ProcessPayment(ctx, Req, callOptions...)
}

func (c *clientImpl) GetPaymentStatus(ctx context.Context, Req *payment.GetPaymentStatusRequest, callOptions ...callopt.Option) (r *payment.GetPaymentStatusResponse, err error) {
	return c.kitexClient.GetPaymentStatus(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdatePaymentStatus(ctx context.Context, Req *payment.UpdatePaymentStatusRequest, callOptions ...callopt.Option) (r *payment.UpdatePaymentStatusResponse, err error) {
	return c.kitexClient.UpdatePaymentStatus(ctx, Req, callOptions...)
}
