package payment

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func ProcessPayment(ctx context.Context, req *payment.ProcessPaymentRequest, callOptions ...callopt.Option) (resp *payment.ProcessPaymentResponse, err error) {
	resp, err = defaultClient.ProcessPayment(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ProcessPayment call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetPaymentStatus(ctx context.Context, req *payment.GetPaymentStatusRequest, callOptions ...callopt.Option) (resp *payment.GetPaymentStatusResponse, err error) {
	resp, err = defaultClient.GetPaymentStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetPaymentStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdatePaymentStatus(ctx context.Context, req *payment.UpdatePaymentStatusRequest, callOptions ...callopt.Option) (resp *payment.UpdatePaymentStatusResponse, err error) {
	resp, err = defaultClient.UpdatePaymentStatus(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdatePaymentStatus call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
