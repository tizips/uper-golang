package blog

type ToCategories struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Order    uint8          `json:"order"`
	Type     string         `json:"type"`
	IsEnable uint8          `json:"is_enable"`
	Children []ToCategories `json:"children,omitempty"`
}

type ToCategoryByParent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ToCategoryByInformation struct {
	ID          string `json:"id"`
	Parent      string `json:"parent,omitempty"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Title       string `json:"title,omitempty"`
	Keyword     string `json:"keyword,omitempty"`
	Description string `json:"description,omitempty"`
	Picture     string `json:"picture,omitempty"`
	Content     string `json:"content,omitempty"`
	Order       uint8  `json:"order"`
	IsComment   uint8  `json:"is_comment,omitempty"`
	IsEnable    uint8  `json:"is_enable"`
}

type ToCategoryByOpening struct {
	ID       string                `json:"id"`
	Type     string                `json:"type,omitempty"`
	Name     string                `json:"name"`
	Children []ToCategoryByOpening `json:"children,omitempty"`
}
