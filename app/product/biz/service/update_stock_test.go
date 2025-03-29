package service

import (
	"context"
	"testing"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
)

func TestUpdateStock_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUpdateStockService(ctx)
	// init req and assert value

	req := &product.UpdateStockRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
