package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func Bootstrap() {
	connection, err := sql.Open("sqlite3", "events_database.db")

	if err != nil {
		panic("Failed to connect to database")
	}

	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(5)

	Database = connection

	createTables()
}

func createTables() {
	eventsTable := getEventStructTable()

	_, err := Database.Exec(eventsTable)

	if err != nil {
		panic("Failed to create events table")
	}

	fmt.Println(">> Events table created or already exists.")
}

func getEventStructTable() string {
	return `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			user_id INTEGER NOT NULL
		)
	`
}
