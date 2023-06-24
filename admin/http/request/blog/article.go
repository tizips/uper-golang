package blog

import "github.com/herhe-com/framework/contracts/http/request"

type ToArticleByPaginate struct {
	request.Paginate
}

type DoArticleByCreate struct {
	Category    string `json:"category" form:"category" binding:"required,snowflake" label:"栏目"`
	Name        string `json:"name" form:"name" binding:"required,max=120" label:"名称"`
	Source      string `json:"source" form:"source" binding:"omitempty,max=32" label:"转载"`
	URL         string `json:"url" form:"url" binding:"omitempty,required_with=Source,max=120,url" label:"转载链接"`
	Picture     string `json:"picture" form:"picture" binding:"omitempty,max=255,url" label:"封面"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"描述"`
	IsComment   uint8  `json:"is_comment" form:"is_comment" binding:"required,oneof=1 2" label:"开启评论"`
	Content     string `json:"content" form:"content" binding:"required" label:"内容"`
	Text        string `json:"text" form:"text" binding:"required" label:"文本内容"`

	request.Enable
}

type DoArticleByUpdate struct {
	Category    string `json:"category" form:"category" binding:"required,snowflake" label:"栏目"`
	Name        string `json:"name" form:"name" binding:"required,max=120" label:"名称"`
	Source      string `json:"source" form:"source" binding:"omitempty,max=32" label:"转载"`
	URL         string `json:"url" form:"url" binding:"omitempty,required_with=Source,max=120,url" label:"转载链接"`
	Picture     string `json:"picture" form:"picture" binding:"omitempty,max=255,url" label:"封面"`
	Title       string `json:"title" form:"title" binding:"omitempty,max=255" label:"标题"`
	Keyword     string `json:"keyword" form:"keyword" binding:"omitempty,max=255" label:"词组"`
	Description string `json:"description" form:"description" binding:"omitempty,max=255" label:"描述"`
	IsComment   uint8  `json:"is_comment" form:"is_comment" binding:"required,oneof=1 2" label:"开启评论"`
	Content     string `json:"content" form:"content" binding:"required" label:"内容"`
	Text        string `json:"text" form:"text" binding:"required" label:"文本内容"`

	request.IDOfSnowflake
	request.Enable
}

type DoArticleByEnable struct {
	request.IDOfSnowflake
	request.Enable
}
