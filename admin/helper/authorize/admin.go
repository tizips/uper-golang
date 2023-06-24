package authorize

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v4"
	"github.com/herhe-com/framework/cache"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/model"
)

func Check(c *app.RequestContext) bool {

	if ID(c) != "" {
		return true
	}

	return false
}

func ID(c *app.RequestContext) string {

	if value, ok := c.Get(constants.ContextOfIdWithAdmin); ok {
		return value.(string)
	}

	return ""
}

func User(ctx context.Context, c *app.RequestContext) (user model.SysUser) {

	if Check(c) {

		if value, ok := c.Get(constants.ContextOfUserWithAdmin); ok {
			return value.(model.SysUser)
		} else {

			if err := cache.FindById(ctx, &user, ID(c)); err == nil {
				c.Set(constants.ContextOfUserWithAdmin, user)
			}
		}
	}

	return user
}

func Claims(c *app.RequestContext) (claims *jwt.RegisteredClaims) {

	if value, ok := c.Get(constants.ContextOfClaimsWithAdmin); ok {
		return value.(*jwt.RegisteredClaims)
	}

	return nil
}
