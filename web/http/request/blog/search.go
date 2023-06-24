package blog

import "github.com/herhe-com/framework/contracts/http/request"

type ToSearch struct {
	Keyword string `json:"keyword" form:"keyword" query:"keyword" validate:"required,max=20" label:"关键词"`

	request.Paginate
}
