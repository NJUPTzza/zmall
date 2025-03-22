package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/joho/godotenv"
)

func TestRegister_Run(t *testing.T) {
	godotenv.Load("../../.env")
	mysql.Init()
	ctx := context.Background()
	s := NewRegisterService(ctx)
	// init req and assert value

	req := &user.RegisterRequest{
		Username: "zza",
		Password: "123456",
		PasswordConfirm: "123456",
		Email: "921212318@qq.com",
		Phone: "18915565387",
	}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestRegisterService_Run(t *testing.T) {
	type args struct {
		req *user.RegisterRequest
	}
	tests := []struct {
		name     string
		s        *RegisterService
		args     args
		wantResp *user.RegisterResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("RegisterService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
