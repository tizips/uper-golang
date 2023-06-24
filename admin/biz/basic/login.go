package basic

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/gookit/goutil/dump"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/contracts/util"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/helper/authorize"
	req "github.com/tizips/uper-go/admin/http/request/basic"
	res "github.com/tizips/uper-go/admin/http/response/basic"
	"github.com/tizips/uper-go/model"
	"gorm.io/gorm"
)

func DoLoginByAccount(ctx context.Context, c *app.RequestContext) {

	var request req.DoLoginByAccount

	if err := c.Bind(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	if err := facades.Validator.Struct(&request); err != nil {
		http.BadRequest(c, err)
		return
	}

	var user model.SysUser

	fu := facades.Gorm.First(&user, "`username`=? and `is_enable`=?", request.Username, util.EnableOfYes)
	if fu.Error != nil {
		http.Fail(c, "用户名或密码错误")
		return
	}

	if !auth.CheckPassword(request.Password, user.Password) {
		http.Fail(c, "用户名或密码错误")
		return
	}

	lifetime := facades.Cfg.GetInt("jwt.lifetime")

	token, err := auth.MakeJwtToken(constants.JwtOfIssuerWithAdmin, user.ID, lifetime)
	if err != nil {
		dump.P(err)
		http.Login(c)
		return
	}

	// 查询最高授权 平台 > 集团 > 商户 > 单店
	var bind model.SysUserBindRole

	fb := facades.Gorm.Order("`role_id` asc").First(&bind, "`user_id`=?", user.ID)
	if errors.Is(fb.Error, gorm.ErrRecordNotFound) {
		http.NotFound(c, "未查询到被授权的角色")
		return
	} else if fb.Error != nil {
		http.Fail(c, "登陆失败：%v", fb.Error)
		return
	}

	responses := res.DoLogin{
		Token:    token,
		Lifetime: lifetime,
	}

	http.Success(c, responses)
}

func DoLoginByOut(ctx context.Context, c *app.RequestContext) {

	if !authorize.BlacklistOfJwt(ctx, c) {
		http.Fail(c, "退出失败，请稍后重试")
		return
	}

	http.Success[any](c)
}
