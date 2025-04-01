package mw

import (
	"context"
	"errors"
	"time"

	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/jwt"
)

var (
	JwtMiddleware *jwt.HertzJWTMiddleware
	IdentityKey   = "user_id"
)

// 初始化 JWT 中间件
func InitJwt(userClient userservice.Client) {
	var err error
	JwtMiddleware, err = jwt.New(&jwt.HertzJWTMiddleware{
		Key:         []byte("secret key"), // 设置 JWT 密钥
		Timeout:     time.Hour,            // 令牌有效期
		MaxRefresh:  time.Hour,            // 令牌刷新时间
		IdentityKey: IdentityKey,

		// 用户登录时的认证逻辑
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginStruct struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}
			if err := c.BindAndValidate(&loginStruct); err != nil {
				return nil, err
			}
			resp, err := userClient.CheckUser(ctx, &user.CheckUserRequest{
				Username: loginStruct.Username,
				Password: loginStruct.Password,
			})

			if err != nil || resp.CommonResponse.Code != 200 {
				return nil, errors.New("用户名或密码错误")
			}

			return map[string]interface{}{
				"user_id": resp.Id,
				"username": loginStruct.Username,
			}, nil
		},

		// 生成 JWT 载荷
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(map[string]interface{}); ok {
				return jwt.MapClaims{
					IdentityKey: v["user_id"],
					"usernmae": v["username"],
				}
			}
			return jwt.MapClaims{}
		},

		// 解析 JWT，获取用户信息
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return map[string]interface{}{
				"user_id": claims[IdentityKey],
				"username": claims["username"],
			}
		},

		// 认证失败处理
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(401, map[string]interface{}{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		panic(err)
	}
}