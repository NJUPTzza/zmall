package service

import (
	"context"
	"testing"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

func TestUpdatePaymentStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdatePaymentStatusService(ctx)
	// init req and assert value

	req := &payment.UpdatePaymentStatusRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
