package blog

type ToArticles struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}

type ToArticle[T any] struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Source      string `json:"source"`
	URL         string `json:"url,omitempty"`
	Title       string `json:"title"`
	Keyword     string `json:"keyword"`
	Description string `json:"description"`
	Picture     string `json:"picture"`
	HTML        string `json:"html"`
	IsComment   uint8  `json:"is_comment"`
	IsEnable    uint8  `json:"is_enable"`
	CreatedAt   string `json:"created_at"`
}
