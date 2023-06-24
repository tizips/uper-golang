package blog

type ToCategories struct {
	ID       string         `json:"id"`
	Type     string         `json:"type,omitempty"`
	Name     string         `json:"name"`
	Children []ToCategories `json:"children,omitempty"`
}

type ToCategory struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Picture     string `json:"picture"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	HTML        string `json:"html,omitempty"`
	IsComment   uint8  `json:"is_comment"`
	IsEnable    uint8  `json:"is_enable"`
	CreatedAt   string `json:"created_at"`
}
