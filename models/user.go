package models

type SignUpInput struct {
	UserName string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type SignInInput struct {
	UserName string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserConfirmation struct {
	Email string `json:"email" binding:"required,email"`
	Code string `json:"code" binding:"required"`
}

type ForgotPasswordInput struct {
	UserName string `json:"username" binding:"required,email"`
}

type ConfirmForgotPasswordInput struct {
	UserName string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	ConfirmationCode string `json:"code" binding:"required"`
}

type RefreshTokenInput struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}