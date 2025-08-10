package routes

import (
	"github.com/gin-gonic/gin"

	"example.com/go-cognito/handlers"
)

func RegisterRoutes(server *gin.Engine, authHandler *handlers.AuthHandler) {
	authGroup := server.Group("/auth")
	{
	authGroup.POST("/signup", authHandler.SignUp)
	authGroup.POST("/login", authHandler.SignIn)
	}
}