package blog

type ToArticleByPaginate struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}

type ToArticleByInformation[T any] struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    T      `json:"category"`
	Source      string `json:"source"`
	URL         string `json:"url,omitempty"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	Content     string `json:"content"`
	IsComment   uint8  `json:"is_comment"`
	IsEnable    uint8  `json:"is_enable"`
}
