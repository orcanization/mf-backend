package repository

import (
	"orca-backend/models"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	Save(project models.Project)
	Update(project models.Project)
	Delete(projectId uint)
	FindById(projectId uint) (models.Project, error)
	FindAll() []models.Project
}

type ProjectRepositoryImpl struct {
	Db *gorm.DB
}

func (p *ProjectRepositoryImpl) Save(project models.Project) {
	p.Db.Create(&project)
}

func (p *ProjectRepositoryImpl) Update(project models.Project) {
	p.Db.Model(&project).Updates(project)
}

func (p *ProjectRepositoryImpl) Delete(projectId uint) {
	p.Db.Delete(&models.Project{}, projectId)
}

func (p *ProjectRepositoryImpl) FindById(projectId uint) (models.Project, error) {
	var project models.Project
	result := p.Db.First(&project, projectId)
	return project, result.Error
}

func (p *ProjectRepositoryImpl) FindAll() []models.Project {
	var projects []models.Project
	p.Db.Find(projects)
	return projects
}
