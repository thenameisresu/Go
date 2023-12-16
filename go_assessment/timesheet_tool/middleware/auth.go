// middleware/auth.go

package middleware

import (
	"net/http"
	"strings"
	"timesheet_tool/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key") // Replace with your own secret key

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := extractTokenFromHeader(c.Request.Header.Get("Authorization"))
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		token, err := validateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			return
		}

		// Extract user details from claims
		userType, _ := claims["user_type"].(string)
		user := models.User{
			UserType: userType,
		}

		// Attach user details to the context for later use
		c.Set("user", user)

		c.Next()
	}
}

func extractTokenFromHeader(header string) string {
	parts := strings.Split(header, " ")
	if len(parts) == 2 {
		return parts[1]
	}
	return ""
}

func validateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
