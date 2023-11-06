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
		photo.POST("/", middlewares.Authentication(), controllers.CreatePhoto)
	}

	return router
}
