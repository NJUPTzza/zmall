package service

import (
	"context"
	user "github.com/NJUPTzza/zmall/client/user/kitex_gen/user"
)

type DeleteUserService struct {
	ctx context.Context
} // NewDeleteUserService new DeleteUserService
func NewDeleteUserService(ctx context.Context) *DeleteUserService {
	return &DeleteUserService{ctx: ctx}
}

// Run create note info
func (s *DeleteUserService) Run(req *user.UserDeleteReq) (resp *user.UserDeleteResp, err error) {
	// Finish your business logic.

	return
}
