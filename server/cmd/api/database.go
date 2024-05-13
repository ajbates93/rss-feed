package main

import "database/sql"

func InitDB() {
	// Initialise the database connection
	db, err := sql.Open("postrges")
}
