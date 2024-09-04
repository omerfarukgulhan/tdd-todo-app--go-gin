package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"todo-app--go-gin/common/util/results"
	"todo-app--go-gin/controller/constants"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/service"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (authController *AuthController) RegisterAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authController.Register)
		authGroup.POST("/login", authController.Login)
	}
}

func (authController *AuthController) Register(ctx *gin.Context) {
	var newUser request.UserCreate
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		log.Printf("error in bind")
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Enter user in valid format"))
		return
	}

	err := authController.authService.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, results.NewResult(true, constants.DataAdded))
}

func (authController *AuthController) Login(context *gin.Context) {

}
