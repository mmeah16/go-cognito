// Loads environment variables and creates configuration

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientId string
	ClientSecret string
	Region string
}

func LoadConfig() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Could not retrieve environment variables.")
	}

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	region := os.Getenv("REGION")

	if clientId == "" || clientSecret == "" {
		log.Fatal("Could not retrieve Client ID or Client Secret.")
	}
	return &Config{
		ClientId: clientId,
		ClientSecret: clientSecret,
		Region: region,
	}
}