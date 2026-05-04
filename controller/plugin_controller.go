package controller

import (
	"net/http"
	url2 "net/url"
	"orca-backend/models"
	"orca-backend/repository"

	"github.com/gin-gonic/gin"
)

type PluginController struct {
	Repo repository.PluginRepository
}

type CreatePluginRequest struct {
	Name string `json:"name" binding:"required,min=2,max=100"`
	Url  string `json:"url" binding:"required"`
}

func (controller *PluginController) Create(ctx *gin.Context) {
	var request CreatePluginRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := url2.ParseRequestURI(request.Url)
	isLocal := err != nil

	plugin := models.Plugin{
		Name:  request.Name,
		Url:   request.Url,
		Local: isLocal,
	}
	controller.Repo.Create(plugin)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Plugin installed successfully"})
}

func (controller *PluginController) FindAll(ctx *gin.Context) {
	plugins := controller.Repo.FindAll()
	ctx.JSON(200, plugins)
}
