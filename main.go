package main

import (
	"log"

	"example.com/go-cognito/config"
	"example.com/go-cognito/handlers"
	"example.com/go-cognito/routes"
	"example.com/go-cognito/services"
	"example.com/go-cognito/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates and returns configuration with environment variables
	config := config.LoadConfig()

	// Create client to interact with AWS Cognito
	client, err := utils.CreateCognitoClient()

	if err != nil {
		log.Fatalf("Failed to create Cognito client: %v", err)
	}

	// Create authService using Cognito client, client ID, and client secret
	authService := services.NewAuthService(client.CognitoClient, config.ClientId, config.ClientSecret, config.Region)
	authHandler := handlers.NewAuthHandler(authService)

	server := gin.Default()
	routes.RegisterRoutes(server, authHandler)
	server.Run(":8080")
}
