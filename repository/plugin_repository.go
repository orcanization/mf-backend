package repository

import (
	"orca-backend/models"

	"gorm.io/gorm"
)

type PluginRepository interface {
	FindAll() []models.Plugin
}

type PluginRepositoryImpl struct {
	DB *gorm.DB
}

func (p *PluginRepositoryImpl) FindAll() []models.Plugin {
	var plugins []models.Plugin
	p.DB.Find(&plugins)
	return plugins
}
