package dto

type CreateTools struct {
	Title       string `json:"title" form:"name" binding:"required" validate:"required,min=3,max=50"`
	Image       string `json:"image"  form:"image" binding:"required" validate:"required,"`
	SubTitle    string `json:"sub_title"  form:"sub_title" binding:"required" validate:"required,"`
	Url         string `json:"url" form:"url" binding:"required" validate:"required,"`
	ResourcesId uint   `json:"resources_id" form:"resources_id" binding:"required" validate:"required,"`
}
