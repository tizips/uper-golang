package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/uper-go/admin/biz/basic"
	"github.com/tizips/uper-go/admin/middleware"
)

func BasicRouter(router *server.Hertz) {

	route := router.Group("basic")
	{
		login := route.Group("login")
		{
			login.POST("account", basic.DoLoginByAccount)
		}

		account := route.Group("account").Use(middleware.Auth())
		{
			account.GET("information", basic.ToAccountByInformation)
			account.GET("modules", basic.ToAccountByModules)
			account.GET("permissions", basic.ToAccountByPermissions)
			account.POST("logout", middleware.Auth(), basic.DoLoginByOut)
		}

		upload := route.Group("upload").Use(middleware.Auth())
		{
			upload.POST("file", basic.DoUploadByFile)
		}
	}
}
