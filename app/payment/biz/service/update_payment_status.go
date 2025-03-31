package service

import (
	"context"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

type UpdatePaymentStatusService struct {
	ctx context.Context
} // NewUpdatePaymentStatusService new UpdatePaymentStatusService
func NewUpdatePaymentStatusService(ctx context.Context) *UpdatePaymentStatusService {
	return &UpdatePaymentStatusService{ctx: ctx}
}

// Run create note info
func (s *UpdatePaymentStatusService) Run(req *payment.UpdatePaymentStatusRequest) (resp *payment.UpdatePaymentStatusResponse, err error) {
	// Finish your business logic.
	
	return
}
