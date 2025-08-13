package middleware

import (
	"net/http"
	"strings"

	"example.com/go-cognito/services"
	"github.com/gin-gonic/gin"
)

type MiddlewareHandler struct {
	Service *services.AuthService
}

func NewMiddlewareHandler(service *services.AuthService) *MiddlewareHandler {
	return &MiddlewareHandler{
		Service: service,
	}
}

func (s *MiddlewareHandler) Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Authorization token is required."})
		return
	}

	token = strings.TrimSpace(strings.TrimPrefix(token, "Bearer "))

	valid, err := s.Service.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token."})
		return
	}

	if !valid {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid authorization token."})
		return
	}

	context.Next()
	
}