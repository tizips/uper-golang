package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/helper/authorize"
)

func Permission(permission string) app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {

		if ok, _ := facades.Casbin.HasRoleForUser(auth.NameOfUser(authorize.ID(c)), auth.NameOfRoleWithDeveloper()); ok {
			c.Next(ctx)
			return
		}

		permissions := []any{auth.NameOfUser(authorize.ID(c)), permission}

		if ok, _ := facades.Casbin.Enforce(permissions...); !ok {
			c.Abort()
			http.Forbidden(c)
			return
		}

		c.Next(ctx)
	}
}
