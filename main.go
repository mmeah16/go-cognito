package main

import (
	"log"

	"example.com/go-cognito/config"
	_ "example.com/go-cognito/docs"
	"example.com/go-cognito/handlers"
	"example.com/go-cognito/middleware"
	"example.com/go-cognito/routes"
	"example.com/go-cognito/services"
	"example.com/go-cognito/utils"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           AWS Cognito and Go Gin
// @version         1.0
// @description     These APIs support an end-to-end authentication workflow using AWS Cognito to support identity and access management for Gin APIs.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// Creates and returns configuration with environment variables
	config := config.LoadConfig()

	// Create client to interact with AWS Cognito
	client, err := utils.CreateCognitoClient()

	if err != nil {
		log.Fatalf("Failed to create Cognito client: %v", err)
	}

	// Create authService using Cognito client, client ID, and client secret
	authService := services.NewAuthService(client.CognitoClient, config.ClientId, config.ClientSecret, config.Region, config.UserPoolId)
	authHandler := handlers.NewAuthHandler(authService)
	middlewareHandler := middleware.NewMiddlewareHandler(authService)

	server := gin.Default()
	routes.RegisterRoutes(server, middlewareHandler, authHandler)
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	server.Run(":8080")
}
