// db/user.go

package db

import (
	"database/sql"
	"fmt"
	"timesheet_tool/models"
)

// GetUserByUsernameAndPassword retrieves user from the database based on username and password
func GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	var user models.User
	fmt.Println("inside GetUserByUsernameAndPassword")
	query := "SELECT user_id, user_type, email, username, password FROM users WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&user.UserID, &user.UserType, &user.Email, &user.Username, &user.Password)
	fmt.Println("GetUserByUsernameAndPassword : query :", query)
	if err == sql.ErrNoRows {
		// No user found
		return nil, nil
	} else if err != nil {
		// Other database error
		return nil, err
	}
	fmt.Println("GetUserByUsernameAndPassword : user :", user)
	return &user, nil
}

// GetUserByEmail retrieves user from the database based on email address
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	query := "SELECT user_id, user_type, email, username FROM users WHERE email = ?"
	err := db.QueryRow(query, email).Scan(&user.UserID, &user.UserType, &user.Email, &user.Username)

	if err == sql.ErrNoRows {
		// No user found
		return nil, nil
	} else if err != nil {
		// Other database error
		return nil, err
	}

	return &user, nil
}

// CreateUser creates a new user entry in the database
func CreateUser(user *models.User) error {
	query := "INSERT INTO users (user_type, email, username, password) VALUES (?, ?, ?, ?)"
	result, err := db.Exec(query, user.UserType, user.Email, user.Username, user.Password)

	if err != nil {
		return err
	}

	// Get the ID of the newly created user and update the user struct
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.UserID = int(lastInsertID)

	return nil
}

// GetUserByUsername retrieves user from the database based on username
func GetUserByUsername(username string) (*models.User, error) {
	var user models.User

	// Replace the placeholders with your actual database connection and query
	row := db.QueryRow("SELECT user_id, user_type, email, username, password FROM users WHERE username = ?", username)
	err := row.Scan(&user.UserID, &user.UserType, &user.Email, &user.Username, &user.Password)

	if err == sql.ErrNoRows {
		// No user found with the given username
		return nil, nil
	} else if err != nil {
		// An error occurred while fetching the user
		return nil, err
	}

	return &user, nil
}
