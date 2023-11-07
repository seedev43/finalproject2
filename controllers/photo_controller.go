package controllers

import (
	"fp2/database"
	"fp2/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(ctx *gin.Context) {
	photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	photo.UserId = userId

	if err := database.DB.Create(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":         photo.Id,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"user_id":    photo.UserId,
		"created_at": photo.CreatedAt,
	})
}

func GetPhotos(ctx *gin.Context) {
	photos := []models.Photo{}

	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}
