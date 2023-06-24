package blog

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/model"
	"github.com/tizips/uper-go/web/constants"
	req "github.com/tizips/uper-go/web/http/request/blog"
	res "github.com/tizips/uper-go/web/http/response/blog"
	"gorm.io/gorm"
)

func ToArticles(ctx context.Context, c *app.RequestContext) {

	var request req.ToArticles

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	responses := response.Paginate[res.ToArticles]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm.Where("`is_enable`=?", util.EnableOfYes)

	if len(request.Categories) > 0 {
		tx = tx.Where("`category_id` IN (?)", request.Categories)
	}

	tx.Model(model.BlgArticle{}).Count(&responses.Total)

	if responses.Total > 0 {

		var articles []model.BlgArticle

		tx.
			Preload("Category", func(tf *gorm.DB) *gorm.DB { return tf.Unscoped() }).
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` desc").
			Find(&articles)

		responses.Data = make([]res.ToArticles, len(articles))

		for index, item := range articles {
			responses.Data[index] = res.ToArticles{
				ID:        item.ID,
				Name:      item.Name,
				Summary:   item.Summary,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
			if item.Category != nil {
				responses.Data[index].Category = item.Category.Name
			}
		}
	}

	http.Success(c, responses)
}

func ToArticle(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var article model.BlgArticle

	fc := facades.Gorm.
		Preload("Category", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfArticle) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfArticle) }).
		First(&article, "`id`=?", id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该文章")
		return
	} else if fc.Error != nil {
		http.Fail(c, "文章查询失败：%v", fc.Error)
		return
	}

	responses := res.ToArticle[[]string]{
		ID:        article.ID,
		Name:      article.Name,
		Source:    article.Source,
		URL:       article.URL,
		Picture:   article.Picture,
		IsComment: article.IsComment,
		IsEnable:  article.IsEnable,
		CreatedAt: article.CreatedAt.ToDateTimeString(),
	}

	if article.Category != nil {
		responses.Category = article.Category.Name
	}

	if article.SEO != nil {
		responses.Title = article.SEO.Title
		responses.Keyword = article.SEO.Keyword
		responses.Description = article.SEO.Description
	}

	if article.HTML != nil {
		responses.HTML = article.HTML.Content
	}

	http.Success(c, responses)
}
