package dto

type CreateTools struct {
	Title       string `json:"title" form:"title" binding:"required,min=3,max=10" `
	Image       string `json:"image"  form:"image" binding:"required,min=3,max=10" `
	SubTitle    string `json:"sub_title"  form:"sub_title" binding:"required,min=3,max=10" `
	Url         string `json:"url" form:"url"  `
	ResourcesId uint   `json:"resources_id" form:"resources_id" binding:"required" `
}

type UpdateTools struct {
	Id          uint   `json:"id" form:"id" binding:"required" validate:"required,"`
	Title       string `json:"title" form:"name" binding:"required" validate:"required,min=3,max=50"`
	Image       string `json:"image"  form:"image" binding:"required" validate:"required,"`
	SubTitle    string `json:"sub_title"  form:"sub_title" binding:"required" validate:"required,"`
	Url         string `json:"url" form:"url" binding:"required" validate:"required,"`
	ResourcesId uint   `json:"resources_id" form:"resources_id" binding:"required" validate:"required,"`
}
