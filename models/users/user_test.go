// models/user_test.go
package user

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
	// Test successful creation
	u := &User{Username: "testuser", Email: "test@example.com", Password: "password123"}
	err := u.Create()
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	// Verify user was created
	var count int
	err = database.DB.QueryRow("SELECT COUNT(*) FROM user WHERE username = ?", u.Username).Scan(&count)
	if err != nil {
		t.Fatalf("Failed to check if user was created: %v", err)
	}
	if count != 1 {
		t.Errorf("Expected user count to be 1, got %d", count)
	}

	// Test duplicate username
	duplicateUser := &User{Username: "testuser", Email: "different@example.com", Password: "password123"}
	err = duplicateUser.Create()
	if err == nil {
		t.Error("Expected error when creating user with duplicate username, got nil")
	}
}

func TestUserGet(t *testing.T) {
	// Create a test user first
	u := &User{Username: "gettest", Email: "get@example.com", Password: "password123"}
	err := u.Create()
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Test successful get
	retrieved, err := Get(u.ID)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if retrieved.Username != u.Username {
		t.Errorf("Expected username %s, got %s", u.Username, retrieved.Username)
	}

	// Test non-existent user
	_, err = Get(99999)
	if err == nil {
		t.Error("Expected error when getting non-existent user, got nil")
	}
}

func TestUserUpdate(t *testing.T) {
	// Create a test user first
	u := &User{Username: "updatetest", Email: "update@example.com", Password: "password123"}
	err := u.Create()
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Update user
	u.Username = "updated_username"
	u.Email = "updated@example.com"
	updated, err := u.Update()
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	// Verify update
	retrieved, err := Get(updated.ID)
	if err != nil {
		t.Fatalf("Failed to get updated user: %v", err)
	}
	if retrieved.Username != "updated_username" {
		t.Errorf("Expected username 'updated_username', got '%s'", retrieved.Username)
	}
	if retrieved.Email != "updated@example.com" {
		t.Errorf("Expected email 'updated@example.com', got '%s'", retrieved.Email)
	}
}

func TestUserDelete(t *testing.T) {
	// Create a test user first
	u := &User{Username: "deletetest", Email: "delete@example.com", Password: "password123"}
	err := u.Create()
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}

	// Delete user
	err = u.Delete()
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}

	// Verify deletion
	_, err = Get(u.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user, got nil")
	}
}

func TestUserCreateValidation(t *testing.T) {
	testCases := []struct {
		name    string
		user    User
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Empty Username",
			user:    User{Username: "", Email: "test@example.com", Password: "password123"},
			wantErr: true,
			errMsg:  "username cannot be empty",
		},
		{
			name:    "Empty Email",
			user:    User{Username: "testuser", Email: "", Password: "password123"},
			wantErr: true,
			errMsg:  "email cannot be empty",
		},
		{
			name:    "Empty Password",
			user:    User{Username: "testuser", Email: "test@example.com", Password: ""},
			wantErr: true,
			errMsg:  "password cannot be empty",
		},
		{
			name:    "Invalid Email Format",
			user:    User{Username: "testuser", Email: "invalid-email", Password: "password123"},
			wantErr: true,
			errMsg:  "invalid email format",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Create()
			if tc.wantErr && err == nil {
				t.Errorf("%s: expected error but got none", tc.name)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("%s: unexpected error: %v", tc.name, err)
			}
		})
	}
}
