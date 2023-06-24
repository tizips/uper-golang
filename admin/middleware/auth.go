package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/http"
	"github.com/tizips/uper-go/admin/helper/authorize"
)

func Auth() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {

		if !authorize.Check(c) {
			c.Abort()
			http.Unauthorized(c)
			return
		}

		if authorize.CheckBlacklistOfJwt(ctx, c) {
			c.Abort()
			http.Unauthorized(c)
			return
		}

		c.Next(ctx)
	}
}
