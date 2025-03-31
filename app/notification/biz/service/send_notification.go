package service

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
)

type SendNotificationService struct {
	ctx context.Context
} // NewSendNotificationService new SendNotificationService
func NewSendNotificationService(ctx context.Context) *SendNotificationService {
	return &SendNotificationService{ctx: ctx}
}

// Run create note info
func (s *SendNotificationService) Run(req *notification.SendNotificationRequest) (resp *notification.SendNotificationResponse, err error) {
	// Finish your business logic.

	return
}
