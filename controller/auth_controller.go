package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-app--go-gin/common/util/results"
	"todo-app--go-gin/controller/constants"
	"todo-app--go-gin/controller/middlewares"
	"todo-app--go-gin/domain/request"
	"todo-app--go-gin/domain/response"
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
		authGroup.GET("/test", middlewares.Authenticate, authController.Test)
	}
}

func (authController *AuthController) Register(ctx *gin.Context) {
	var newUser request.UserCreate
	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Enter user in valid format"))
		return
	}

	token, err := authController.authService.Register(newUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	authResponse := response.NewAuthResponse(token)

	ctx.JSON(http.StatusCreated, results.NewDataResult(true, constants.RegisterSuccess, authResponse))
}

func (authController *AuthController) Login(ctx *gin.Context) {
	var newSignInCredentials request.SignInCredentials
	if err := ctx.ShouldBindJSON(&newSignInCredentials); err != nil {
		ctx.JSON(http.StatusBadRequest, results.NewResult(false, "Enter user in valid format"))
		return
	}

	token, err := authController.authService.Login(newSignInCredentials)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, results.NewResult(false, err.Error()))
		return
	}

	authResponse := response.NewAuthResponse(token)

	ctx.JSON(http.StatusCreated, results.NewDataResult(true, constants.LoginSuccess, authResponse))
}

func (authController *AuthController) Test(context *gin.Context) {
	// Retrieve the user ID from the context
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, results.NewResult(false, "User ID not found in context"))
		return
	}

	// Return the user ID in the JSON response
	context.JSON(http.StatusOK, results.NewDataResult(true, "User ID retrieved successfully", gin.H{"userId": userId}))
}
