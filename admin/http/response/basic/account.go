package basic

type ToAccountByInformation struct {
	Nickname string `json:"nickname"`
	Username string `json:"username,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Avatar   string `json:"avatar"`
}

type ToAccountByPlatform struct {
	Id     string `json:"id"`
	Code   uint16 `json:"code"`
	Name   string `json:"name"`
	Group  string `json:"group"`
	Clique string `json:"clique,omitempty"`
	Bak    string `json:"bak,omitempty"`
}

type ToAccountByModules struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
