package main

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
	"github.com/NJUPTzza/zmall/app/notification/biz/service"
)

// NotificationServiceImpl implements the last service interface defined in the IDL.
type NotificationServiceImpl struct{}

// SendNotification implements the NotificationServiceImpl interface.
func (s *NotificationServiceImpl) SendNotification(ctx context.Context, req *notification.SendNotificationRequest) (resp *notification.SendNotificationResponse, err error) {
	resp, err = service.NewSendNotificationService(ctx).Run(req)

	return resp, err
}

// GetNotifications implements the NotificationServiceImpl interface.
func (s *NotificationServiceImpl) GetNotifications(ctx context.Context, req *notification.GetNotificationsRequest) (resp *notification.GetNotificationsResponse, err error) {
	resp, err = service.NewGetNotificationsService(ctx).Run(req)

	return resp, err
}
