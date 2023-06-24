package authorize

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-module/carbon/v2"
	"github.com/herhe-com/framework/auth"
	"time"
)

func CheckBlacklistOfJwt(ctx context.Context, c *app.RequestContext) bool {

	if Claims(c) == nil {
		return false
	}

	return auth.CheckBlacklist(ctx, BlacklistOfKeyWithJwt(Claims(c).ID)...)
}

func BlacklistOfJwt(ctx context.Context, c *app.RequestContext) bool {

	if Claims(c) == nil {
		return false
	}

	now := carbon.Now()

	expires := time.Duration(Claims(c).ExpiresAt.Unix()+12*60*60-now.Timestamp()) * time.Second

	return auth.Blacklist(ctx, now.Timestamp(), expires, BlacklistOfKeyWithJwt(Claims(c).ID)...)
}

func BlacklistOfKeyWithJwt(id string) []any {
	return []any{"jwt", id}
}
