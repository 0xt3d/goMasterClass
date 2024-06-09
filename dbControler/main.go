package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Connection string for your PostgreSQL database
	connStr := "user=your_user dbname=your_db sslmode=disable password=your_password host=localhost port=5432"

	// Open a connection to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query the database
	rows, err := db.Query("SELECT * FROM ")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterate through the rows and print them out
	for rows.Next() {
		var id int
		var name string
		// Add more variables here to match your table's columns

		err := rows.Scan(&id, &name /* Add more variables here */)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name /* Print out more variables here */)
	}

	// Check for any errors that happened during the iteration
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
