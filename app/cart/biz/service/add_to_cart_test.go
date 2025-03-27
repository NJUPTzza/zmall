package service

import (
	"context"
	"testing"

	cart "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/cart"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
)

func TestAddToCart_Run(t *testing.T) {
	ctx := context.Background()
	cli, err := productservice.NewClient("")
	
	s := NewAddToCartService(ctx, cli)
	// init req and assert value

	req := &cart.AddToCartRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
