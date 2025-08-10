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

