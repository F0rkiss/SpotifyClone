package main

import (
	"log"
	"spotify-clone/controllers"
	"spotify-clone/database"
	"spotify-clone/repository"
	"spotify-clone/routes"
	"spotify-clone/services"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.ConnectDB()

	userRepo := repository.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	router := gin.Default()
	routes.AuthRoutes(router, authController)

	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"massage" : "pong", 
		})
	})

	log.Println("server starting on : 8080...")
	router.Run(":8080")
}