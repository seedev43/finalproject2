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

func CreateComment(ctx *gin.Context) {
	comment := models.Comment{}
	photo := models.Photo{}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	comment.UserId = userId

	if err := database.DB.Select("id").First(&photo, comment.PhotoId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Photo not found",
		})
		return
	}

	if err := database.DB.Create(&comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.CommentCreateResponse{
		Id:        comment.Id,
		Message:   comment.Message,
		PhotoId:   comment.PhotoId,
		UserId:    comment.UserId,
		CreatedAt: comment.CreatedAt,
	}
	ctx.JSON(http.StatusCreated, res)

}

func GetComment(ctx *gin.Context) {
	comment := models.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	if err := database.DB.Preload("User").Preload("Photo").First(&comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

func GetComments(ctx *gin.Context) {
	comments := []models.Comment{}

	if err := database.DB.Preload("User").Preload("Photo").Find(&comments).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comments)
}

func UpdateComment(ctx *gin.Context) {
	comment := models.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	if err := database.DB.Select("user_id").First(&comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	if comment.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to edit this comment",
		})
		return
	}

	if err := database.DB.Model(&comment).Where("id = ?", commentId).Updates(comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := dto.CommentUpdateResponse{
		Id:        commentId,
		Message:   comment.Message,
		UserId:    comment.UserId,
		UpdatedAt: comment.UpdatedAt,
	}

	ctx.JSON(http.StatusOK, res)
}

func DeleteComment(ctx *gin.Context) {
	comment := models.Comment{}
	commentId, _ := strconv.Atoi(ctx.Param("commentId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := database.DB.Select("user_id").First(&comment, commentId).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "Data not found",
		})
		return
	}

	if comment.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not allowed to delete this comment",
		})
		return
	}

	if err := database.DB.Where("id = ?", commentId).Delete(&comment).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Your comment has been successfully deleted"})
}
