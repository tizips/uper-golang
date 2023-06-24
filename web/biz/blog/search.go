package blog

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/meilisearch/meilisearch-go"
	blogService "github.com/tizips/uper-go/admin/service/blog"
	"github.com/tizips/uper-go/model"
	req "github.com/tizips/uper-go/web/http/request/blog"
	res "github.com/tizips/uper-go/web/http/response/blog"
	"gorm.io/gorm"
)

func ToSearch(ctx context.Context, c *app.RequestContext) {

	var request req.ToSearch

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	result, err := blogService.SearchIndexForArticle().Search(request.Keyword, &meilisearch.SearchRequest{
		Limit:  int64(request.GetSize()),
		Offset: int64(request.GetOffset()),
	})

	if err != nil || result == nil {
		http.Fail(c, "搜索失败")
		return
	}

	responses := response.Paginate[res.ToSearch]{
		Size:  request.GetPage(),
		Page:  request.GetPage(),
		Total: result.EstimatedTotalHits,
		Data:  make([]res.ToSearch, len(result.Hits)),
	}

	ids := make([]string, len(result.Hits))

	for index, item := range result.Hits {

		if value, ok := item.(map[string]any); ok {

			responses.Data[index].ID, _ = value["id"].(string)
			responses.Data[index].Name, _ = value["name"].(string)
			responses.Data[index].CreatedAt, _ = value["created_at"].(string)

			ids[index] = responses.Data[index].ID
		}
	}

	if len(ids) > 0 {

		var articles []model.BlgArticle

		facades.Gorm.
			Preload("Category", func(t *gorm.DB) *gorm.DB { return t.Unscoped() }).
			Find(&articles, "`id` IN (?)", ids)

		for index, item := range responses.Data {

			for _, value := range articles {

				if item.ID == value.ID {

					responses.Data[index].Summary = value.Summary

					if value.Category != nil {
						responses.Data[index].Category = value.Category.Name
					}
				}
			}
		}
	}

	http.Success(c, responses)
}
