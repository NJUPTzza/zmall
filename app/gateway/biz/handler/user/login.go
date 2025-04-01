package handler

import (
	"context"

	"github.com/NJUPTzza/zmall/app/gateway/biz/mw"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Login(userClient userservice.Client) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		var req user.LoginRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(consts.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}

		resp, err := userClient.Login(ctx, &req)
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}

		// 生成jwt
		jwtToken, _, err := mw.JwtMiddleware.TokenGenerator(map[string]interface{}{
			"user_id":  resp.User.Id,
			"username": resp.User.Username,
		})
		if err != nil {
			c.JSON(consts.StatusInternalServerError, map[string]string{"error": "生成Token失败"})
			return
		}

		// 返回 Token
		c.JSON(consts.StatusOK, map[string]string{
			"token": jwtToken,
		})
	}
}

