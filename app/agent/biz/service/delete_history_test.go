package service

import (
	"context"
	"testing"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

func TestDeleteHistory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewDeleteHistoryService(ctx)
	// init req and assert value

	req := &agent.DeleteHistoryRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
