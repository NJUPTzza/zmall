package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CheckUserService struct {
	ctx context.Context
} // NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckUserService {
	return &CheckUserService{ctx: ctx}
}

// Run create note info
func (s *CheckUserService) Run(req *user.CheckUserRequest) (resp *user.CheckUserResponse, err error) {
	userData, err := mysql.GetByUsername(mysql.DB, req.Username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user.CheckUserResponse{
				CommonResponse: &user.CommonResponse{
					Code:    404,
					Message: "用户不存在",
				},
			}, nil
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.PasswordHashed), []byte(req.Password))
	if err != nil {
		return &user.CheckUserResponse{
			CommonResponse: &user.CommonResponse{
				Code:    401,
				Message: "密码错误",
			},
		}, nil
	}

	// 返回验证成功的用户信息
	return &user.CheckUserResponse{
		Id:       int64(userData.ID),
		Username: userData.Username,
		CommonResponse: &user.CommonResponse{
			Code:    200,
			Message: "验证成功",
		},
	}, nil
}
