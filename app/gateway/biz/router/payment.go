package router

import (
	"log"

	handler "github.com/NJUPTzza/zmall/app/gateway/biz/handler/payment"
	"github.com/NJUPTzza/zmall/rpc_gen/kitex_gen/payment/paymentservice"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func paymentRegister(r *server.Hertz) {
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}
	client, err := paymentservice.NewClient("payment_service", client.WithResolver(resolver))
	if err != nil {
		log.Fatal(err)
	}

	r.POST("/api/v1/payment/process", handler.ProcessPayment(client))
	r.POST("/api/v1/payment/update", handler.UpdatePaymentStatus(client))

}