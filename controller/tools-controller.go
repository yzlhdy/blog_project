package controller

import (
	"blog_project/dto"
	"blog_project/helper"
	"blog_project/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ToolsController interface {
	InsertTools(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type toolsController struct {
	toolSer service.ToolService
}

func NewToolsController(toolSer service.ToolService) ToolsController {
	return &toolsController{toolSer: toolSer}
}

func (s *toolsController) InsertTools(ctx *gin.Context) {
	var toolsCreate dto.CreateTools
	if err := ctx.ShouldBindJSON(&toolsCreate); err != nil {
		res := helper.BuildErrorResponse(401, "参数错误", helper.EmptyObj{}, err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}
	res := s.toolSer.InsertTools(toolsCreate)
	response := helper.BuildResponse(200, "添加成功", res)
	ctx.JSON(http.StatusOK, response)

}

func (s *toolsController) FindAll(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.Query("page"))
	limit, _ := strconv.Atoi(ctx.Query("limit"))
	res := s.toolSer.FindAll(page, limit)
	response := helper.BuildResponse(200, "查询成功", res)
	ctx.JSON(http.StatusOK, response)
}
