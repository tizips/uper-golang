package blog

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/contracts/http/response"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	req "github.com/tizips/uper-go/admin/http/request/blog"
	res "github.com/tizips/uper-go/admin/http/response/blog"
	"github.com/tizips/uper-go/model"
	"gorm.io/gorm"
)

func ToLinkByPaginate(ctx context.Context, c *app.RequestContext) {

	var request req.ToLinkByPaginate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	responses := response.Paginate[res.ToLinkByPaginate]{
		Page: request.GetPage(),
		Size: request.GetSize(),
	}

	tx := facades.Gorm

	tx.Model(&model.BlgLink{}).Count(&responses.Total)

	if responses.Total > 0 {

		var links []model.BlgLink

		tx.
			Order("`id` desc").
			Limit(request.GetLimit()).
			Offset(request.GetOffset()).
			Find(&links)

		responses.Data = make([]res.ToLinkByPaginate, len(links))

		for index, item := range links {
			responses.Data[index] = res.ToLinkByPaginate{
				ID:        item.ID,
				Name:      item.Name,
				URL:       item.URL,
				Logo:      item.Logo,
				Email:     item.Email,
				Position:  item.Position,
				Order:     item.Order,
				IsEnable:  item.IsEnable,
				CreatedAt: item.CreatedAt.ToDateTimeString(),
			}
		}
	}

	http.Success(c, responses)
}

func DoLinkByCreate(ctx context.Context, c *app.RequestContext) {

	var request req.DoLinkByCreate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	link := model.BlgLink{
		Name:     request.Name,
		URL:      request.URL,
		Logo:     request.Logo,
		Email:    request.Email,
		Position: request.Position,
		Order:    request.Order.Order,
		IsEnable: request.IsEnable,
	}

	if cl := facades.Gorm.Create(&link); cl.Error != nil {
		http.Fail(c, "写入失败：%v", cl.Error)
		return
	}

	http.Success[any](c)
}

func DoLinkByUpdate(ctx context.Context, c *app.RequestContext) {

	var request req.DoLinkByUpdate

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var link model.BlgLink

	fl := facades.Gorm.First(&link, "`id`=?", request.ID)

	if errors.Is(fl.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fl.Error != nil {
		http.Fail(c, "数据查询失败：%v", fl.Error)
		return
	}

	link.Name = request.Name
	link.URL = request.URL
	link.Logo = request.Logo
	link.Email = request.Email
	link.Position = request.Position
	link.Order = request.Order.Order
	link.IsEnable = request.IsEnable

	if ul := facades.Gorm.Updates(&link); ul.Error != nil {
		http.Fail(c, "修改失败：%v", ul.Error)
		return
	}

	http.Success[any](c)
}

func DoLinkByEnable(ctx context.Context, c *app.RequestContext) {

	var request req.DoLinkByEnable

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var link model.BlgLink

	fl := facades.Gorm.First(&link, "`id`=?", request.ID)

	if errors.Is(fl.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fl.Error != nil {
		http.Fail(c, "数据查询失败：%v", fl.Error)
		return
	}

	if ul := facades.Gorm.Model(&link).Update("is_enable", request.IsEnable); ul.Error != nil {
		http.Fail(c, "启禁失败：%v", ul.Error)
		return
	}

	http.Success[any](c)
}

func DoLinkByDelete(ctx context.Context, c *app.RequestContext) {

	id := c.Param("id")

	var link model.BlgLink

	fl := facades.Gorm.First(&link, "`id`=?", id)

	if errors.Is(fl.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未找到该数据")
		return
	} else if fl.Error != nil {
		http.Fail(c, "数据查询失败：%v", fl.Error)
		return
	}

	if ul := facades.Gorm.Delete(&link, "`id`=?", id); ul.Error != nil {
		http.Fail(c, "启禁失败：%v", ul.Error)
		return
	}

	http.Success[any](c)
}
