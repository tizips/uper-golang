package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Router(router *server.Hertz) {

	BlogRouter(router)

}
