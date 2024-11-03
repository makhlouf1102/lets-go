// models/user_test.go
package user

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	// Initialize the in-memory database for testing
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("Failed to open test database: %v", err)
	}
	defer db.Close()

	database.DB = db

	// Load and execute the database setup script
	build_script, err := os.ReadFile("./models/users/user_test_db_script.sql")
	data_script, err := os.ReadFile("./database/scripts/db_build.sql")
	if err != nil {
		log.Fatalf("Failed to read database setup script: %v", err)
	}

	if _, err = db.Exec(string(build_script)); err != nil {
		log.Fatalf("Failed to execute database setup build script: %v", err)
	}

	if _, err = db.Exec(string(data_script)); err != nil {
		log.Fatalf("Failed to execute database setup data script: %v", err)
	}

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestUserCreate(t *testing.T) {
	u := &User{Username: "testuser", Email: "test@example.com", Password: "password123"}
	err := u.Create()
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verify user has an ID after creation
	if u.ID == 0 {
		t.Error("Expected non-zero user ID after creation")
	}
}

func TestUserGet(t *testing.T) {
	u := &User{Username: "testuser", Email: "test@example.com", Password: "password123"}
	u.Create()

	// Retrieve the user
	retrievedUser, err := Get(u.ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	// Verify retrieved user details
	if retrievedUser.Username != u.Username || retrievedUser.Email != u.Email {
		t.Errorf("Retrieved user details do not match: got %+v, want %+v", retrievedUser, u)
	}
}

func TestUserUpdate(t *testing.T) {
	u := &User{Username: "testuser", Email: "test@example.com", Password: "password123"}
	u.Create()

	// Update user information
	u.Username = "updateduser"
	u.Email = "updated@example.com"
	updatedUser, err := u.Update()
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Verify update
	if updatedUser.Username != "updateduser" || updatedUser.Email != "updated@example.com" {
		t.Errorf("User update failed: got %+v", updatedUser)
	}
}

func TestUserDelete(t *testing.T) {
	u := &User{Username: "testuser", Email: "test@example.com", Password: "password123"}
	u.Create()

	// Delete the user
	err := u.Delete()
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify user no longer exists
	_, err = Get(u.ID)
	if err == nil || err.Error() != "user not found" {
		t.Errorf("Expected 'user not found' error, got %v", err)
	}
}
