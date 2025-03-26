package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

// register registers all routers.
func Register(r *server.Hertz) {
	userRegister(r)
	productRegister(r)
}
