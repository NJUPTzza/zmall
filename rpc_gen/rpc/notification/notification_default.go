package notification

import (
	"context"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func SendNotification(ctx context.Context, req *notification.SendNotificationRequest, callOptions ...callopt.Option) (resp *notification.SendNotificationResponse, err error) {
	resp, err = defaultClient.SendNotification(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "SendNotification call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetNotifications(ctx context.Context, req *notification.GetNotificationsRequest, callOptions ...callopt.Option) (resp *notification.GetNotificationsResponse, err error) {
	resp, err = defaultClient.GetNotifications(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetNotifications call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
