package controllers

import (
	"net/http"
	"spotify-clone/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct{
	authService services.AuthService
}

func NewAuthController (authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

type LoginRequest struct {
	Email 	string	`json:"email"`
	Password 	string	`json:"password"`
}

func (controller *AuthController) Login (ctx *gin.Context){
	var req LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}
	user, err := controller.authService.Login(req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"mesage": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user" : gin.H{
			"id": user.Id,
			"email": user.Email,
		},
	})
}
