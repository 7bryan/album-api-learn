package main

import (
	"log"

	"github.com/7bryan/album-api-learn/config"
	"github.com/7bryan/album-api-learn/controllers"
	"github.com/7bryan/album-api-learn/models"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()

	err := db.AutoMigrate(&models.Album{})
	if err != nil {
		log.Fatalf("Database Migration Failed: %v", err)
	}

	albumController := &controllers.AlbumController{DB: db}

	router := gin.Default()

	router.GET("/albums", albumController.FindAlbums)
	router.POST("/albums", albumController.CreateAlbum)

	router.Run(":8080")
}
