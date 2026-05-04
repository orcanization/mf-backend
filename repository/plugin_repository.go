package repository

import (
	"orca-backend/models"

	"gorm.io/gorm"
)

type PluginRepository interface {
	Create(plugin models.Plugin)
	FindAll() []models.Plugin
}

type PluginRepositoryImpl struct {
	DB *gorm.DB
}

func (p *PluginRepositoryImpl) Create(plugin models.Plugin) {
	p.DB.Create(&plugin)
}

func (p *PluginRepositoryImpl) FindAll() []models.Plugin {
	var plugins []models.Plugin
	p.DB.Find(&plugins)
	return plugins
}
