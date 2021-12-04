package repository

import (
	"blog_project/entity"

	"gorm.io/gorm"
)

type ResourcesRepository interface {
	InsertResource(reseau entity.Resources) entity.Resources
	UpdateResource(reseau entity.Resources) entity.Resources
	DeleteResource(id int) entity.Resources
	ProfileResource(id int) entity.Resources
}

type resourcesRepository struct {
	connection *gorm.DB
}

func NewResourceRepository(db *gorm.DB) ResourcesRepository {
	return &resourcesRepository{
		connection: db,
	}
}
func (db *resourcesRepository) InsertResource(res entity.Resources) entity.Resources {
	db.connection.Save(&res)
	return res
}

func (db *resourcesRepository) UpdateResource(reseau entity.Resources) entity.Resources {
	db.connection.Save(&reseau)
	return reseau
}

func (db *resourcesRepository) DeleteResource(id int) entity.Resources {
	var reseau entity.Resources
	db.connection.First(&reseau, id)
	db.connection.Delete(&reseau)
	return reseau
}

func (db *resourcesRepository) ProfileResource(id int) entity.Resources {
	var reseau entity.Resources
	db.connection.First(&reseau, id)
	return reseau
}
