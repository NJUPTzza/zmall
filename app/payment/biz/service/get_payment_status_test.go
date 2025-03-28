package service

import (
	"context"
	"testing"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

func TestGetPaymentStatus_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetPaymentStatusService(ctx)
	// init req and assert value

	req := &payment.GetPaymentStatusRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
