package main

import (
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http/middleware"
	"github.com/tizips/uper-go/web/bootstrap"
	"github.com/tizips/uper-go/web/route"
)

func main() {

	bootstrap.Boot()

	if facades.Server != nil {

		facades.Server.Use(middleware.AccessMiddleware())

		route.Router(facades.Server)

		facades.Server.Spin()
	}
}
