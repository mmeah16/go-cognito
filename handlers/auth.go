package handlers

import (
	"net/http"

	"example.com/go-cognito/models"
	"example.com/go-cognito/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) SignUp(context *gin.Context) {

	var user models.SignUpInput

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid input data."})
		return
	}

	err = h.Service.SignUp(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Successfully signed up user!"})

}

func (h *AuthHandler) SignIn(context *gin.Context) {
	var user models.SignInInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid input data."})
		return
	}

	token, err := h.Service.SignIn(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Login successful.", "token" : token})

}

func (h *AuthHandler) ConfirmAccount(context *gin.Context) {
	var user models.UserConfirmation

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid input data."})
		return
	}

	err = h.Service.ConfirmAccount(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Account confirmed."})
}

func (h *AuthHandler) ForgotPassword(context *gin.Context) {
	var user models.ForgotPasswordInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid input data."})
		return
	}

	output, err := h.Service.ForgotPassword(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : output})
}

func (h *AuthHandler) ConfirmForgotPassword(context *gin.Context) {
	var user models.ConfirmForgotPasswordInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : "Invalid input data."})
		return
	}

	err = h.Service.ConfirmForgotPassword(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message" : err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message" : "Password successfully changed."})
}
