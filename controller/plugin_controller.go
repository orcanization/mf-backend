package controller

import (
	"orca-backend/repository"

	"github.com/gin-gonic/gin"
)

type PluginController struct {
	Repo repository.PluginRepository
}

func (controller *PluginController) FindAll(ctx *gin.Context) {
	plugins := controller.Repo.FindAll()
	ctx.JSON(200, plugins)
}
