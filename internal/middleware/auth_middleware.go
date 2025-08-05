package middleware

import (
	"net/http"
	"strings"

	"lsp-api/internal/services"
	"lsp-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(authService services.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Authorization header is required"))
			return
		}

		// Check if the header has the Bearer prefix
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid authorization format, expected 'Bearer {token}'"))
			return
		}

		tokenString := parts[1]

		// Validate token
		token, err := authService.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid or expired token"))
			return
		}

		// Extract claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Failed to parse token claims"))
			return
		}

		// Set user ID in context
		userID, ok := claims["user_id"].(float64)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, utils.ErrorResponse("Invalid token claims"))
			return
		}

		c.Set("userID", uint(userID))
		c.Set("email", claims["email"])
		c.Set("username", claims["username"])

		c.Next()
	}
}
