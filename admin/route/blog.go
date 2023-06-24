package route

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/tizips/uper-go/admin/biz/blog"
	"github.com/tizips/uper-go/admin/middleware"
)

func BlogRouter(routes *server.Hertz) {

	route := routes.Group("blog")
	{

		categories := route.Group("categories")
		{
			categories.GET(":id", blog.ToCategoryByInformation)
			categories.GET("", middleware.Permission("blog.category.tree"), blog.ToCategories)
			categories.PUT(":id", middleware.Permission("blog.category.update"), blog.DoCategoryByUpdate)
			categories.DELETE(":id", middleware.Permission("blog.category.delete"), blog.DoCategoryByDelete)
		}

		category := route.Group("category")
		{
			category.POST("", middleware.Permission("blog.category.create"), blog.DoCategoryByCreate)
			category.PUT("enable", middleware.Permission("blog.category.enable"), blog.DoCategoryByEnable)
			category.GET("parent", blog.ToCategoryByParent)
			category.GET("opening", blog.ToCategoryByOpening)
		}

		articles := route.Group("articles")
		{
			articles.GET("", middleware.Permission("blog.article.paginate"), blog.ToArticleByPaginate)
			articles.GET(":id", blog.ToArticleByInformation)
			articles.PUT(":id", middleware.Permission("blog.article.update"), blog.DoArticleByUpdate)
			articles.DELETE(":id", middleware.Permission("blog.article.delete"), blog.DoArticleByDelete)
		}

		article := route.Group("article")
		{
			article.POST("", middleware.Permission("blog.article.create"), blog.DoArticleByCreate)
			article.PUT("enable", middleware.Permission("blog.article.enable"), blog.DoArticleByEnable)
		}

		links := route.Group("links")
		{
			links.GET("", middleware.Permission("blog.link.paginate"), blog.ToLinkByPaginate)
			links.PUT(":id", middleware.Permission("blog.link.update"), blog.DoLinkByUpdate)
			links.DELETE(":id", middleware.Permission("blog.link.delete"), blog.DoLinkByDelete)
		}

		link := route.Group("link")
		{
			link.POST("", middleware.Permission("blog.link.create"), blog.DoLinkByCreate)
			link.PUT("enable", middleware.Permission("blog.link.enable"), blog.DoLinkByEnable)
		}

		setting := route.Group("setting")
		{
			setting.GET("", middleware.Permission("blog.setting.list"), blog.ToSetting)
			setting.PUT("", middleware.Permission("blog.setting.update"), blog.DoSetting)
		}

	}
}
