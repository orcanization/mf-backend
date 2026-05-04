package controller

import (
	"net/http"
	"orca-backend/models"
	"orca-backend/repository"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProjectController struct {
	Repo repository.ProjectRepository
}

type CreateProjectRequest struct {
	Name        string `json:"name" binding:"required,min=2,max=100"`
	Description string `json:"description" binding:"max=255"`
}

func (controller *ProjectController) Create(ctx *gin.Context) {
	var request CreateProjectRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		Name:        request.Name,
		Description: request.Description,
	}
	controller.Repo.Save(project)
	ctx.JSON(http.StatusCreated, gin.H{"message": "Project created successfully"})
}

func (controller *ProjectController) FindById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	project, err := controller.Repo.FindById(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	ctx.JSON(http.StatusOK, project)
}
