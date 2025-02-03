package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func Bootstrap() {
	connection, err := sql.Open("sqlite3", "database.db")

	if err != nil {
		panic("Failed to open database connection")
	}

	connection.SetMaxOpenConns(10)
	connection.SetMaxIdleConns(5)

	Database = connection

	createTables()
}

func createTables() {
	_, err := Database.Exec(getUsersStructTable())

	if err != nil {
		panic("Failed to create users table: " + err.Error())
	}

	fmt.Println(">> Users table created or already exists.")

	_, err = Database.Exec(getEventStructTable())

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
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`
}

func getUsersStructTable() string {
	return `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`
}
