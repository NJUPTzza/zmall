package handler

import (
	"context"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(userClient userservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.RegisterRequest

		if err := c.Bind(&req); err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		klog.Infof("Request: %+v", &req)

		resp, err := userClient.Register(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(consts.StatusOK, resp)
	}
}

