package service

import (
	"blog_project/dto"
	"blog_project/entity"
	"blog_project/repository"

	"github.com/mashingan/smapping"
)

type ToolService interface {
	InsertTools(tools dto.CreateTools) entity.Tools
	UpdateTools(tools entity.Tools) entity.Tools
	DeleteTools(id int) entity.Tools
	ProfileTools(id int) entity.Tools
	FindAll(page int, limit int) []entity.Tools
}

type toolService struct {
	toolRepo repository.ToolsRepository
}

func NewToolService(toolRep repository.ToolsRepository) ToolService {
	return &toolService{toolRepo: toolRep}
}
func (re *toolService) InsertTools(tools dto.CreateTools) entity.Tools {
	createTools := entity.Tools{}
	err := smapping.FillStruct(&createTools, smapping.MapFields(&tools))
	if err != nil {
		panic(err)
	}
	result := re.toolRepo.InsertTools(createTools)
	return result

}

func (re *toolService) UpdateTools(tools entity.Tools) entity.Tools {
	panic("not implemented") // TODO: Implement
}

func (re *toolService) DeleteTools(id int) entity.Tools {
	panic("not implemented") // TODO: Implement
}

func (re *toolService) ProfileTools(id int) entity.Tools {
	panic("not implemented") // TODO: Implement
}
func (re *toolService) FindAll(page int, limit int) []entity.Tools {
	result := re.toolRepo.FindAll(page, limit)
	return result
}
