package blog

import "github.com/herhe-com/framework/contracts/http/request"

type ToArticles struct {
	Categories []string `json:"categories" form:"categories" query:"categories" validate:"omitempty,unique,dive,required,snowflake" label:"栏目"`

	request.Paginate
}
