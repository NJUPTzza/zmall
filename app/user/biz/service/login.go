package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/user/biz/model"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	ctx context.Context
} // NewLoginService new LoginService
func NewLoginService(ctx context.Context) *LoginService {
	return &LoginService{ctx: ctx}
}

// Run create note info
func (s *LoginService) Run(req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// Finish your business logic.
	if req.Username == "" || req.Password == "" {
		return nil, errors.New("username or password is empty")
	}
	row, err := model.GetByUsername(mysql.DB, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResponse{
		Token: "token",
	}
	return resp, nil
}
