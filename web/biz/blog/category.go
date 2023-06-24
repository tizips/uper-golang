package blog

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gookit/goutil/strutil"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/model"
	res "github.com/tizips/uper-go/web/http/response/blog"
	"gorm.io/gorm"
)

func ToCategories(ctx context.Context, c *app.RequestContext) {

	var categories []model.BlgCategory

	facades.Gorm.Order("`order` asc, `created_at` asc").Find(&categories, "is_enable=?", util.EnableOfYes)

	responses := make([]res.ToCategories, 0)

	for _, item := range categories {
		if item.ParentID == nil || strutil.IsEmpty(*item.ParentID) {
			responses = append(responses, res.ToCategories{
				ID:       item.ID,
				Name:     item.Name,
				Type:     item.Type,
				Children: make([]res.ToCategories, 0),
			})
		}
	}

	for index, item := range responses {

		for _, value := range categories {

			if value.ParentID != nil && item.ID == *value.ParentID {
				responses[index].Children = append(responses[index].Children, res.ToCategories{
					ID:   value.ID,
					Name: value.Name,
					Type: value.Type,
				})
			}
		}
	}

	http.Success(c, responses)
}

func ToCategory(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var category model.BlgCategory

	fc := facades.Gorm.
		Preload("SEO").
		First(&category, "`id`=? and `type` IN (?)", id, []string{model.BlogCategoryForTypeOfList, model.BlogCategoryForTypeOfPage})

	if errors.Is(fc.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fc.Error != nil {
		http.Fail(c, "数据查询失败")
		return
	}

	if category.Type == model.BlogCategoryForTypeOfPage {
		_ = facades.Gorm.Model(&category).Association("HTML").Find(&category.HTML)
	}

	responses := res.ToCategory{
		ID:          category.ID,
		Type:        category.Type,
		Name:        category.Name,
		Picture:     category.Picture,
		Title:       category.SEO.Title,
		Keyword:     category.SEO.Keyword,
		Description: category.SEO.Description,
		IsComment:   category.IsComment,
		IsEnable:    category.IsEnable,
		CreatedAt:   category.CreatedAt.ToDateTimeString(),
	}

	if category.HTML != nil {
		responses.HTML = category.HTML.Content
	}

	http.Success(c, responses)
}
