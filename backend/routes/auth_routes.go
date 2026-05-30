package routes

import (
	"spotify-clone/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, authController *controllers.AuthController) {
	auth := router.Group("/auth")

	auth.POST("/login", authController.Login)
}