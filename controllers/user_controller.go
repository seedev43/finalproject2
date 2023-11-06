package controllers

import (
	"fp2/database"
	"fp2/dto"
	"fp2/helpers"
	"fp2/models"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func UserRegister(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	if err := database.DB.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"id":       user.Id,
		"email":    user.Email,
		"username": user.Username,
		"age":      user.Age,
	})
}

func UserLogin(ctx *gin.Context) {
	var user models.User
	var request dto.UserLogin

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

	password := request.Password

	if err := database.DB.Where("email = ?", request.Email).Take(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": "Email not registered",
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(user.Password), []byte(password))

	if !comparePass {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unathorized",
			"message": "Invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(user.Id, user.Email)
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func UserUpdate(ctx *gin.Context) {
	var user models.User
	var request dto.UserUpdate
	id, _ := strconv.Atoi(ctx.Param("userId"))

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Input must be in JSON format"})
		return
	}

	user.Id = userId

	if user.Id != uint(id) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorized",
			"message": "You are not authorized to edit this user.",
		})
		return
	}

	if err := database.DB.Model(&user).Where("id = ?", id).Updates(request).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"email":      user.Email,
		"username":   user.Username,
		"age":        user.Age,
		"updated_at": user.UpdatedAt,
	})

}

func UserDelete(ctx *gin.Context) {
	var user models.User

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := uint(userData["id"].(float64))

	user.Id = userId

	if err := database.DB.Delete(&user).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":   "Unauthorizated",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Your account has been succesfully deleted",
	})

}
