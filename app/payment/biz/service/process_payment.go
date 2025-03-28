package service

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

type ProcessPaymentService struct {
	ctx context.Context
} // NewProcessPaymentService new ProcessPaymentService
func NewProcessPaymentService(ctx context.Context) *ProcessPaymentService {
	return &ProcessPaymentService{ctx: ctx}
}

// Run create note info
func (s *ProcessPaymentService) Run(req *payment.ProcessPaymentRequest) (resp *payment.ProcessPaymentResponse, err error) {
	// Finish your business logic.

	return
}
