// models/models.go

package models

import "time"

// Project represents the structure of a project
type Project struct {
	ProjectID   int    `json:"project_id"`
	ProjectName string `json:"project_name"`
}

// SubProject represents the structure of a subproject
type SubProject struct {
	SubProjectID   int    `json:"sub_project_id"`
	SubProjectName string `json:"sub_project_name"`
	ProjectID      int    `json:"project_id"`
}

// Timesheet represents the structure of a timesheet entry
type Timesheet struct {
	ProjectID       string    `json:"project_id"`
	SubProjectID    string    `json:"sub_project_id"`
	JiraSnowID      string    `json:"jira_snow_id"`
	TaskDescription string    `json:"task_description"`
	HoursSpent      int       `json:"hours_spent"`
	Comments        string    `json:"comments"`
	CreatedAt       time.Time `json:"created_at"`
}
