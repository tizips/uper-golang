package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gookit/goutil/strutil"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/http/response/common"
	"github.com/tizips/uper-go/model"
)

func ToSetting(ctx context.Context, c *app.RequestContext, module string) {

	var settings []model.ComSetting

	facades.Gorm.Order("`order` asc, `id` asc").Find(&settings, "`module`=?", module)

	responses := make([]common.ToSetting, len(settings))

	for index, item := range settings {
		responses[index] = common.ToSetting{
			ID:         item.ID,
			Type:       item.Type,
			Label:      item.Label,
			Key:        item.Key,
			Val:        item.Val,
			IsRequired: item.IsRequired,
			CreatedAt:  item.CreatedAt.ToDateTimeString(),
		}
	}

	http.Success(c, responses)
}

func DoSetting(ctx context.Context, c *app.RequestContext, module string) {

	var request map[string]string

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var settings []model.ComSetting

	facades.Gorm.Find(&settings, "`module`=?", module)

	updates := make(map[uint]string, 0)

	for _, item := range settings {

		req, ok := request[item.Key]

		if item.IsRequired == model.ComSettingForIsRequiredOfYes && (!ok || strutil.IsEmpty(req)) {
			http.BadRequest(c, item.Label+"不能为空")
			return
		}

		if !strutil.IsEmpty(req) {

			var err error

			if item.Type == model.ComSettingForTypeOfEnable {
				err = facades.Validator.Var(req, "oneof=1 2")
			} else if item.Type == model.ComSettingForTypeOfURL || item.Type == model.ComSettingForTypeOfPicture {
				err = facades.Validator.Var(req, "url")
			} else if item.Type == model.ComSettingForTypeOfEmail {
				err = facades.Validator.Var(req, "email")
			}

			if err != nil {
				http.BadRequest(c, err)
				return
			}
		}

		if req != item.Val {
			updates[item.ID] = req
		}
	}

	if len(updates) > 0 {

		tx := facades.Gorm.Begin()

		for index, item := range updates {
			if us := tx.Model(model.ComSetting{}).Where("`id`=?", index).Update("val", item); us.Error != nil {
				tx.Rollback()
				http.Fail(c, "保存失败：%v", us.Error)
				return
			}
		}

		tx.Commit()
	}

	http.Success[any](c)
}
