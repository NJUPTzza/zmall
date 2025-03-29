package service

import (
	"context"
	"errors"

	"github.com/NJUPTzza/zmall/app/user/biz/dal/mysql"
	user "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"gorm.io/gorm"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.GetUserRequest) (resp *user.GetUserResponse, err error) {
	// Finish your business logic.
	if req.Id == 0 {
		return nil, errors.New("id is empty")
	}
	userDetail, err := mysql.GetUserById(mysql.DB, req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &user.GetUserResponse{
				CommonResponse: &user.CommonResponse{
					Code:    404,
					Message: "user not found",
				},
			}, nil
		}
		return nil, err
	}

	return &user.GetUserResponse{
		User: &user.User{
			Id: int64(userDetail.ID),
			Username: userDetail.Username,
			Email: userDetail.Email,
			Phone: userDetail.Phone,
		},
		CommonResponse: &user.CommonResponse{
			Code: 200,
			Message: "get user detail success",
		},
	}, nil
}
