package service

import (
	"context"
	"testing"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
)

func TestCheckUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCheckUserService(ctx)
	// init req and assert value

	req := &user.CheckUserRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test
	
}
