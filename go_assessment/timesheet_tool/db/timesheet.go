// db/timesheet.go

package db

import (
	"time"
	"timesheet_tool/models"
)

// GetTimesheetDataByDateRange retrieves timesheet data from the database based on the date range
func GetTimesheetDataByDateRange(startDate, endDate time.Time) ([]models.Timesheet, error) {
	var timesheets []models.Timesheet

	// Replace the placeholders with your actual database connection and query
	rows, err := db.Query("SELECT project_id, sub_project_id, jira_snow_id, task_description, hours_spent, comments, created_at FROM timesheets WHERE created_at BETWEEN ? AND ?", startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var timesheet models.Timesheet
		err := rows.Scan(&timesheet.ProjectID, &timesheet.SubProjectID, &timesheet.JiraSnowID, &timesheet.TaskDescription, &timesheet.HoursSpent, &timesheet.Comments, &timesheet.CreatedAt)
		if err != nil {
			return nil, err
		}
		timesheets = append(timesheets, timesheet)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return timesheets, nil
}
