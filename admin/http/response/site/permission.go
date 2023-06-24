package site

type ToPermissions struct {
	Code     string          `json:"code"`
	Name     string          `json:"name"`
	Children []ToPermissions `json:"children,omitempty"`
}
