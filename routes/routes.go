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
		authGroup.POST("/signup", authHandler.SignUp)
		authGroup.POST("/login", authHandler.SignIn)
		authGroup.POST("/confirmAccount", authHandler.ConfirmAccount)
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