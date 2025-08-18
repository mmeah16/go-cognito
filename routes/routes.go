package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/go-cognito/handlers"
	"example.com/go-cognito/middleware"
)

func RegisterRoutes(server *gin.Engine, middlewareHandler *middleware.MiddlewareHandler, authHandler *handlers.AuthHandler) {
	authGroup := server.Group("/auth")
	{
		authGroup.POST("/signUp", authHandler.SignUp)
		authGroup.POST("/signIn", authHandler.SignIn)
		authGroup.POST("/confirmAccount", authHandler.ConfirmAccount)
		authGroup.POST("/forgotPassword", authHandler.ForgotPassword)
		authGroup.POST("/confirmForgotPassword", authHandler.ConfirmForgotPassword)
		authGroup.POST("/resendConfirmationCode", authHandler.ResendConfirmationCode)
		authGroup.POST("/refreshToken", authHandler.GetTokensFromRefreshToken)
		authGroup.POST("/signOut", authHandler.SignOut)
	}

	authenticated := server.Group("/")
	authenticated.Use(middlewareHandler.Authenticate)
	authenticated.GET("/health", health)
}

func health(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}