package router

import (
	"log"

	handler "github.com/NJUPTzza/zmall/app/gateway/biz/handler/product"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/product/productservice"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func productRegister(r *server.Hertz) {
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379}"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := productservice.NewClient("product_service", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

	r.GET("/api/v1/product", handler.GetProduct(client))
	r.GET("api/v1/products", handler.ListProducts(client))
}