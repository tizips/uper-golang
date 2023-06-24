package basic

type ToAccountByPermissions struct {
	Module string `json:"module" query:"module" form:"module" validate:"required" label:"模块"`
}
