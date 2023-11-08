package routers

import (
	"fp2/controllers"
	"fp2/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	user := router.Group("/users")
	{

		user.POST("/register", controllers.UserRegister)
		user.POST("/login", controllers.UserLogin)
		user.PUT("/:userId", middlewares.Authentication(), controllers.UserUpdate)
		user.DELETE("/", middlewares.Authentication(), controllers.UserDelete)
	}

	photo := router.Group("/photos")
	{
		photo.Use(middlewares.Authentication())
		photo.POST("/", controllers.CreatePhoto)
		photo.GET("/", controllers.GetPhotos)
		photo.PUT("/:photoId", controllers.UpdatePhoto)
		photo.DELETE("/:photoId", controllers.DeletePhoto)
	}

	comment := router.Group("/comments")
	{
		comment.Use(middlewares.Authentication())
		comment.POST("/", controllers.CreateComment)
		comment.GET("/", controllers.GetComments)
		comment.PUT("/:commentId", controllers.UpdateComment)
		comment.DELETE("/:commentId", controllers.DeleteComment)
	}

	return router
}
