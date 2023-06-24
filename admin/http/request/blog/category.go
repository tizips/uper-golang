package blog

import "github.com/herhe-com/framework/contracts/http/request"

type DoCategoryByCreate struct {
	Parent      string `json:"parent" form:"parent" validate:"omitempty,snowflake" label:"父级"`
	Type        string `json:"type" form:"type" validate:"required,oneof=parent page list" label:"类型"`
	Name        string `json:"name" form:"name" validate:"required,max=120" label:"名称"`
	Picture     string `json:"picture" form:"picture" validate:"omitempty,max=255,url" label:"封面"`
	Title       string `json:"title" form:"title" validate:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" validate:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" validate:"omitempty,max=255" label:"描述"`
	IsComment   uint8  `json:"is_comment" form:"is_comment" validate:"required,oneof=1 2" label:"开启评论"`
	Content     string `json:"content" form:"content" validate:"required_if=Type page" label:"内容"`
	Text        string `json:"text" form:"text" validate:"required_if=Type page" label:"纯文本"`

	request.Order
	request.Enable
}

type DoCategoryByUpdate struct {
	request.IDOfSnowflake

	Name        string `json:"name" form:"name" validate:"required,max=120" label:"名称"`
	Picture     string `json:"picture" form:"picture" validate:"omitempty,max=255,url" label:"封面"`
	Title       string `json:"title" form:"title" validate:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" validate:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" validate:"omitempty,max=255" label:"描述"`
	IsComment   uint8  `json:"is_comment" form:"is_comment" validate:"required,oneof=1 2" label:"开启评论"`
	Content     string `json:"content" form:"content" validate:"required_if=Type page" label:"内容"`
	Text        string `json:"text" form:"text" validate:"required_if=Type page" label:"纯文本"`

	request.Order
	request.Enable
}

type DoCategoryByEnable struct {
	request.IDOfSnowflake
	request.Enable
}
