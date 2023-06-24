package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/uper-go/web/biz/blog"
)

func BlogRouter(routes *server.Hertz) {

	route := routes.Group("blog")
	{

		categories := route.Group("categories")
		{
			categories.GET("", blog.ToCategories)
			categories.GET(":id", blog.ToCategory)
		}

		articles := route.Group("articles")
		{
			articles.GET("", blog.ToArticles)
			articles.GET(":id", blog.ToArticle)
		}

		links := route.Group("links")
		{
			links.GET("", blog.ToLinks)
		}

		search := route.Group("search")
		{
			search.GET("", blog.ToSearch)
		}

		setting := route.Group("setting")
		{
			setting.GET("", blog.ToSetting)
		}

	}
}
