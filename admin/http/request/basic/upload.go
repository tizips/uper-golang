package basic

type DoUploadByFile struct {
	Dir string `json:"dir" form:"dir" validate:"required,max=100,dirs" label:"目录"`
}
