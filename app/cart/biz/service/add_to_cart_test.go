package service

import (
	"context"
	"testing"
	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
)

func TestAddToCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewAddToCartService(ctx)
	// init req and assert value

	req := &cart.AddToCartRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
