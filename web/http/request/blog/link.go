package blog

type ToLinks struct {
	Positions []string `json:"positions" form:"positions" query:"positions" validate:"omitempty,unique,dive,required,oneof=all bottom other"`
}
