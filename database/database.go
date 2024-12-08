package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

var DB *sql.DB

func InitializeDB(dataSourceName string) error {
	// Check if the database file exists
	if _, err := os.Stat(dataSourceName); os.IsNotExist(err) {
		log.Println("Database does not exist, initializing...")

		DB, err = sql.Open("sqlite3", dataSourceName)
		if err != nil {
			return err
		}

		buildScript, err := os.ReadFile("./database/scripts/db_build.sql")
		if err != nil {
			return err
		}

		productionInsertsScript, err := os.ReadFile("./database/scripts/production_inserts.sql")
		if err != nil {
			return err
		}
		buildScript = append(buildScript, productionInsertsScript...)

		if _, err = DB.Exec(string(buildScript)); err != nil {
			return err
		}
	} else {
		// Open the existing database
		DB, err = sql.Open("sqlite3", dataSourceName)
		if err != nil {
			return err
		}
	}

	// Test the connection
	if err := DB.Ping(); err != nil {
		return err
	}

	log.Println("Database connection successfully established")
	return nil
}

func SetUpDBForTests(buildScriptPath string, dataScriptPath string) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}
	DB = db

	build_script, err := os.ReadFile(buildScriptPath)
	if err != nil {
		log.Fatalf("Failed to read database setup script: %v", err)
	}

	data_script, err := os.ReadFile(dataScriptPath)
	if err != nil {
		log.Fatalf("Failed to read database setup script: %v", err)
	}

	if _, err = db.Exec(string(build_script)); err != nil {
		log.Fatalf("Failed to execute database setup build script: %v", err)
	}

	if _, err = db.Exec(string(data_script)); err != nil {
		log.Fatalf("Failed to execute database setup data script: %v", err)
	}
}
