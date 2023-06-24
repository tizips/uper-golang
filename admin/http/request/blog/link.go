package blog

import "github.com/herhe-com/framework/contracts/http/request"

type ToLinkByPaginate struct {
	request.Paginate
}

type DoLinkByCreate struct {
	Name     string `json:"name" form:"name" binding:"required,max=20" label:"名称"`
	URL      string `json:"url" form:"url" binding:"required,max=64,url" label:"链接"`
	Logo     string `json:"logo" form:"logo" binding:"omitempty,max=120,url" label:"LOGO"`
	Email    string `json:"email" form:"email" binding:"omitempty,max=64,email" label:"邮箱"`
	Position string `json:"position" form:"position" binding:"required,oneof=all bottom other" label:"位置"`

	request.Order
	request.Enable
}

type DoLinkByUpdate struct {
	Name     string `json:"name" form:"name" binding:"required,max=20" label:"名称"`
	URL      string `json:"url" form:"url" binding:"required,max=64,url" label:"链接"`
	Logo     string `json:"logo" form:"logo" binding:"omitempty,max=120,url" label:"LOGO"`
	Email    string `json:"email" form:"email" binding:"omitempty,max=64,email" label:"邮箱"`
	Position string `json:"position" form:"position" binding:"required,oneof=all bottom other" label:"位置"`

	request.IDOfUint
	request.Order
	request.Enable
}

type DoLinkByEnable struct {
	request.IDOfUint
	request.Enable
}
