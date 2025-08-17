package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"example.com/go-cognito/models"
	"example.com/go-cognito/utils"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/golang-jwt/jwt/v5"
)

// Define struct fields
type AuthService struct {
	CognitoClient *cognitoidentityprovider.Client
	ClientID string
	ClientSecret string
	Region string
	UserPoolID string
}

// Define constructor 
func NewAuthService(client *cognitoidentityprovider.Client, clientId, clientSecret, region, userPoolId string) *AuthService {
	if client == nil {
		log.Fatalf("Cognito Client cannot be nil.")
	}

	if clientId == "" || clientSecret == ""{
		log.Fatalf("Client ID or Client Secret cannot be nil")
	}
	
	return &AuthService{
		CognitoClient: client,
		ClientID: clientId,
		ClientSecret: clientSecret,
		Region: region,
		UserPoolID: userPoolId,
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
		return fmt.Errorf("Could not create new user: %w", err)
	}

	return nil
}

func (s *AuthService) SignIn(context context.Context, user models.SignInInput) (models.AuthResponse, error) {
	// var authResult *types.AuthenticationResultType

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
			return models.AuthResponse{}, errors.New(*resetRequired.Message)
		} 
		return models.AuthResponse{}, fmt.Errorf("Could not sign in user %s: %w", user.UserName, err)
	}

	if output.AuthenticationResult == nil || output.AuthenticationResult.IdToken == nil {
		return models.AuthResponse{}, errors.New("Authentication result or ID Token is nil.")
	} 

	authResult := output.AuthenticationResult
	response := models.NewAuthResponse(authResult.AccessToken, authResult.IdToken, authResult.RefreshToken, authResult.TokenType, authResult.ExpiresIn)
	
	return response, err
}

func (s *AuthService) ConfirmAccount(context context.Context, user models.UserConfirmation) (error) {
	_, err := s.CognitoClient.ConfirmSignUp(context, &cognitoidentityprovider.ConfirmSignUpInput{
		Username: aws.String(user.Email),
		ConfirmationCode: aws.String(user.Code),
		ClientId: aws.String(s.ClientID),
		SecretHash: aws.String(utils.GetSecretHash(s.ClientID, s.ClientSecret, user.Email)),
	})

	if err != nil {
		log.Printf("ConfirmSignUp failed for user %s: %v", user.Email, err)
		return fmt.Errorf("account confirmation failed: %w", err)
	}

	return nil
}

func (s *AuthService) VerifyToken(jwtToken string) (bool, error) {
	
	region := s.Region
	userPoolId := s.UserPoolID

	if region == "" || userPoolId == "" {
		return false, fmt.Errorf("Region or User Pool ID environment variables are not set.")
	}

	jwksURL := fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", region, userPoolId)

	// Create the keyfunc.Keyfunc.
	jwks, err := keyfunc.NewDefault([]string{jwksURL})
	if err != nil {
		return false, fmt.Errorf("Failed to create JWKS from URL: %v", err)
	}

	// Parse the JWT
	token, err := jwt.Parse(jwtToken, jwks.Keyfunc)
	if err != nil {
		return false, fmt.Errorf("Failed to parse JWT: %v", err)
	}

	if !token.Valid {
		return false, fmt.Errorf("The token is not valid.")
	}

	log.Println("The token is valid.")

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		return false, fmt.Errorf("Failed to parse claims.")
	}

	tokenUse, ok := claims["token_use"].(string)

	if !ok || tokenUse != "access" {
		return false, fmt.Errorf("Invalid token use - must be an access token.")
	}

	log.Printf("Token Claims: %+v\n", claims)

	exp, ok := claims["exp"].(float64)

	if !ok {
		return false, fmt.Errorf("Failed to get exp claim.")
	}

	if time.Now().Unix() > int64(exp) {
		return false, fmt.Errorf("Token has expired.")
	}

	return true, nil
}

func (s *AuthService) ForgotPassword(context context.Context, user models.ForgotPasswordInput) (*types.CodeDeliveryDetailsType, error) {
	output, err := s.CognitoClient.ForgotPassword(context, &cognitoidentityprovider.ForgotPasswordInput{
		ClientId: aws.String(s.ClientID),
		Username: aws.String(user.UserName),
		SecretHash: aws.String(utils.GetSecretHash(s.ClientID, s.ClientSecret, user.UserName)),
	})

	if err != nil {
		log.Printf("Couldn't start password reset for user '%v'. Here;s why: %v\n", user.UserName, err)
		return nil, fmt.Errorf("Password reset failed: %w", err)
	}

	log.Println(output.CodeDeliveryDetails)
	return output.CodeDeliveryDetails, nil 
}

func (s *AuthService) ConfirmForgotPassword(context context.Context, user models.ConfirmForgotPasswordInput) error {
	_, err := s.CognitoClient.ConfirmForgotPassword(context, &cognitoidentityprovider.ConfirmForgotPasswordInput{
		ClientId:         aws.String(s.ClientID),
		ConfirmationCode: aws.String(user.ConfirmationCode),
		Password:         aws.String(user.Password),
		Username:         aws.String(user.UserName),
		SecretHash: 	  aws.String(utils.GetSecretHash(s.ClientID, s.ClientSecret, user.UserName)),
	})
	if err != nil {
		var invalidPassword *types.InvalidPasswordException
		if errors.As(err, &invalidPassword) {
			log.Println(*invalidPassword.Message)
		} else {
			log.Printf("Couldn't confirm user %v. Here's why: %v", user.UserName, err)
			return fmt.Errorf("Password reset failed: %w", err)
		}
	}
	return nil
}

func (s *AuthService) ResendConfirmationCode(context context.Context, user models.ForgotPasswordInput) (*types.CodeDeliveryDetailsType, error) {
	output, err := s.CognitoClient.ResendConfirmationCode(context, &cognitoidentityprovider.ResendConfirmationCodeInput{
		ClientId: aws.String(s.ClientID),
		Username: aws.String(user.UserName),
		SecretHash: aws.String(utils.GetSecretHash(s.ClientID, s.ClientSecret, user.UserName)),
	})

	if err != nil {
		log.Printf("Couldn't resend confirmation code to user '%v'. Here;s why: %v\n", user.UserName, err)
		return nil, fmt.Errorf("Password reset failed: %w", err)
	}

	return output.CodeDeliveryDetails, nil
}

func (s *AuthService) GetTokensFromRefreshToken(context context.Context, user models.RefreshTokenInput) (*types.AuthenticationResultType, error) {
	output, err := s.CognitoClient.GetTokensFromRefreshToken(context, &cognitoidentityprovider.GetTokensFromRefreshTokenInput{
		ClientId: aws.String(s.ClientID),
		RefreshToken: aws.String(user.RefreshToken),
		ClientSecret: aws.String(s.ClientSecret),
	})

	if err != nil {
		return nil, fmt.Errorf("Could not obtain new token: %w", err)
	}

	return output.AuthenticationResult, nil
}