package service

import (
	"context"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// Finish your business logic.
	return
}
