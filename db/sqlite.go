package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", os.Getenv("SQLITE3_FILENAME"))

	if err != nil {
		log.Fatal(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        email TEXT NOT NULL UNIQUE,
        password TEXT NOT NULL
    );
    `

	_, err := DB.Exec(createUsersTable)

	// Check for errors
	if err != nil {
		panic("Error creating users table: " + err.Error())
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL,
        description TEXT NOT NULL,
        location TEXT NOT NULL,
        date DATETIME NOT NULL,
        user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
    );
    `

	_, err = DB.Exec(createEventsTable)

	// Check for errors
	if err != nil {
		panic("Error creating events table: " + err.Error())
	}
}
