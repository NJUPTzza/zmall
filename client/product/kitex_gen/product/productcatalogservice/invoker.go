// Code generated by Kitex v0.9.1. DO NOT EDIT.

package productcatalogservice

import (
	product "github.com/NJUPTzza/zmall/client/product/kitex_gen/product"
	server "github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler product.ProductCatalogService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
