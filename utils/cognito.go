// Create Cognito Client

package utils

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type CognitoActions struct {
	CognitoClient *cognitoidentityprovider.Client
}

func CreateCognitoClient() (CognitoActions, error) {
	context := context.Background()

	sdkConfig, err := config.LoadDefaultConfig(context)

	if err != nil {
		log.Fatalf("Couldn't load default configuration: %v", err)
	}

	cognitoClient := cognitoidentityprovider.NewFromConfig(sdkConfig)

	actor := CognitoActions{CognitoClient: cognitoClient}

	return actor, nil
}

