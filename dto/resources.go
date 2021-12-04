package dto

import (
	"blog_project/entity"

	"gorm.io/gorm"
)

type CreateResourceRequest struct {
	Name  string `json:"name" form:"name" binding:"required" validate:"required,min=3,max=50"`
	Image string `json:"image" form:"image" binding:"required"`
	// ToolsData []entity.Tools `json:"tools_data" form:"tools_data" binding:"required"`
}
type UpdateResource struct {
	gorm.Model
	Name      string         `json:"name"`
	Image     string         `json:"image"`
	ToolsData []entity.Tools `json:"tools_data"`
}
