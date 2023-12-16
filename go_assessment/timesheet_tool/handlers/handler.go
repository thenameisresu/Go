// handlers/handlers.go

package handlers

import (
	"net/http"
	"strconv"

	"timesheet_tool/db"
	"timesheet_tool/models"

	"github.com/gin-gonic/gin"
)

// GetProjects retrieves a list of all projects
func GetProjects(c *gin.Context) {
	projects, err := db.FetchProjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"projects": projects})
}

// GetSubProjects retrieves a list of subprojects for a given project ID
func GetSubProjects(c *gin.Context) {
	projectIDStr := c.Query("projectID")
	if projectIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing projectID parameter"})
		return
	}

	// Convert projectIDStr to an integer
	projectID, err := strconv.Atoi(projectIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid projectID parameter"})
		return
	}

	subprojects, err := db.FetchAllSubProjects(projectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch subprojects"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"subprojects": subprojects})
}

func SubmitTimesheet(c *gin.Context) {
	var timesheet models.Timesheet

	if err := c.ShouldBindJSON(&timesheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate and save the timesheet to the database
	err := db.SaveTimesheet(timesheet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save timesheet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Timesheet submitted successfully"})
}
