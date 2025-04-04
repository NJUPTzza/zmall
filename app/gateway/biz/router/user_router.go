// Code generated by hertz generator.

package router

import (
	"log"

	handler "github.com/NJUPTzza/zmall/app/gateway/biz/handler/user"
	"github.com/NJUPTzza/zmall/app/gateway/biz/mw"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

// customizeRegister registers customize routers.
func userRegister(r *server.Hertz) {
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := userservice.NewClient("user_service", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

	// 初始化 JWT 中间件
	mw.InitJwt(client)

	r.POST("/api/v1/register", handler.Register(client))
	r.POST("/api/v1/login", handler.Login(client))
	r.GET("/api/v1/user", handler.GetUser(client))
}
