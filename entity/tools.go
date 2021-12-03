package entity

import "gorm.io/gorm"

type Tools struct {
	gorm.Model
	Title       string `json:"title"`
	Image       string `json:"image"`
	SubTitle    string `json:"sub_title"`
	Url         string `json:"url"`
	ResourcesId uint   `json:"resources_id"`
}
