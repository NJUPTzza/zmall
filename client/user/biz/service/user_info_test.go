package service

import (
	"context"
	"testing"
	user "github.com/NJUPTzza/zmall/client/user/kitex_gen/user"
)

func TestUserInfo_Run(t *testing.T) {
	ctx := context.Background()
	s := NewUserInfoService(ctx)
	// init req and assert value

	req := &user.UserInfoReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
