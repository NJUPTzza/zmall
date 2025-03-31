package main

import (
	"context"

	"github.com/NJUPTzza/zmall/app/payment/biz/service"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order/orderservice"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{
	orderClient orderservice.Client
}

// ProcessPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) ProcessPayment(ctx context.Context, req *payment.ProcessPaymentRequest) (resp *payment.ProcessPaymentResponse, err error) {
	resp, err = service.NewProcessPaymentService(ctx, s.orderClient).Run(req)

	return resp, err
}

// GetPaymentStatus implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetPaymentStatus(ctx context.Context, req *payment.GetPaymentStatusRequest) (resp *payment.GetPaymentStatusResponse, err error) {
	resp, err = service.NewGetPaymentStatusService(ctx).Run(req)

	return resp, err
}

// UpdatePaymentStatus implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) UpdatePaymentStatus(ctx context.Context, req *payment.UpdatePaymentStatusRequest) (resp *payment.UpdatePaymentStatusResponse, err error) {
	resp, err = service.NewUpdatePaymentStatusService(ctx).Run(req)

	return resp, err
}
