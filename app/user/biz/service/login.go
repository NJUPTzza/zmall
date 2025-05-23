package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
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
	row, err := mysql.GetByUsername(mysql.DB, req.Username)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHashed), []byte(req.Password))
	if err != nil {
		return nil, err
	}
	resp = &user.LoginResponse{
		CommonResponse: &user.CommonResponse{
			Code: 200,
			Message:  "success",
		},
		User: &user.User{
			Id: int64(row.ID),
			Username: row.Username,
			Email: row.Email,
			Phone: row.Phone,
		},
	}
	return resp, nil
}
