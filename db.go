package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.AutoMigrate(&User{}, &Role{}, &Permission{})
}
