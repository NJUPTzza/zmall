// Code generated by Kitex v0.9.1. DO NOT EDIT.

package notificationservice

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SendNotification(ctx context.Context, Req *notification.SendNotificationRequest, callOptions ...callopt.Option) (r *notification.SendNotificationResponse, err error)
	GetNotifications(ctx context.Context, Req *notification.GetNotificationsRequest, callOptions ...callopt.Option) (r *notification.GetNotificationsResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kNotificationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kNotificationServiceClient struct {
	*kClient
}

func (p *kNotificationServiceClient) SendNotification(ctx context.Context, Req *notification.SendNotificationRequest, callOptions ...callopt.Option) (r *notification.SendNotificationResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendNotification(ctx, Req)
}

func (p *kNotificationServiceClient) GetNotifications(ctx context.Context, Req *notification.GetNotificationsRequest, callOptions ...callopt.Option) (r *notification.GetNotificationsResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetNotifications(ctx, Req)
}
