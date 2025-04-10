package service

import (
	"context"
	"testing"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

func TestSendMessage_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendMessageService(ctx)
	// init req and assert value

	req := &agent.ChatRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
