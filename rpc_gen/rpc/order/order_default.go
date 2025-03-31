package order

import (
	"context"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func CreateOrder(ctx context.Context, req *order.CreateOrderRequest, callOptions ...callopt.Option) (resp *order.CreateOrderResponse, err error) {
	resp, err = defaultClient.CreateOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOrder(ctx context.Context, req *order.GetOrderRequest, callOptions ...callopt.Option) (resp *order.GetOrderResponse, err error) {
	resp, err = defaultClient.GetOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateOrderStatus(ctx context.Context, req *order.UpdateOrderStatusRequest, callOptions ...callopt.Option) (resp *order.UpdateOrderStatusResponse, err error) {
	resp, err = defaultClient.UpdateOrderStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateOrderStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
