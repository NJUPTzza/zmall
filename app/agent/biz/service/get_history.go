package service

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

type GetHistoryService struct {
	ctx context.Context
} // NewGetHistoryService new GetHistoryService
func NewGetHistoryService(ctx context.Context) *GetHistoryService {
	return &GetHistoryService{ctx: ctx}
}

// Run create note info
func (s *GetHistoryService) Run(req *agent.GetHistoryRequest) (resp *agent.GetHistoryResponse, err error) {
	// Finish your business logic.

	return
}
