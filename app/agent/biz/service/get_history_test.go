package service

import (
	"context"
	"testing"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

func TestGetHistory_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetHistoryService(ctx)
	// init req and assert value

	req := &agent.GetHistoryRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
