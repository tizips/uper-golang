package blog

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/go-redsync/redsync/v4"
	"github.com/gookit/goutil/strutil"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/helper/authorize"
	"github.com/tizips/uper-go/admin/http/request/blog"
	res "github.com/tizips/uper-go/admin/http/response/blog"
	blogService "github.com/tizips/uper-go/admin/service/blog"
	"github.com/tizips/uper-go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func ToArticleByPaginate(ctx context.Context, c *app.RequestContext) {

	var request blog.ToArticleByPaginate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	responses := response.Paginate[res.ToArticleByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm

	tx.Model(model.BlgArticle{}).Count(&responses.Total)

	if responses.Total > 0 {

		var articles []model.BlgArticle

		tx.
			Preload("Category", func(tf *gorm.DB) *gorm.DB { return tf.Unscoped() }).
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Order("`created_at` desc").
			Find(&articles)

		responses.Data = make([]res.ToArticleByPaginate, len(articles))

		for index, item := range articles {
			responses.Data[index] = res.ToArticleByPaginate{
				ID:        item.ID,
				Name:      item.Name,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
			if item.Category != nil {
				responses.Data[index].Category = item.Category.Name
			}
		}
	}

	http.Success(c, responses)
}

func DoArticleByCreate(ctx context.Context, c *app.RequestContext) {

	var request blog.DoArticleByCreate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var category model.BlgCategory

	fc := facades.Gorm.First(&category, "`id`=?", request.Category)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该栏目")
		return
	} else if fc.Error != nil {
		http.Fail(c, "栏目查询失败：%v", fc.Error)
		return
	}

	if category.Type != model.BlogCategoryForTypeOfList {
		http.Fail(c, "该栏目非列表类型，无法添加文章")
		return
	}

	tx := facades.Gorm.Begin()

	article := model.BlgArticle{
		ID:         facades.Snowflake.Generate().String(),
		CategoryID: request.Category,
		UserID:     authorize.ID(c),
		Name:       request.Name,
		Picture:    request.Picture,
		Source:     request.Source,
		URL:        request.URL,
		IsComment:  request.IsComment,
		IsEnable:   request.IsEnable,
	}

	if length := strutil.Utf8Len(request.Text); length >= 210 {
		article.Summary = strutil.Substr(request.Text, 0, 200) + "..."
	} else {
		article.Summary = request.Text
	}

	if ca := tx.Create(&article); ca.Error != nil {
		tx.Rollback()
		http.Fail(c, "写入失败：%v", ca.Error)
		return
	}

	seo := model.BlgSEO{
		Type:        constants.BlogTypeOfArticle,
		OtherID:     article.ID,
		Title:       request.Title,
		Keyword:     request.Keyword,
		Description: request.Description,
	}

	if cs := tx.Create(&seo); cs.Error != nil {
		tx.Rollback()
		http.Fail(c, "写入失败：%v", cs.Error)
		return
	}

	html := model.BlgHTML{
		Type:    constants.BlogTypeOfArticle,
		OtherID: article.ID,
		Content: request.Content,
		Text:    request.Text,
	}

	if ct := tx.Create(&html); ct.Error != nil {
		tx.Rollback()
		http.Fail(c, "写入失败：%v", ct.Error)
		return
	}

	if article.IsEnable == util.EnableOfYes {

		documents := []map[string]any{
			{
				"id":          article.ID,
				"name":        article.Name,
				"picture":     article.Picture,
				"title":       seo.Title,
				"keyword":     seo.Keyword,
				"description": seo.Description,
				"text":        html.Text,
				"created_at":  article.CreatedAt.ToDateTimeString(),
			},
		}

		_, err := blogService.SearchIndexForArticle().AddDocuments(documents, "id")

		if err != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", err)
			return
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoArticleByUpdate(ctx context.Context, c *app.RequestContext) {

	var request blog.DoArticleByUpdate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	mutex := facades.Locker.NewMutex(locker.Keys("article", request.ID))

	if err := mutex.Lock(); err != nil {
		http.Fail(c, "处理失败：%v", err)
		return
	}

	defer func(lock *redsync.Mutex) {
		_, _ = lock.Unlock()
	}(mutex)

	var article model.BlgArticle

	fa := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfArticle) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfArticle) }).
		First(&article, "`id`=?", request.ID)

	if errors.Is(fa.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该文章")
		return
	} else if fa.Error != nil {
		http.Fail(c, "文章查询失败：%v", fa.Error)
		return
	}

	if request.Category != article.CategoryID {

		var category model.BlgCategory

		fc := facades.Gorm.First(&category, "`id`=?", request.Category)

		if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
			http.NotFound(c, "未找到该栏目")
			return
		} else if fc.Error != nil {
			http.Fail(c, "栏目查询失败：%v", fc.Error)
			return
		}

		if category.Type != model.BlogCategoryForTypeOfList {
			http.Fail(c, "该栏目非列表类型，无法添加文章")
			return
		}
	}

	tx := facades.Gorm.Begin()

	article.CategoryID = request.Category
	article.Name = request.Name
	article.Source = request.Source
	article.URL = request.URL
	article.Picture = request.Picture
	article.IsEnable = request.IsEnable
	article.IsComment = request.IsComment

	if length := strutil.Utf8Len(request.Text); length >= 210 {
		article.Summary = strutil.Substr(request.Text, 0, 200) + "..."
	} else {
		article.Summary = request.Text
	}

	if ua := tx.Omit(clause.Associations).Save(&article); ua.Error != nil {
		tx.Rollback()
		http.Fail(c, "修改失败：%v", ua.Error)
		return
	}

	if article.SEO != nil {

		article.SEO.Title = request.Title
		article.SEO.Keyword = request.Keyword
		article.SEO.Description = request.Description

		if us := tx.Save(&article.SEO); us.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", us.Error)
			return
		}
	}

	if article.HTML != nil {

		article.HTML.Content = request.Content
		article.HTML.Text = request.Text

		if ut := tx.Save(&article.HTML); ut.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", ut.Error)
			return
		}
	}

	if article.IsEnable == util.EnableOfYes {

		documents := []map[string]any{
			{
				"id":          article.ID,
				"name":        article.Name,
				"picture":     article.Picture,
				"title":       article.SEO.Title,
				"keyword":     article.SEO.Keyword,
				"description": article.SEO.Description,
				"text":        request.Text,
				"created_at":  article.CreatedAt.ToDateTimeString(),
			},
		}

		if _, err := blogService.SearchIndexForArticle().AddDocuments(documents, "id"); err != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", err)
			return
		}
	} else {
		if _, err := blogService.SearchIndexForArticle().DeleteDocument(article.ID); err != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", err)
			return
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoArticleByEnable(ctx context.Context, c *app.RequestContext) {

	var request blog.DoArticleByEnable

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var article model.BlgArticle

	fc := facades.Gorm.First(&article, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该文章")
		return
	} else if fc.Error != nil {
		http.Fail(c, "文章查询失败：%v", fc.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if ec := tx.Model(&article).Update("is_enable", request.IsEnable); ec.Error != nil {
		tx.Rollback()
		http.Fail(c, "操作失败：%v", ec.Error)
		return
	}

	if article.IsEnable == util.EnableOfYes {

		_ = facades.Gorm.Model(&article).Association("SEO").Find(&article.SEO)
		_ = facades.Gorm.Model(&article).Association("HTML").Find(&article.HTML)

		documents := []map[string]any{
			{
				"id":          article.ID,
				"name":        article.Name,
				"picture":     article.Picture,
				"title":       article.SEO.Title,
				"keyword":     article.SEO.Keyword,
				"description": article.SEO.Description,
				"text":        article.HTML.Text,
				"created_at":  article.CreatedAt.ToDateTimeString(),
			},
		}

		if _, err := blogService.SearchIndexForArticle().AddDocuments(documents, "id"); err != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", err)
			return
		}
	} else {
		if _, err := blogService.SearchIndexForArticle().DeleteDocument(article.ID); err != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", err)
			return
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoArticleByDelete(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var article model.BlgArticle

	fc := facades.Gorm.First(&article, "`id`=?", id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该文章")
		return
	} else if fc.Error != nil {
		http.Fail(c, "文章查询失败：%v", fc.Error)
		return
	}

	tx := facades.Gorm.Begin()

	if dc := tx.Delete(&article); dc.Error != nil {
		tx.Rollback()
		http.Fail(c, "删除失败：%v", dc.Error)
		return
	}

	if _, err := blogService.SearchIndexForArticle().DeleteDocument(article.ID); err != nil {
		tx.Rollback()
		http.Fail(c, "写入失败：%v", err)
		return
	}

	tx.Commit()

	http.Success[any](c)
}

func ToArticleByInformation(ctx context.Context, c *app.RequestContext) {

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

	responses := res.ToArticleByInformation[[]string]{
		ID:        article.ID,
		Name:      article.Name,
		Category:  make([]string, 0),
		Source:    article.Source,
		URL:       article.URL,
		Picture:   article.Picture,
		IsComment: article.IsComment,
		IsEnable:  article.IsEnable,
	}

	if article.Category != nil {
		if article.Category.ParentID != nil {
			responses.Category = append(responses.Category, *article.Category.ParentID)
		}
		responses.Category = append(responses.Category, article.CategoryID)
	}

	if article.SEO != nil {
		responses.Title = article.SEO.Title
		responses.Keyword = article.SEO.Keyword
		responses.Description = article.SEO.Description
	}

	if article.HTML != nil {
		responses.Content = article.HTML.Content
	}

	http.Success(c, responses)
}
