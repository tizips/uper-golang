package blog

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/tizips/uper-go/web/biz/common"
)

func ToSetting(ctx context.Context, c *app.RequestContext) {

	common.ToSetting(ctx, c, "blog")
}
