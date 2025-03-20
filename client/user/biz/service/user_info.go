package service

import (
	"context"
	user "github.com/NJUPTzza/zmall/client/user/kitex_gen/user"
)

type UserInfoService struct {
	ctx context.Context
} // NewUserInfoService new UserInfoService
func NewUserInfoService(ctx context.Context) *UserInfoService {
	return &UserInfoService{ctx: ctx}
}

// Run create note info
func (s *UserInfoService) Run(req *user.UserInfoReq) (resp *user.UserInfoResp, err error) {
	// Finish your business logic.

	return
}
