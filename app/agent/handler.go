package main

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
	"github.com/NJUPTzza/zmall/app/agent/biz/service"
)

// AgentServiceImpl implements the last service interface defined in the IDL.
type AgentServiceImpl struct{}

// SendMessage implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) SendMessage(ctx context.Context, req *agent.ChatRequest) (resp *agent.ChatResponse, err error) {
	resp, err = service.NewSendMessageService(ctx).Run(req)

	return resp, err
}

func (s *AgentServiceImpl) StreamMessage(req *agent.ChatRequest, stream agent.AgentService_StreamMessageServer) (err error) {
	ctx := context.Background()
	err = service.NewStreamMessageService(ctx).Run(req, stream)
	return
}

// GetHistory implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) GetHistory(ctx context.Context, req *agent.GetHistoryRequest) (resp *agent.GetHistoryResponse, err error) {
	resp, err = service.NewGetHistoryService(ctx).Run(req)

	return resp, err
}

// DeleteHistory implements the AgentServiceImpl interface.
func (s *AgentServiceImpl) DeleteHistory(ctx context.Context, req *agent.DeleteHistoryRequest) (resp *agent.DeleteHistoryResponse, err error) {
	resp, err = service.NewDeleteHistoryService(ctx).Run(req)

	return resp, err
}
