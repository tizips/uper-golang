package blog

type ToLinkByPaginate struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	Logo      string `json:"logo"`
	Email     string `json:"email"`
	Position  string `json:"position"`
	Order     uint8  `json:"order"`
	IsEnable  uint8  `json:"is_enable"`
	CreatedAt string `json:"created_at"`
}
