package service

import (
	"context"
	"testing"
	order "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/order"
)

func TestCreateOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateOrderService(ctx, nil)
	// init req and assert value

	req := &order.CreateOrderRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
