package user_role_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./user_role_test_data_script.sql")

	defer database.DB.Close()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestUserRoleCreate(t *testing.T) {
	userRole := &UserRole{
		ID:     "testuserrole4",
		UserID: "testuser3",
		RoleID: "testrole2",
	}

	err := userRole.Create()
	if err != nil {
		t.Fatalf("Failed to create user-role: %v", err)
	}

	var retrievedUserRole UserRole
	query := `SELECT user_role_id, user_id, role_id FROM user_role WHERE user_role_id = ?`
	row := database.DB.QueryRow(query, userRole.ID)
	err = row.Scan(&retrievedUserRole.ID, &retrievedUserRole.UserID, &retrievedUserRole.RoleID)
	if err != nil {
		t.Fatalf("Failed to retrieve created user-role: %v", err)
	}

	if userRole.ID != retrievedUserRole.ID {
		t.Errorf("Expected ID '%s', got '%s'", userRole.ID, retrievedUserRole.ID)
	}
	if userRole.UserID != retrievedUserRole.UserID {
		t.Errorf("Expected UserID '%s', got '%s'", userRole.UserID, retrievedUserRole.UserID)
	}
	if userRole.RoleID != retrievedUserRole.RoleID {
		t.Errorf("Expected RoleID '%s', got '%s'", userRole.RoleID, retrievedUserRole.RoleID)
	}
}

func TestUserRoleGet(t *testing.T) {
	userRole, err := Get("testuserrole1")
	if err != nil {
		t.Fatalf("Failed to get user-role: %v", err)
	}

	if userRole.UserID != "testuser1" || userRole.RoleID != "testrole1" {
		t.Errorf("Retrieved user-role data doesn't match expected values")
	}

	_, err = Get("doesnotexist")
	if err == nil {
		t.Error("Expected error when getting non-existent user-role, got nil")
	}
}

func TestUserRoleGetByUserID(t *testing.T) {
	userRoles, err := GetByUserID("testuser1")
	if err != nil {
		t.Fatalf("Failed to get user-roles by user ID: %v", err)
	}

	if len(userRoles) == 0 {
		t.Error("Expected at least one user-role, got none")
	}

	firstRole := userRoles[0]
	if firstRole.UserID != "testuser1" {
		t.Errorf("Expected UserID 'testuser1', got '%s'", firstRole.UserID)
	}
}

func TestUserRoleUpdate(t *testing.T) {
	userRole := &UserRole{
		ID:     "testuserrole2",
		UserID: "updateduser2",
		RoleID: "updatedrole2",
	}

	updatedRole, err := userRole.Update()
	if err != nil {
		t.Fatalf("Failed to update user-role: %v", err)
	}

	if updatedRole.UserID != userRole.UserID {
		t.Errorf("Expected UserID '%s', got '%s'", userRole.UserID, updatedRole.UserID)
	}
	if updatedRole.RoleID != userRole.RoleID {
		t.Errorf("Expected RoleID '%s', got '%s'", userRole.RoleID, updatedRole.RoleID)
	}
}

func TestUserRoleDelete(t *testing.T) {
	userRole := &UserRole{ID: "testuserrole3"}

	err := userRole.Delete()
	if err != nil {
		t.Fatalf("Failed to delete user-role: %v", err)
	}

	_, err = Get(userRole.ID)
	if err == nil {
		t.Error("Expected error when getting deleted user-role, got nil")
	}
}

func TestUserRoleCheckDuplicate(t *testing.T) {
	userRole := &UserRole{
		UserID: "testuser1",
		RoleID: "testrole1",
	}

	isDuplicate, err := userRole.CheckDuplicate()
	if err != nil {
		t.Fatalf("Failed to check for duplicates: %v", err)
	}

	if !isDuplicate {
		t.Error("Expected duplicate to be true, got false")
	}

	userRole.UserID = "newuser"
	userRole.RoleID = "newrole"

	isDuplicate, err = userRole.CheckDuplicate()
	if err != nil {
		t.Fatalf("Failed to check for duplicates: %v", err)
	}

	if isDuplicate {
		t.Error("Expected duplicate to be false, got true")
	}
}
