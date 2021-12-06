package main

import (
	"blog_project/config"
	"blog_project/controller"
	"blog_project/middleware"
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
	// tools
	toolsRepository repository.ToolsRepository = repository.NewToolsRepository(db)
	toolsService    service.ToolService        = service.NewToolService(toolsRepository)
	toolsController controller.ToolsController = controller.NewToolsController(toolsService)
)

func main() {
	defer config.CloseDatabase(db)
	server := gin.New()
	server.Use(gin.Logger())
	server.Use(gin.Recovery())
	server.Use(middleware.Translations())

	// resources
	resources := server.Group("/api/v1")
	{
		resources.POST("/resources", resourController.InsertResource)
		resources.POST("/resources/update", resourController.UpdateResource)
	}
	// tools
	tools := server.Group("/api/v1")
	{
		tools.POST("/tools", toolsController.InsertTools)
		tools.GET("/tools", toolsController.FindAll)
	}

	server.Run(":8848")
}
