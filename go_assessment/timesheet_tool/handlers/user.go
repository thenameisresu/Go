// handlers/user.go

package handlers

import (
	"net/http"
	"timesheet_tool/db"
	"timesheet_tool/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = []byte("your-secret-key")

// Login handles user login
func Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := db.GetUserByUsernameAndPassword(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to authenticate user"})
		return
	}

	if user != nil {
		// Generate a JWT token
		token, err := generateToken(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
			return
		}

		// Provide the token in the response
		c.JSON(http.StatusOK, gin.H{"success": true, "token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": "Invalid username or password"})
	}
}

// generateToken generates a JWT token for the provided user
func generateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_type": user.UserType,
		// Add more claims if needed
	})

	// Sign the token with a secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Register(c *gin.Context) {
	var userData models.User

	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before saving it to the database
	if err := userData.SetPassword(userData.Password); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash the password"})
		return
	}

	if err := db.CreateUser(&userData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}

func UserDetails(c *gin.Context) {
	// Retrieve user details from the context
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User details not found"})
		return
	}

	// Use user details as needed
	c.JSON(http.StatusOK, user)
}
