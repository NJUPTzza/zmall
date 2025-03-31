package notification

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification/notificationservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() notificationservice.Client
	Service() string
	SendNotification(ctx context.Context, Req *notification.SendNotificationRequest, callOptions ...callopt.Option) (r *notification.SendNotificationResponse, err error)
	GetNotifications(ctx context.Context, Req *notification.GetNotificationsRequest, callOptions ...callopt.Option) (r *notification.GetNotificationsResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := notificationservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient notificationservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() notificationservice.Client {
	return c.kitexClient
}

func (c *clientImpl) SendNotification(ctx context.Context, Req *notification.SendNotificationRequest, callOptions ...callopt.Option) (r *notification.SendNotificationResponse, err error) {
	return c.kitexClient.SendNotification(ctx, Req, callOptions...)
}

func (c *clientImpl) GetNotifications(ctx context.Context, Req *notification.GetNotificationsRequest, callOptions ...callopt.Option) (r *notification.GetNotificationsResponse, err error) {
	return c.kitexClient.GetNotifications(ctx, Req, callOptions...)
}
