package handlers

import (
	"log"
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

// SignUp godoc
// @Summary      Sign up a new user
// @Description  Creates a new user account with the provided details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.SignUpInput  true  "Sign up data"
// @Success      200   "Successfully signed up user!"
// @Failure      400   "Invalid input data or signup failed".
// @Router       /auth/signUp [post]
func (h *AuthHandler) SignUp(context *gin.Context) {

	var user models.SignUpInput

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	err = h.Service.SignUp(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully signed up user!"})
}

// SignIn godoc
// @Summary      Sign in a user
// @Description  Logins a user account with provided details
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.SignInInput  true  "Sign in data"
// @Success      200   {object}  models.AuthResponse "Successfully logged in user!"
// @Failure      400   "Invalid input data or sign in failed".
// @Router       /auth/signIn [post]
func (h *AuthHandler) SignIn(context *gin.Context) {
	var user models.SignInInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	authResult, err := h.Service.SignIn(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, authResult)
}

// SignUp godoc
// @Summary      Confirm user account.
// @Description  Confirms a user account with provided confirmation code sent by AWS via email.
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        user  body      models.UserConfirmationInput  true  "Confirmation code"
// @Success      200   "Account confirmed."
// @Failure      400   "Invalid input data".
// @Router       /auth/confirmAccount [post]
func (h *AuthHandler) ConfirmAccount(context *gin.Context) {
	var user models.UserConfirmationInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	err = h.Service.ConfirmAccount(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Account confirmed."})
}

func (h *AuthHandler) ForgotPassword(context *gin.Context) {
	var user models.ForgotPasswordInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	output, err := h.Service.ForgotPassword(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": output})
}

func (h *AuthHandler) ConfirmForgotPassword(context *gin.Context) {
	var user models.ConfirmForgotPasswordInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	err = h.Service.ConfirmForgotPassword(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Password successfully changed."})
}

func (h *AuthHandler) ResendConfirmationCode(context *gin.Context) {
	var user models.ForgotPasswordInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	output, err := h.Service.ResendConfirmationCode(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": output})
}

func (h *AuthHandler) GetTokensFromRefreshToken(context *gin.Context) {
	var user models.RefreshTokenInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	log.Printf("refreshToken: %q", user.RefreshToken)

	output, err := h.Service.GetTokensFromRefreshToken(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": output})
}

func (h *AuthHandler) SignOut(context *gin.Context) {
	var user models.SignOutInput

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input data."})
		return
	}

	log.Printf("refreshToken: %q", user.AccessToken)

	output, err := h.Service.SignOut(context, user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": output})
}
