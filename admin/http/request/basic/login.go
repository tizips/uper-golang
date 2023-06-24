package basic

type DoLoginByAccount struct {
	Username string `json:"username" form:"username" validate:"required,username" label:"用户名"`
	Password string `json:"password" form:"password" validate:"required,password" label:"密码"`
}
