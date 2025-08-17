package models

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type AuthResponse struct {
    AccessToken  string `json:"accessToken"`
    IdToken      string `json:"idToken"`
    RefreshToken string `json:"refreshToken"`
    TokenType    string `json:"tokenType"`
    ExpiresIn    int32  `json:"expiresIn"`
}

func NewAuthResponse(accessToken, idToken, refreshToken, tokenType *string, expiresIn int32) AuthResponse {
    return AuthResponse{
        AccessToken:  aws.ToString(accessToken),
        IdToken:      aws.ToString(idToken),
        RefreshToken: aws.ToString(refreshToken),
        TokenType:    aws.ToString(tokenType),
        ExpiresIn:    expiresIn,
    }
}
