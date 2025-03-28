package service

import (
	"context"
	"testing"
	payment "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment"
)

func TestProcessPayment_Run(t *testing.T) {
	ctx := context.Background()
	s := NewProcessPaymentService(ctx)
	// init req and assert value

	req := &payment.ProcessPaymentRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
