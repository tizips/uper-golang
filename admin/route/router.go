package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/uper-go/admin/middleware"
)

func Router(router *server.Hertz) {

	router.Use(middleware.Jwt())

	BasicRouter(router)

	SiteRouter(router)

	BlogRouter(router)

}
