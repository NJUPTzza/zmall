package service

import (
	"context"
	product "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product"
	"testing"
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
