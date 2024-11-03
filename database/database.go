// database/connection.go
package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func InitializeDB(dataSourceName string, dbRef *sql.DB) error {
	var err error
	dbRef, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}

	// Test the connection
	if err := dbRef.Ping(); err != nil {
		return err
	}

	log.Println("Database connection successfully established")
	return nil
}
