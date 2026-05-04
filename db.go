package main

import (
	"orca-backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() *gorm.DB {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	err = DB.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.Project{}, &models.ProjectAssignment{}, &models.Plugin{})
	if err != nil {
		panic("Failed to migrate database")
	}

	return DB
}
