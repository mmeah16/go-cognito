package services

import (
	"context"
	"errors"
	"fmt"
	"log"

	"example.com/go-cognito/models"
	"example.com/go-cognito/utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

// Define struct fields
type AuthService struct {
	CognitoClient *cognitoidentityprovider.Client
	ClientID string
	ClientSecret string
}

// Define constructor 
func NewAuthService(client *cognitoidentityprovider.Client, clientID, clientSecret string) *AuthService {
	if client == nil {
		log.Fatalf("Cognito Client cannot be nil.")
	}

	if clientID == "" || clientSecret == ""{
		log.Fatalf("Client ID or Client Secret cannot be nil")
	}
	
	return &AuthService{
		CognitoClient: client,
		ClientID: clientID,
		ClientSecret: clientSecret,
	}
}

// Implement SignUp business logic
func (s *AuthService) SignUp(context context.Context, user models.SignUpInput) error {
	// Use SignUp API to register the user
	_, err := s.CognitoClient.SignUp(context, &cognitoidentityprovider.SignUpInput{
		ClientId: aws.String(s.ClientID),
		Username: aws.String(user.UserName),
		Password: aws.String(user.Password),
		UserAttributes: []types.AttributeType{
			{Name: aws.String("name"), Value: aws.String(user.Name)},
		},
		SecretHash: aws.String(utils.GetSecretHash(s.ClientID, s.ClientSecret, user.UserName)),
	})

	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			return errors.New(*invalidPassword.Message)
		}
		return fmt.Errorf("could not create new user: %w", err)
	}

	return nil
}

func (s *AuthService) SignIn(context context.Context, user models.SignInInput) (string, error) {
	var authResult *types.AuthenticationResultType

	output, err := s.CognitoClient.InitiateAuth(context, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: "USER_PASSWORD_AUTH",
		ClientId: aws.String(s.ClientID),
		AuthParameters: map[string]string{
			"USERNAME": user.UserName, 
			"PASSWORD": user.Password, 
			"SECRET_HASH": utils.GetSecretHash(s.ClientID, s.ClientSecret, user.UserName),
		},
	})

	if err != nil {
		var resetRequired *types.PasswordResetRequiredException
		if errors.As(err, &resetRequired) {
			log.Println(*resetRequired.Message)
			return "", errors.New(*resetRequired.Message)
		} 
		return "", fmt.Errorf("couldn't sign in user %s: %w", user.UserName, err)
	}

	if output.AuthenticationResult == nil || output.AuthenticationResult.IdToken == nil {
		return "", errors.New("Authentication result or ID Token is nil.")
	} else {
		authResult = output.AuthenticationResult
	}
	
	return *authResult.IdToken, err
}