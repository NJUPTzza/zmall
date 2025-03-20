package service

import (
	"context"
	"testing"
	cart "github.com/NJUPTzza/zmall/client/cart/kitex_gen/cart"
)

func TestDeleteItem_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteItemService(ctx)
	// init req and assert value

	req := &cart.DeleteItemReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
