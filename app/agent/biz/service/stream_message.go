package service

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

type StreamMessageService struct {
	ctx context.Context
}

// NewStreamMessageService new StreamMessageService
func NewStreamMessageService(ctx context.Context) *StreamMessageService {
	return &StreamMessageService{ctx: ctx}
}

func (s *StreamMessageService) Run(req *agent.ChatRequest, stream agent.AgentService_StreamMessageServer) (err error) {
	return
}
