package blog

type ToSearch struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"created_at"`
}
