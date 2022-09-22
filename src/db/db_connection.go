package db

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type Tag struct {
	actor_id    int    `json:"actor_id"`
	first_name  string `json:"first_name"`
	last_name   string `json:"last_name"`
	last_update string `json:"last_update"`
}

func CreateCon() {
	db, err := sql.Open("mysql", "root:Shiv@131218@tcp(127.0.0.1:3306)/sakila")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("db is connected")

		// make sure connection is available
		err = db.Ping()
		fmt.Println("Error:", err)
		if err != nil {
			fmt.Println("MySQL db is not connected")
		}
	}

	//create output file
	file, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("File does not exists or cannot be created")
		os.Exit(1)
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	//SQL Connection
	res, err := db.Query("SELECT * FROM actor LIMIT 10")

	for res.Next() {
		var tag Tag
		err = res.Scan(&tag.actor_id, &tag.first_name, &tag.last_name, &tag.last_update)

		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Fprintf(w, "actor_id: %5v       first_name: %10v       last_name: %15v       last_update: %15v\n", tag.actor_id, tag.first_name, tag.last_name, tag.last_update)

	}
	w.Flush()
}
