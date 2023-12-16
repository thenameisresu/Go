// main.go

package main

import (
	"timesheet_tool/db"
	"timesheet_tool/handlers"
	"timesheet_tool/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Use the authentication middleware
	r.Use(middleware.AuthMiddleware())

	// Initialize the database
	_, err := db.InitDB()
	if err != nil {
		panic(err)
	}

	// Serve static files from the "static" directory
	r.Static("/static", "./static")

	// Use the authentication middleware
	r.Use(middleware.AuthMiddleware())

	// Define API routes
	api := r.Group("/api/timesheet")
	{
		api.GET("/projects", handlers.GetProjects)
		api.GET("/subprojects", handlers.GetSubProjects)
		api.POST("/", handlers.SubmitTimesheet)
		api.POST("", handlers.SubmitTimesheet)
		api.POST("/login", handlers.Login)
		api.GET("/userDetails", handlers.UserDetails)
		api.GET("/downloadTimesheet", handlers.DownloadTimesheet)
	}

	// Define the route to serve the HTML file
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// Run the server
	r.Run(":8080")
}
