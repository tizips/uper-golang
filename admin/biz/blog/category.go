package blog

import (
	"context"
	"errors"
	"fmt"
	"github.com/bsm/redislock"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gookit/goutil/strutil"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/herhe-com/framework/microservice/locker"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/http/request/blog"
	res "github.com/tizips/uper-go/admin/http/response/blog"
	"github.com/tizips/uper-go/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func ToCategories(ctx context.Context, c *app.RequestContext) {

	var categories []model.BlgCategory

	facades.Gorm.Order("`order` asc, `created_at` asc").Find(&categories)

	responses := make([]res.ToCategories, 0)

	for _, item := range categories {
		if item.ParentID == nil || strutil.IsEmpty(*item.ParentID) {
			responses = append(responses, res.ToCategories{
				ID:       item.ID,
				Name:     item.Name,
				Order:    item.Order,
				Type:     item.Type,
				IsEnable: item.IsEnable,
				Children: make([]res.ToCategories, 0),
			})
		}
	}

	for index, item := range responses {

		for _, value := range categories {

			if value.ParentID != nil && item.ID == *value.ParentID {
				responses[index].Children = append(responses[index].Children, res.ToCategories{
					ID:       value.ID,
					Name:     value.Name,
					Order:    value.Order,
					Type:     value.Type,
					IsEnable: value.IsEnable,
				})
			}
		}
	}

	http.Success(c, responses)
}

func DoCategoryByCreate(ctx context.Context, c *app.RequestContext) {

	var request blog.DoCategoryByCreate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if !strutil.IsEmpty(request.Parent) {

		var parent model.BlgCategory

		fp := facades.Gorm.First(&parent, "`id`=?", request.Parent)

		if errors.Is(fp.Error, gorm.ErrRecordNotFound) {
			http.NotFound(c, "未找到该父级栏目")
			return
		} else if fp.Error != nil {
			http.Fail(c, "父级栏目查询失败：%v", fp.Error)
			return
		}

		if parent.Type != model.BlogCategoryForTypeOfParent {
			http.Fail(c, "该栏目下无法添加子栏目")
			return
		}
	}

	category := model.BlgCategory{
		ID:        facades.Snowflake.Generate().String(),
		Type:      request.Type,
		Name:      request.Name,
		Picture:   request.Picture,
		Order:     request.Order.Order,
		IsComment: request.IsComment,
		IsEnable:  request.IsEnable,
	}

	if !strutil.IsEmpty(request.Parent) {
		category.ParentID = &request.Parent
	}

	tx := facades.Gorm.Begin()

	if cc := tx.Create(&category); cc.Error != nil {
		tx.Rollback()
		http.Fail(c, "写入失败：%v", cc.Error)
		return
	}

	if request.Type != model.BlogCategoryForTypeOfParent {

		seo := model.BlgSEO{
			Type:        constants.BlogTypeOfCategory,
			OtherID:     category.ID,
			Title:       request.Title,
			Keyword:     request.Keyword,
			Description: request.Description,
		}

		if cs := tx.Create(&seo); cs.Error != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", cs.Error)
			return
		}
	}

	if request.Type == model.BlogCategoryForTypeOfPage {

		html := model.BlgHTML{
			Type:    constants.BlogTypeOfCategory,
			OtherID: category.ID,
			Content: request.Content,
		}

		if ct := tx.Create(&html); ct.Error != nil {
			tx.Rollback()
			http.Fail(c, "写入失败：%v", ct.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoCategoryByUpdate(ctx context.Context, c *app.RequestContext) {

	var request blog.DoCategoryByUpdate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	obtain, err := facades.Locker.Obtain(ctx, locker.Keys(request.ID), time.Second*60, &redislock.Options{RetryStrategy: redislock.LinearBackoff(100 * time.Millisecond)})

	if err != nil {
		http.Fail(c, "处理失败：%v", err)
		return
	}

	defer obtain.Release(ctx)

	var category model.BlgCategory

	fc := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfCategory) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfCategory) }).
		First(&category, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该栏目")
		return
	} else if fc.Error != nil {
		http.Fail(c, "栏目查询失败：%v", fc.Error)
		return
	}

	tx := facades.Gorm.Begin()

	category.Name = request.Name
	category.Picture = request.Picture
	category.Order = request.Order.Order
	category.IsEnable = request.IsEnable
	category.IsComment = request.IsComment

	if uc := tx.Omit(clause.Associations).Save(&category); uc.Error != nil {
		tx.Rollback()
		http.Fail(c, "修改失败：%v", uc.Error)
		return
	}

	if category.SEO != nil {

		category.SEO.Title = request.Title
		category.SEO.Keyword = request.Keyword
		category.SEO.Description = request.Description

		if us := tx.Save(&category.SEO); us.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", us.Error)
			return
		}
	}

	if category.HTML != nil {

		category.HTML.Content = request.Content

		if ut := tx.Save(&category.HTML); ut.Error != nil {
			tx.Rollback()
			http.Fail(c, "修改失败：%v", ut.Error)
			return
		}
	}

	tx.Commit()

	http.Success[any](c)
}

func DoCategoryByDelete(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var category model.BlgCategory

	fc := facades.Gorm.First(&category, "`id`=?", id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该栏目")
		return
	} else if fc.Error != nil {
		http.Fail(c, fmt.Sprintf("栏目查询失败：%v", fc.Error))
		return
	}

	if dc := facades.Gorm.Delete(&category); dc.Error != nil {
		http.Fail(c, "删除失败：%v", dc.Error)
		return
	}

	http.Success[any](c)
}

func DoCategoryByEnable(ctx context.Context, c *app.RequestContext) {

	var request blog.DoCategoryByEnable

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var category model.BlgCategory

	fc := facades.Gorm.First(&category, "`id`=?", request.ID)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该栏目")
		return
	} else if fc.Error != nil {
		http.Fail(c, "栏目查询失败：%v", fc.Error)
		return
	}

	if ec := facades.Gorm.Model(&category).Update("is_enable", request.IsEnable); ec.Error != nil {
		http.Fail(c, "操作失败：%v", ec.Error)
		return
	}

	http.Success[any](c)
}

func ToCategoryByInformation(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var category model.BlgCategory

	fc := facades.Gorm.
		Preload("SEO", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfCategory) }).
		Preload("HTML", func(t *gorm.DB) *gorm.DB { return t.Where("`type`=?", constants.BlogTypeOfCategory) }).
		First(&category, "`id`=?", id)

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该栏目")
		return
	} else if fc.Error != nil {
		http.Fail(c, "栏目查询失败：%v", fc.Error)
		return
	}

	responses := res.ToCategoryByInformation{
		ID:        category.ID,
		Name:      category.Name,
		Type:      category.Type,
		Picture:   category.Picture,
		Order:     category.Order,
		IsComment: category.IsComment,
		IsEnable:  category.IsEnable,
	}

	if category.ParentID != nil {
		responses.Parent = *category.ParentID
	}

	if category.SEO != nil {
		responses.Title = category.SEO.Title
		responses.Keyword = category.SEO.Keyword
		responses.Description = category.SEO.Description
	}

	if category.HTML != nil {
		responses.Content = category.HTML.Content
	}

	http.Success(c, responses)
}

func ToCategoryByParent(ctx context.Context, c *app.RequestContext) {

	var categories []model.BlgCategory

	facades.Gorm.Order("`order` asc, `created_at` asc").Find(&categories, "`type`=?", model.BlogCategoryForTypeOfParent)

	responses := make([]res.ToCategoryByParent, len(categories))

	for index, item := range categories {
		responses[index] = res.ToCategoryByParent{
			ID:   item.ID,
			Name: item.Name,
		}
	}

	http.Success(c, responses)
}

func ToCategoryByOpening(ctx context.Context, c *app.RequestContext) {

	var categories []model.BlgCategory

	facades.Gorm.Order("`order` asc, `created_at` asc").Find(&categories, "`type` in (?)", []string{model.BlogCategoryForTypeOfParent, model.BlogCategoryForTypeOfList})

	parent := make([]res.ToCategoryByOpening, 0)

	for _, item := range categories {
		if item.ParentID == nil || strutil.IsEmpty(*item.ParentID) {
			parent = append(parent, res.ToCategoryByOpening{
				ID:       item.ID,
				Type:     item.Type,
				Name:     item.Name,
				Children: make([]res.ToCategoryByOpening, 0),
			})
		}
	}

	for index, item := range parent {

		for _, value := range categories {

			if value.ParentID != nil && item.ID == *value.ParentID {
				parent[index].Children = append(parent[index].Children, res.ToCategoryByOpening{
					ID:   value.ID,
					Name: value.Name,
				})
			}
		}
	}

	responses := make([]res.ToCategoryByOpening, 0)

	for _, item := range parent {
		if item.Type == model.BlogCategoryForTypeOfList || len(item.Children) > 0 {
			responses = append(responses, res.ToCategoryByOpening{
				ID:       item.ID,
				Name:     item.Name,
				Children: item.Children,
			})
		}
	}

	http.Success(c, responses)
}
