package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUserTable := `CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
)`

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("could not create user table")
	}

	createEventTable := `CREATE TABLE IF NOT EXISTS event (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255) NOT NULL,
    location VARCHAR(255) NOT NULL,
    date_time DATETIME NOT NULL,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES user (id)
)`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("could not create event table")
	}
}
