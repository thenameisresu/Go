// main.go

package main

import (
	"loginapp/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Define API routes
	api := r.Group("/api")
	{
		api.POST("/login", handlers.Login)
		api.GET("/userDetails", handlers.UserDetails)
	}

	// Define the route to serve the HTML file
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Run the server
	r.Run(":8080")
}
