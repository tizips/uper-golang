package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/model"
)

func ToSetting(ctx context.Context, c *app.RequestContext, module string) {

	var settings []model.ComSetting

	facades.Gorm.Order("`order` asc, `id` asc").Find(&settings, "`module`=?", module)

	responses := make(map[string]string, len(settings))

	for _, item := range settings {
		responses[item.Key] = item.Val
	}

	http.Success(c, responses)
}
