// handlers/timesheet.go

package handlers

import (
	"fmt"
	"net/http"
	"time"
	"timesheet_tool/db"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func DownloadTimesheet(c *gin.Context) {
	// Get start and end dates from the query parameters
	startDateString := c.Query("start_date")
	endDateString := c.Query("end_date")

	// Parse dates
	startDate, err := time.Parse("2006-01-02", startDateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	// Fetch timesheet data based on the date range
	timesheets, err := db.GetTimesheetDataByDateRange(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch timesheet data"})
		return
	}

	// Create a new Excel file
	file := excelize.NewFile()

	// Create a new sheet
	sheetName := "Sheet1"
	file.NewSheet(sheetName)

	// Set header row
	headers := []string{"Project ID", "Sub Project ID", "Jira/SNOW ID", "Task Description", "Hours Spent", "Comments", "Created At"}
	for col, header := range headers {
		colName, _ := excelize.ColumnNumberToName(col + 1)
		cell := fmt.Sprintf("%s1", colName)
		file.SetCellValue(sheetName, cell, header)
	}

	// Add data to the sheet
	for row, timesheet := range timesheets {
		rowIndex := row + 2 // Excel rows start from 1, so we add 1
		file.SetCellValue(sheetName, fmt.Sprintf("A%d", rowIndex), timesheet.ProjectID)
		file.SetCellValue(sheetName, fmt.Sprintf("B%d", rowIndex), timesheet.SubProjectID)
		file.SetCellValue(sheetName, fmt.Sprintf("C%d", rowIndex), timesheet.JiraSnowID)
		file.SetCellValue(sheetName, fmt.Sprintf("D%d", rowIndex), timesheet.TaskDescription)
		file.SetCellValue(sheetName, fmt.Sprintf("E%d", rowIndex), timesheet.HoursSpent)
		file.SetCellValue(sheetName, fmt.Sprintf("F%d", rowIndex), timesheet.Comments)
		file.SetCellValue(sheetName, fmt.Sprintf("G%d", rowIndex), timesheet.CreatedAt.Format("2006-01-02 15:04:05"))
	}

	// Save the Excel file
	filePath := fmt.Sprintf("./static/timesheet_%s_%s.xlsx", startDateString, endDateString)
	if err := file.SaveAs(filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save Excel file"})
		return
	}

	// Respond with the file for download
	c.File(filePath)
}
