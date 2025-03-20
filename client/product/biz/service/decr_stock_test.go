package service

import (
	"context"
	"testing"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

func TestDecrStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDecrStockService(ctx)
	// init req and assert value

	req := &product.DecrStockReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
