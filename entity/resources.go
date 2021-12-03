package entity

import "gorm.io/gorm"

type Resources struct {
	gorm.Model
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	ToolsData []Tools `json:"tools_data"`
}
