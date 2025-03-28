package service

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

type GetPaymentStatusService struct {
	ctx context.Context
} // NewGetPaymentStatusService new GetPaymentStatusService
func NewGetPaymentStatusService(ctx context.Context) *GetPaymentStatusService {
	return &GetPaymentStatusService{ctx: ctx}
}

// Run create note info
func (s *GetPaymentStatusService) Run(req *payment.GetPaymentStatusRequest) (resp *payment.GetPaymentStatusResponse, err error) {
	// Finish your business logic.

	return
}
