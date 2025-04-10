package service

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

type SendMessageService struct {
	ctx context.Context
} // NewSendMessageService new SendMessageService
func NewSendMessageService(ctx context.Context) *SendMessageService {
	return &SendMessageService{ctx: ctx}
}

// Run create note info
func (s *SendMessageService) Run(req *agent.ChatRequest) (resp *agent.ChatResponse, err error) {
	// Finish your business logic.

	return
}
