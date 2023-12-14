// db.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Component represents a record in the Component table
type Component struct {
	CompID      int
	CompName    string
	Description string
}

func InitDB() {
	dataSourceName := "root:@tcp(127.0.0.1:3306)/test"
	var err error
	db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")
}

func CloseDB() {
	db.Close()
}

func GetComponents() ([]Component, error) {
	rows, err := db.Query("SELECT CompID, CompName, Description FROM Component")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var components []Component
	for rows.Next() {
		var component Component
		err := rows.Scan(
			&component.CompID,
			&component.CompName,
			&component.Description,
		)
		if err != nil {
			return nil, err
		}
		components = append(components, component)
	}

	return components, nil
}

func GetComponentByID(compID int) (Component, error) {
	var c Component
	err := db.QueryRow("SELECT CompID, CompName, Description FROM Component WHERE CompID=?", compID).
		Scan(&c.CompID, &c.CompName, &c.Description)
	if err != nil {
		return Component{}, err
	}

	return c, nil
}

func CreateComponent(c Component) (int64, error) {
	result, err := db.Exec("INSERT INTO Component (CompID, CompName, Description) VALUES (?, ?, ?)",
		c.CompID, c.CompName, c.Description)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateComponent(c Component) error {
	_, err := db.Exec("UPDATE Component SET CompName=?, Description=? WHERE CompID=?",
		c.CompName, c.Description, c.CompID)
	return err
}

func DeleteComponent(compID int) error {
	_, err := db.Exec("DELETE FROM Component WHERE CompID=?", compID)
	return err
}
