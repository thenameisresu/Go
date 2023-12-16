// db/mysql.go

package db

import (
	"database/sql"
	"fmt"

	"timesheet_tool/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/test"
	database, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = database.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")

	// Set the global database variable
	db = database

	return db, nil
}

// FetchProjects retrieves a list of all projects from the database
func FetchProjects() ([]models.Project, error) {
	rows, err := db.Query("SELECT ProjectID, ProjectName FROM Projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []models.Project
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.ProjectID, &project.ProjectName)
		fmt.Println("Project ID: ", project.ProjectID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

// FetchAllSubProjects retrieves a list of all subprojects for a given project ID
func FetchAllSubProjects(projectID int) ([]models.SubProject, error) {
	rows, err := db.Query("SELECT SubProjectID, SubProjectName FROM SubProjects WHERE ProjectID = ?", projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subprojects []models.SubProject
	for rows.Next() {
		var subproject models.SubProject
		err := rows.Scan(&subproject.SubProjectID, &subproject.SubProjectName)
		if err != nil {
			return nil, err
		}
		subprojects = append(subprojects, subproject)
	}

	return subprojects, nil
}

// SaveTimesheet saves a timesheet entry to the database
func SaveTimesheet(timesheet models.Timesheet) error {
	// Prepare the SQL query
	query := `
		INSERT INTO Timesheets (ProjectID, SubProjectID, JiraSnowID, TaskDescription, HoursSpent, Comments, CreatedAt)
		VALUES (?, ?, ?, ?, ?, ?, NOW())
	`

	// Execute the SQL query with timesheet data
	_, err := db.Exec(
		query,
		timesheet.ProjectID,
		timesheet.SubProjectID,
		timesheet.JiraSnowID,
		timesheet.TaskDescription,
		timesheet.HoursSpent,
		timesheet.Comments,
	)
	fmt.Println(query)
	if err != nil {
		fmt.Println("Error saving timesheet:", err)
		return err
	}

	fmt.Println("Timesheet saved successfully!")
	return nil
}
