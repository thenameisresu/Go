// handlers/user.go

package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
