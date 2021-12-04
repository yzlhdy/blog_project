package main

import (
	"blog_project/config"
	"blog_project/controller"
	"blog_project/repository"
	"blog_project/service"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                       = config.SetUpDatabaseConnection()
	resourRepository repository.ResourcesRepository = repository.NewResourceRepository(db)
	resourService    service.ResourcesService       = service.NewResourcesService(resourRepository)
	resourController controller.ResourcesController = controller.NewResourceController(resourService)
)

func main() {
	defer config.CloseDatabase(db)
	server := gin.Default()

	resources := server.Group("/api/v1")
	{
		resources.POST("/resources", resourController.InsertResource)
		resources.POST("/resources/update", resourController.UpdateResource)
	}

	server.Run(":8848")
}
