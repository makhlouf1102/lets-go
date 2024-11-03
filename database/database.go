package database

import(
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

func ConnectToDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)

	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Successfully connected to the database!")
    return db, nil
}