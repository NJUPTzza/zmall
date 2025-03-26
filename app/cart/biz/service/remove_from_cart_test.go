package service

import (
	"context"
	"testing"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

func TestRemoveFromCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewRemoveFromCartService(ctx)
	// init req and assert value

	req := &cart.RemoveFromCartRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
