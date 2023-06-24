package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/herhe-com/framework/auth"
	"github.com/herhe-com/framework/facades"
	"github.com/tizips/uper-go/admin/constants"
	"github.com/tizips/uper-go/admin/helper/authorize"
)

func Jwt() app.HandlerFunc {

	return func(ctx context.Context, c *app.RequestContext) {

		if token := c.GetHeader(constants.JwtOfAuthorization); len(token) > 0 {

			lifetime := facades.Cfg.GetInt("jwt.lifetime")

			claims, refresh, ok := auth.CheckJwtToken(ctx, string(token), constants.JwtOfIssuerWithAdmin, lifetime)

			if ok && claims != nil {
				c.Set(constants.ContextOfIdWithAdmin, claims.Subject)
				c.Set(constants.ContextOfClaimsWithAdmin, claims)
			}

			if refresh != "" {

				c.Header(constants.JwtOfAuthorization, refresh)

				if authorize.Check(c) {
					_ = auth.DoRoleOfRefresh(ctx, authorize.ID(c))
				}
			}
		}
	}
}
