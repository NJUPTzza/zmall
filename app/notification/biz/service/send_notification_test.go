package service

import (
	"context"
	"testing"
	notification "github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/notification"
)

func TestSendNotification_Run(t *testing.T) {
	ctx := context.Background()
	s := NewSendNotificationService(ctx)
	// init req and assert value

	req := &notification.SendNotificationRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
