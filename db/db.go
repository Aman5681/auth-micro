package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Failed to connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		eventId TEXT PRIMARY KEY, -- UUID as TEXT
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		userId TEXT NOT NULL,  -- string userId
		FOREIGN KEY(userId) REFERENCES users(userId)
	);
	`

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		userId TEXT PRIMARY KEY, -- UUID AS TEXT
		emailId TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Failed to create users table")
	}

	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Failed to create events table")
	}

}
