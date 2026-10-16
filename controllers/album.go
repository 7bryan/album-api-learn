package controllers

import (
	"net/http"

	"github.com/7bryan/album-api-learn/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AlbumController struct {
	DB *gorm.DB
}

func (ctrl *AlbumController) FindAlbums(c *gin.Context) {
	var albums []models.Album
	if err := ctrl.DB.Find(&albums).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)
}

func (ctrl *AlbumController) CreateAlbum(c *gin.Context) {
	var input models.Album

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}
