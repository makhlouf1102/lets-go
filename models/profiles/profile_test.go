package profile

import (
	"lets-go/database"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./profile_test_data_script.sql")
	defer database.DB.Close()
	code := m.Run()
	os.Exit(code)
}

func cleanupProfiles(t *testing.T) {
	if err := DeleteAll(); err != nil {
		t.Fatalf("Failed to cleanup profiles: %v", err)
	}
}

func TestProfileCreate(t *testing.T) {
	cleanupProfiles(t)

	// Test creating profile for existing user
	profile := &Profile{
		UserID:    1,
		FirstName: "John",
		LastName:  "Doe",
		BirthDate: time.Now(),
		Biography: "Test bio",
	}

	err := profile.Create()
	if err != nil {
		t.Fatalf("Failed to create profile: %v", err)
	}

	if profile.ID == 0 {
		t.Error("Expected profile ID to be set after creation")
	}

	// Test duplicate profile creation
	duplicateProfile := &Profile{
		UserID:    1,
		FirstName: "Another",
		LastName:  "Name",
		BirthDate: time.Now(),
		Biography: "Should fail",
	}

	err = duplicateProfile.Create()
	if err == nil {
		t.Error("Expected error when creating duplicate profile")
	}
}

func TestProfileGet(t *testing.T) {
	cleanupProfiles(t)

	// Create a profile first
	profile := &Profile{
		UserID:    2,
		FirstName: "Jane",
		LastName:  "Smith",
		BirthDate: time.Now(),
		Biography: "Test bio",
	}

	err := profile.Create()
	if err != nil {
		t.Fatalf("Failed to create test profile: %v", err)
	}

	// Test successful retrieval
	retrieved, err := Get(profile.ID)
	if err != nil {
		t.Fatalf("Failed to get profile: %v", err)
	}

	if retrieved.FirstName != profile.FirstName {
		t.Errorf("Expected first name %s, got %s", profile.FirstName, retrieved.FirstName)
	}

	// Test non-existent profile
	_, err = Get(99999)
	if err == nil {
		t.Error("Expected error when getting non-existent profile")
	}
}

func TestProfileUpdate(t *testing.T) {
	cleanupProfiles(t)

	// Create initial profile
	profile := &Profile{
		UserID:    3,
		FirstName: "Update",
		LastName:  "Test",
		BirthDate: time.Now(),
		Biography: "Original bio",
	}

	err := profile.Create()
	if err != nil {
		t.Fatalf("Failed to create test profile: %v", err)
	}

	// Update profile
	profile.FirstName = "Updated"
	profile.Biography = "Updated bio"

	updated, err := profile.Update()
	if err != nil {
		t.Fatalf("Failed to update profile: %v", err)
	}

	// Verify update
	retrieved, err := Get(updated.ID)
	if err != nil {
		t.Fatalf("Failed to get updated profile: %v", err)
	}

	if retrieved.FirstName != "Updated" {
		t.Errorf("Expected first name 'Updated', got '%s'", retrieved.FirstName)
	}
	if retrieved.Biography != "Updated bio" {
		t.Errorf("Expected biography 'Updated bio', got '%s'", retrieved.Biography)
	}
}

func TestProfileDelete(t *testing.T) {
	cleanupProfiles(t)

	// Create a profile to delete
	profile := &Profile{
		UserID:    1,
		FirstName: "Delete",
		LastName:  "Test",
		BirthDate: time.Now(),
		Biography: "To be deleted",
	}

	err := profile.Create()
	if err != nil {
		t.Fatalf("Failed to create test profile: %v", err)
	}

	// Delete profile
	err = profile.Delete()
	if err != nil {
		t.Fatalf("Failed to delete profile: %v", err)
	}

	// Verify deletion
	_, err = Get(profile.ID)
	if err == nil {
		t.Error("Expected error when getting deleted profile")
	}
}
