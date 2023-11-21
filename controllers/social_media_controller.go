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

func CreateSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Input must be in JSON format",
		})
		return
	}

	socialMedia.UserId = userId

	if err := database.DB.Create(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	res := dto.SocialMediaCreateResponse{
		Id:             socialMedia.Id,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		CreatedAt:      socialMedia.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, res)

}

func GetSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := database.DB.Preload("User").First(&socialMedia, socialMediaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, socialMedia)
}

func GetSocialMedias(ctx *gin.Context) {
	socialMedias := []models.SocialMedia{}

	if err := database.DB.Preload("User").Find(&socialMedias).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"socialMedias": socialMedias})
}

func UpdateSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&socialMedia); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": "Input must be in JSON format",
		})
		return
	}

	if err := database.DB.Select("user_id").First(&socialMedia, socialMediaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	if socialMedia.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "You are not allowed to edit this social media data",
		})
		return
	}

	if err := database.DB.Model(&socialMedia).Where("id = ?", socialMediaId).Updates(socialMedia).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	res := dto.SocialMediaUpdateResponse{
		Id:             socialMediaId,
		Name:           socialMedia.Name,
		SocialMediaUrl: socialMedia.SocialMediaUrl,
		UserId:         socialMedia.UserId,
		UpdatedAt:      socialMedia.UpdatedAt,
	}
	ctx.JSON(http.StatusOK, res)
}

func DeleteSocialMedia(ctx *gin.Context) {
	socialMedia := models.SocialMedia{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := database.DB.Select("user_id").First(&socialMedia, socialMediaId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code":  http.StatusNotFound,
			"error": "Data not found",
		})
		return
	}

	if socialMedia.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":  http.StatusUnauthorized,
			"error": "You are not allowed to delete this social media data",
		})
		return
	}

	if err := database.DB.Where("id = ?", socialMediaId).Delete(&socialMedia).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":  http.StatusBadRequest,
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Your social media has been successfully deleted"})
}
