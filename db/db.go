package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connect() *sql.DB {
	db, err := sql.Open("sqlite3", "tasks.db")
	if err != nil {
		panic(err)
	}
	return db
}

func Init() {
	db := Connect()
	defer db.Close()
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		due_date DATETIME NOT NULL,
		completed BOOLEAN NOT NULL
	)`)
	if err != nil {
		panic(err)
	}
}
