package repository

import (
	"blog_project/entity"

	"gorm.io/gorm"
)

type ToolsRepository interface {
	InsertTools(tools entity.Tools) entity.Tools
	UpdateTools(tools entity.Tools) entity.Tools
	DeleteTools(id int) entity.Tools
	ProfileTools(id int) entity.Tools
	FindAll(page int, limit int) []entity.Tools
}

type toolsRepository struct {
	connection *gorm.DB
}

func NewToolsRepository(db *gorm.DB) ToolsRepository {
	return &toolsRepository{
		connection: db,
	}
}

func (t *toolsRepository) InsertTools(tools entity.Tools) entity.Tools {

	t.connection.Create(&tools)
	return tools
}

func (t *toolsRepository) UpdateTools(tools entity.Tools) entity.Tools {

	t.connection.Create(&tools)
	return tools
}

func (t *toolsRepository) DeleteTools(id int) entity.Tools {
	var tools entity.Tools
	t.connection.Where("id = ?", id).Delete(&tools)
	return tools
}

func (t *toolsRepository) ProfileTools(id int) entity.Tools {
	var tools entity.Tools
	t.connection.Where("id = ?", id).Find(&tools)
	return tools
}
func (t *toolsRepository) FindAll(page int, limit int) []entity.Tools {
	var tools []entity.Tools
	t.connection.Preload("Recources").Limit(limit).Offset((page - 1) * limit).Find(&tools)
	return tools
}
