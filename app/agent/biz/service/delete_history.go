package service

import (
	"context"
	agent "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/agent"
)

type DeleteHistoryService struct {
	ctx context.Context
} // NewDeleteHistoryService new DeleteHistoryService
func NewDeleteHistoryService(ctx context.Context) *DeleteHistoryService {
	return &DeleteHistoryService{ctx: ctx}
}

// Run create note info
func (s *DeleteHistoryService) Run(req *agent.DeleteHistoryRequest) (resp *agent.DeleteHistoryResponse, err error) {
	// Finish your business logic.

	return
}
