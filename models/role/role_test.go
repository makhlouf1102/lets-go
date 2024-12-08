package role_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./role_test_data_script.sql")

	defer database.DB.Close()

	code := m.Run()
	os.Exit(code)
}

func TestGetAllRoles(t *testing.T) {
	roles, err := GetAllRoles()
	if err != nil {
		t.Fatalf("Failed to get all roles: %v", err)
	}

	if len(roles) != 3 {
		t.Errorf("Expected 3 roles, got %d", len(roles))
	}
}

func TestGetRole(t *testing.T) {
	role, err := GetRole("Admin")
	if err != nil {
		t.Fatalf("Failed to get role: %v", err)
	}

	if role.Name != "Admin" {
		t.Errorf("Expected name to be 'Admin', got %s", role.Name)
	}
}

func TestCreateRole(t *testing.T) {
	role := &Role{
		Name: "Tester",
	}

	err := role.Create()
	if err != nil {
		t.Fatalf("Failed to create role: %v", err)
	}

	createdRole, err := GetRole("Tester")
	if err != nil {
		t.Fatalf("Failed to get created role: %v", err)
	}

	if createdRole.Name != "Tester" {
		t.Errorf("Expected name to be 'Tester', got %s", createdRole.Name)
	}
}

func TestUpdateRole(t *testing.T) {
	role := &Role{
		Name: "Admin",
	}

	err := role.Update("SuperAdmin")
	if err != nil {
		t.Fatalf("Failed to update role: %v", err)
	}

	updatedRole, err := GetRole("SuperAdmin")
	if err != nil {
		t.Fatalf("Failed to get updated role: %v", err)
	}

	if updatedRole.Name != "SuperAdmin" {
		t.Errorf("Expected name to be 'SuperAdmin', got %s", updatedRole.Name)
	}

	
}

func TestDeleteRole(t *testing.T) {
	roles, err := GetAllRoles()
	if err != nil {
		t.Fatalf("Failed to get all roles: %v", err)
	}

	role, err := GetRole("SuperAdmin")
	if err != nil {
		t.Fatalf("Failed to get role: %v", err)
	}

	err = role.Delete()
	if err != nil {
		t.Fatalf("Failed to delete role: %v", err)
	}

	newListRoles, err := GetAllRoles()
	if err != nil {
		t.Fatalf("Failed to get all roles: %v", err)
	}

	if len(newListRoles) == len(roles) {
		t.Errorf("Expected length to be %d, got %d", len(roles)-1, len(newListRoles))
	}
}
