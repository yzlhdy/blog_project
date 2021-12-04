package service

import (
	"blog_project/dto"
	"blog_project/entity"
	"blog_project/repository"
	"log"

	"github.com/mashingan/smapping"
)

type ResourcesService interface {
	InsertResource(reseau dto.CreateResourceRequest) entity.Resources
	UpdateResource(reseau dto.UpdateResource) entity.Resources
	DeleteResource(id int) entity.Resources
	ProfileResource(id int) entity.Resources
}

type resourcesService struct {
	resRepository repository.ResourcesRepository
}

func NewResourcesService(resRep repository.ResourcesRepository) ResourcesService {
	return &resourcesService{
		resRepository: resRep,
	}
}
func (service *resourcesService) InsertResource(reseau dto.CreateResourceRequest) entity.Resources {
	resDtoCreate := entity.Resources{}
	err := smapping.FillStruct(&resDtoCreate, smapping.MapFields(&reseau))
	if err != nil {
		log.Fatalf("failed to fill %v", err)
	}
	res := service.resRepository.InsertResource(resDtoCreate)
	return res
}

func (service *resourcesService) UpdateResource(reseau dto.UpdateResource) entity.Resources {

	resDtoUpdate := entity.Resources{}
	err := smapping.FillStruct(&resDtoUpdate, smapping.MapFields(&reseau))
	if err != nil {
		log.Fatalf("failed to fill %v", err)
	}
	res := service.resRepository.InsertResource(resDtoUpdate)
	return res
}

func (service *resourcesService) DeleteResource(id int) entity.Resources {
	return service.resRepository.DeleteResource(id)
}

func (service *resourcesService) ProfileResource(id int) entity.Resources {
	return service.resRepository.ProfileResource(id)
}
