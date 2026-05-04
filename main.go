package main

import (
	"net/http"
	"orca-backend/controller"
	"orca-backend/repository"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	db := initDatabase()
	router := buildRouter(db)

	server := &http.Server{
		Addr:           ":8888",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func buildRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(corsConfig))

	pluginRepo := &repository.PluginRepositoryImpl{DB: db}
	pluginController := &controller.PluginController{Repo: pluginRepo}

	router.Static("/plugins", "./plugins")

	router.GET("", func(context *gin.Context) { context.JSON(http.StatusOK, "Hello from Orca") })
	router.GET("/plugins", pluginController.FindAll)
	router.POST("/plugins", pluginController.Create)

	return router
}
