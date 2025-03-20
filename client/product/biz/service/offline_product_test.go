package service

import (
	"context"
	"testing"
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
)

func TestOfflineProduct_Run(t *testing.T) {
	ctx := context.Background()
	s := NewOfflineProductService(ctx)
	// init req and assert value

	req := &product.OfflineProductReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
