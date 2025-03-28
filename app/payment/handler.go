package main

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
	"github.com/NJUPTzza/zmall/app/payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// ProcessPayment implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) ProcessPayment(ctx context.Context, req *payment.ProcessPaymentRequest) (resp *payment.ProcessPaymentResponse, err error) {
	resp, err = service.NewProcessPaymentService(ctx).Run(req)

	return resp, err
}

// GetPaymentStatus implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) GetPaymentStatus(ctx context.Context, req *payment.GetPaymentStatusRequest) (resp *payment.GetPaymentStatusResponse, err error) {
	resp, err = service.NewGetPaymentStatusService(ctx).Run(req)

	return resp, err
}
