package cart

import (
	"context"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddToCart(ctx context.Context, req *cart.AddToCartRequest, callOptions ...callopt.Option) (resp *cart.AddToCartResponse, err error) {
	resp, err = defaultClient.AddToCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddToCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetCart(ctx context.Context, req *cart.GetCartRequest, callOptions ...callopt.Option) (resp *cart.GetCartResponse, err error) {
	resp, err = defaultClient.GetCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func RemoveFromCart(ctx context.Context, req *cart.RemoveFromCartRequest, callOptions ...callopt.Option) (resp *cart.RemoveFromCartResponse, err error) {
	resp, err = defaultClient.RemoveFromCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "RemoveFromCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateItemQuantity(ctx context.Context, req *cart.UpdateQuantityRequest, callOptions ...callopt.Option) (resp *cart.UpdateQuantityResponse, err error) {
	resp, err = defaultClient.UpdateItemQuantity(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateItemQuantity call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ClearCart(ctx context.Context, req *cart.ClearCartRequest, callOptions ...callopt.Option) (resp *cart.ClearCartResponse, err error) {
	resp, err = defaultClient.ClearCart(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ClearCart call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
