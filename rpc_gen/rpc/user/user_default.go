package user

import (
	"context"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (resp *user.RegisterResponse, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.LoginRequest, callOptions ...callopt.Option) (resp *user.LoginResponse, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUser(ctx context.Context, req *user.GetUserRequest, callOptions ...callopt.Option) (resp *user.GetUserResponse, err error) {
	resp, err = defaultClient.GetUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckUser(ctx context.Context, req *user.CheckUserRequest, callOptions ...callopt.Option) (resp *user.CheckUserResponse, err error) {
	resp, err = defaultClient.CheckUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
