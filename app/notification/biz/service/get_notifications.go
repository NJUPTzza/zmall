package service

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
)

type GetNotificationsService struct {
	ctx context.Context
} // NewGetNotificationsService new GetNotificationsService
func NewGetNotificationsService(ctx context.Context) *GetNotificationsService {
	return &GetNotificationsService{ctx: ctx}
}

// Run create note info
func (s *GetNotificationsService) Run(req *notification.GetNotificationsRequest) (resp *notification.GetNotificationsResponse, err error) {
	// Finish your business logic.

	return
}
