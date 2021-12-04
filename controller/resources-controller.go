package controller

import (
	"blog_project/dto"
	"blog_project/helper"
	"blog_project/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResourcesController interface {
	InsertResource(context *gin.Context)
	UpdateResource(context *gin.Context)
	DeleteResource(context *gin.Context)
	ProfileResource(context *gin.Context)
}

type resourcesController struct {
	resService service.ResourcesService
}

func NewResourceController(res service.ResourcesService) ResourcesController {
	return &resourcesController{
		resService: res,
	}
}
func (c *resourcesController) InsertResource(context *gin.Context) {
	var userCreate dto.CreateResourceRequest
	errDto := context.ShouldBindJSON(&userCreate)
	if errDto != nil {
		res := helper.BuildErrorResponse(401, "参数错误", helper.EmptyObj{}, errDto.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	u := c.resService.InsertResource(userCreate)
	response := helper.BuildResponse(200, "success", u)
	context.JSON(http.StatusOK, response)
}

func (c *resourcesController) UpdateResource(context *gin.Context) {
	var userCreate dto.UpdateResource
	errDto := context.ShouldBindJSON(&userCreate)
	if errDto != nil {
		res := helper.BuildErrorResponse(401, "参数错误", helper.EmptyObj{}, errDto.Error())
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	u := c.resService.UpdateResource(userCreate)
	response := helper.BuildResponse(200, "success", u)
	context.JSON(http.StatusOK, response)
}

func (c *resourcesController) DeleteResource(context *gin.Context) {
	// id := context.Param("id")
	// c.resService.DeleteResource(id)
}

func (c *resourcesController) ProfileResource(context *gin.Context) {
	panic("not implemented") // TODO: Implement
}
