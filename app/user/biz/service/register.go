package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	"github.com/NJUPTzza/zmall/app/user/biz/model"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// Finish your business logic.
	if req.Username == "" || req.Password == "" || req.PasswordConfirm == "" || req.Email == "" || req.Phone == "" {
		return nil, errors.New("all fields are required")
	}	

	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password and password confirmation do not match")
	}
	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Username: req.Username,
		PasswordHashed: string(passwordHashed),
		Email: req.Email,
		Phone: req.Phone,
	}
	err = mysql.Create(mysql.DB, newUser)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		User: &user.User{
			Id: int64(newUser.ID),
			Username: req.Username,
			Email: req.Email,
			Phone: req.Phone,
		},
		CommonResponse: &user.CommonResponse {
			Code: 200,
			Message: "register success",
		},	
	}, nil
}
