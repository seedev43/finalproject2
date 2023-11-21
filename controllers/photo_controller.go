package controllers

import (
	"fp2/database"
	"fp2/dto"
	"fp2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreatePhoto(ctx *gin.Context) {
	photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Input must be in JSON format",
		})
		return
	}

	photo.UserId = userId

	if err := database.DB.Create(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	res := dto.PhotoCreateResponse{
		Id:        photo.Id,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		CreatedAt: photo.CreatedAt,
	}

	ctx.JSON(http.StatusCreated, res)
}

func GetPhoto(ctx *gin.Context) {
	photo := models.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	if err := database.DB.Preload("User").First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, photo)
}

func GetPhotos(ctx *gin.Context) {
	photos := []dto.PhotosResponse{}

	if err := database.DB.Preload("User").Find(&photos).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, photos)
}

func UpdatePhoto(ctx *gin.Context) {
	photo := models.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&photo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Input must be in JSON format",
		})
		return
	}

	if err := database.DB.Select("user_id").First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	if photo.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "You are not allowed to edit this photo data",
		})
		return
	}

	if err := database.DB.Model(&photo).Where("id = ?", photoId).Updates(photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	res := dto.PhotoUpdateResponse{
		Id:        photoId,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoUrl:  photo.PhotoUrl,
		UserId:    photo.UserId,
		UpdatedAt: photo.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

func DeletePhoto(ctx *gin.Context) {
	photo := models.Photo{}
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := database.DB.Select("user_id").First(&photo, photoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	if photo.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "You are not allowed to delete this photo data",
		})
		return
	}

	if err := database.DB.Where("id = ?", photoId).Delete(&photo).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Your photo has been successfully deleted"})
}
