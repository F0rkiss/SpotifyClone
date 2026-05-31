package controllers

import (
	"net/http"
	"spotify-clone/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

type LoginRequest struct {
	Identifier    string `json:"identifier"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (controller *AuthController) Login(ctx *gin.Context) {
	var req LoginRequest
		// ambil data json dari request body
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

		// kirim data register ke service layer
	user, err := controller.authService.Login(req.Identifier, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"mesage": err.Error(),
		})
		return
	}
		// kirim data register ke service layer
	ctx.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"user": gin.H{
			"id":    user.Id,
			"email": user.Email,
		},
	})
}

func (controller *AuthController) Register(ctx *gin.Context){
	var req RegisterRequest
	
	// ambil data json dari request body
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request",
		})
		return
	}

	// kirim data register ke service layer
	user, err := controller.authService.Register(req.Email, req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// ngasih tau kalo register berhasil
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "register success", 
		"user": gin.H{
			"id": user.Id,
			"email": user.Email,
			"username": user.Username,
		}, 
	})
}
