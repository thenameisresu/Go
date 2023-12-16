// excel/excel.go

package excel

import (
	"fmt"
	"time"

	"timesheet_tool/models"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// CreateExcel generates an Excel file based on timesheet data
func CreateExcel(timesheets []models.Timesheet) (*excelize.File, error) {
	file := excelize.NewFile()

	// Create a new sheet
	sheet := file.NewSheet("Sheet1")

	// Set header row
	headers := []string{"ID", "Project Name", "Sub Project Name", "JIRA/SNOW ID", "Task Description", "Hours Spent", "Comments", "Created At"}
	for col, header := range headers {
		cell := excelize.ToAlphaString(col+1) + "1"
		file.SetCellValue("Sheet1", cell, header)
	}

	// Populate data rows
	for row, timesheet := range timesheets {
		file.SetCellValue("Sheet1", fmt.Sprintf("A%d", row+2), timesheet.ID)
		file.SetCellValue("Sheet1", fmt.Sprintf("B%d", row+2), timesheet.ProjectName)
		file.SetCellValue("Sheet1", fmt.Sprintf("C%d", row+2), timesheet.SubProjectName)
		file.SetCellValue("Sheet1", fmt.Sprintf("D%d", row+2), timesheet.JiraSnowID)
		file.SetCellValue("Sheet1", fmt.Sprintf("E%d", row+2), timesheet.TaskDescription)
		file.SetCellValue("Sheet1", fmt.Sprintf("F%d", row+2), timesheet.HoursSpent)
		file.SetCellValue("Sheet1", fmt.Sprintf("G%d", row+2), timesheet.Comments)
		file.SetCellValue("Sheet1", fmt.Sprintf("H%d", row+2), timesheet.CreatedAt.Format(time.RFC3339))
	}

	// Set active sheet to the newly created sheet
	file.SetActiveSheet(sheet)

	return file, nil
}
