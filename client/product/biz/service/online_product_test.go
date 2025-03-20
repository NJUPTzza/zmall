package service

import (
	"context"
	"testing"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

func TestOnlineProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOnlineProductService(ctx)
	// init req and assert value

	req := &product.OnlineProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
