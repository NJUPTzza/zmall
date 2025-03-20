package service

import (
	"context"
	"testing"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

func TestIncrStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewIncrStockService(ctx)
	// init req and assert value

	req := &product.IncrStockReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
