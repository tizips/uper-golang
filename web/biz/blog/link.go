package blog

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/model"
	req "github.com/tizips/uper-go/web/http/request/blog"
	res "github.com/tizips/uper-go/web/http/response/blog"
)

func ToLinks(ctx context.Context, c *app.RequestContext) {

	var request req.ToLinks

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var links []model.BlgLink

	tx := facades.Gorm

	if len(request.Positions) > 0 {
		tx = tx.Where("`position` IN (?)", request.Positions)
	}

	tx.
		Order("`id` desc").
		Find(&links)

	responses := make([]res.ToLinks, len(links))

	for index, item := range links {
		responses[index] = res.ToLinks{
			ID:        item.ID,
			Name:      item.Name,
			URL:       item.URL,
			Logo:      item.Logo,
			CreatedAt: item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(c, responses)
}
