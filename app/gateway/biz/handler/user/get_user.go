package handler

import (
	"context"
	"strconv"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func GetUser(userClient userservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		idStr := c.Query("id")
		if idStr == "" {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid id"})
			return
		}

		req := user.GetUserRequest {
			Id: int64(id),
		}

		resp, err := userClient.GetUser(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		c.JSON(consts.StatusOK, resp)
	}
}

