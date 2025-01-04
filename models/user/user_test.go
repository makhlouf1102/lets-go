// models/user/user_test.go
package user_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./user_test_data_script.sql")

	defer database.DB.Close()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestUserCreate(t *testing.T) {
	// Define a user for testing
	user := &User{
		ID:        "testuser4",
		Username:  "testuser4",
		Email:     "test4@example.com",
		Password:  "password4",
		FirstName: "Jack",
		LastName:  "Daniels",
	}

	// Call the Create method
	err := user.Create()
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Verify the user was inserted
	var retrievedUser User
	query := `SELECT user_id, username, email, password, first_name, last_name FROM user WHERE user_id = ?`
	row := database.DB.QueryRow(query, user.ID)
	err = row.Scan(&retrievedUser.ID, &retrievedUser.Username, &retrievedUser.Email, &retrievedUser.Password, &retrievedUser.FirstName, &retrievedUser.LastName)
	if err != nil {
		t.Fatalf("Failed to retrieve created user: %v", err)
	}

	// Verify the fields
	if user.ID != retrievedUser.ID {
		t.Errorf("Expected ID '%s', got '%s'", user.ID, retrievedUser.ID)
	}
	if user.Username != retrievedUser.Username {
		t.Errorf("Expected Username '%s', got '%s'", user.Username, retrievedUser.Username)
	}
	if user.Email != retrievedUser.Email {
		t.Errorf("Expected Email '%s', got '%s'", user.Email, retrievedUser.Email)
	}
	if user.Password != retrievedUser.Password {
		t.Errorf("Expected Password '%s', got '%s'", user.Password, retrievedUser.Password)
	}
	if user.FirstName != retrievedUser.FirstName {
		t.Errorf("Expected FirstName '%s', got '%s'", user.FirstName, retrievedUser.FirstName)
	}
	if user.LastName != retrievedUser.LastName {
		t.Errorf("Expected LastName '%s', got '%s'", user.LastName, retrievedUser.LastName)
	}
}

func TestUserGet(t *testing.T) {
	// Test getting an existing user
	user, err := Get("testuser1")
	if err != nil {
		t.Fatalf("Failed to get existing user: %v", err)
	}

	// Verify the retrieved data
	if user.Username != "testuser1" ||
		user.Email != "test1@example.com" ||
		user.FirstName != "Test1" ||
		user.LastName != "User1" {
		t.Errorf("Retrieved user data doesn't match expected values")
	}

	// Test getting non-existent user
	_, err = Get("doesnotexist")
	if err == nil {
		t.Error("Expected error when getting non-existent user, got nil")
	}
}

func TestUserUpdate(t *testing.T) {
	// Create a user to update
	user := &User{
		ID:        "testuser2",
		Username:  "updateduser2",
		Email:     "updated2@example.com",
		FirstName: "UpdatedFirst",
		LastName:  "UpdatedLast",
	}

	// Update the user
	updatedUser, err := user.Update()
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Verify the update
	_, err = Get(user.ID)
	if err != nil {
		t.Fatalf("Failed to get updated user: %v", err)
	}

	if updatedUser.Username != user.Username ||
		updatedUser.Email != user.Email ||
		updatedUser.FirstName != user.FirstName ||
		updatedUser.LastName != user.LastName {
		t.Error("Updated user data doesn't match expected values")
	}
}

func TestUserDelete(t *testing.T) {
	// Create a user to delete
	user := &User{ID: "testuser3"}

	// Delete the user
	err := user.Delete()
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify the deletion
	_, err = Get(user.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user, got nil")
	}
}

func TestGetByExistingEmail(t *testing.T) {
	// Prepare a user record in the test database
	expectedUser := &User{
		ID:        "testuseremail1",
		Username:  "emailuser1",
		Email:     "testemail1@example.com",
		Password:  "password123",
		FirstName: "Email",
		LastName:  "User1",
	}
	err := expectedUser.Create()
	if err != nil {
		t.Fatalf("Failed to create test user for email retrieval test: %v", err)
	}

	// Test retrieving a user by existing email
	retrievedUser, err := GetByEmail(expectedUser.Email)
	if err != nil {
		t.Fatalf("Failed to get user by existing email: %v", err)
	}

	// Verify the retrieved data matches the expected data
	if retrievedUser.Email != expectedUser.Email ||
		retrievedUser.Username != expectedUser.Username ||
		retrievedUser.FirstName != expectedUser.FirstName ||
		retrievedUser.LastName != expectedUser.LastName {
		t.Errorf("Retrieved user data doesn't match expected values: got %+v, expected %+v", retrievedUser, expectedUser)
	}
}

func TestGetNonExistingEmail(t *testing.T) {
	// Test retrieving a user by a non-existent email
	_, err := GetByEmail("nonexistent@example.com")
	if err == nil {
		t.Error("Expected error when getting user by non-existing email, got nil")
	}
}
