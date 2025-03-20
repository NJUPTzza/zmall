package service

import (
	"context"
	"testing"
	cart "github.com/NJUPTzza/zmall/client/cart/kitex_gen/cart"
)

func TestUpdateCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateCartService(ctx)
	// init req and assert value

	req := &cart.UpdateCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
