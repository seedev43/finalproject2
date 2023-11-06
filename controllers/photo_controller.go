package controllers

import (
	"fp2/database"
	"fp2/dto"
	"fp2/models"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(ctx *gin.Context) {
	var photo models.Photo
	var request dto.PostPhoto

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	_, err := govalidator.ValidateStruct(request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	photo = models.Photo{
		Title:    request.Title,
		Caption:  request.Caption,
		PhotoUrl: request.PhotoUrl,
		UserId:   userId,
	}

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
