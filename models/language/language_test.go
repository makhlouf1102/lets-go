package language_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	// Set up the test database and seed data
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./language_test_data_script.sql")

	defer database.DB.Close()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestGetAllLanguages(t *testing.T) {
	languages, err := GetAllLanguages()
	if err != nil {
		t.Fatalf("Failed to get all languages: %v", err)
	}

	if len(languages) != 3 {
		t.Errorf("Expected 3 languages, got %d", len(languages))
	}
}

func TestGetLanguage(t *testing.T) {
	language, err := GetLanguage("uuid1")
	if err != nil {
		t.Fatalf("Failed to get language: %v", err)
	}

	if language.Name != "Golang" {
		t.Errorf("Expected name to be 'Golang', got %s", language.Name)
	}
}

func TestCreateLanguage(t *testing.T) {
	language := &Language{
		ID:   "uuid4",
		Name: "Python",
	}

	err := language.Create()
	if err != nil {
		t.Fatalf("Failed to create language: %v", err)
	}

	createdLanguage, err := GetLanguage("uuid4")
	if err != nil {
		t.Fatalf("Failed to get created language: %v", err)
	}

	if createdLanguage.Name != "Python" {
		t.Errorf("Expected name to be 'Python', got %s", createdLanguage.Name)
	}
}

func TestUpdateLanguage(t *testing.T) {
	language := &Language{
		ID:   "uuid1",
		Name: "UpdatedGolang",
	}

	err := language.Update()
	if err != nil {
		t.Fatalf("Failed to update language: %v", err)
	}

	updatedLanguage, err := GetLanguage("uuid1")
	if err != nil {
		t.Fatalf("Failed to get updated language: %v", err)
	}

	if updatedLanguage.Name != "UpdatedGolang" {
		t.Errorf("Expected name to be 'UpdatedGolang', got %s", updatedLanguage.Name)
	}
}

func TestDeleteLanguage(t *testing.T) {
	languages, err := GetAllLanguages()
	if err != nil {
		t.Fatalf("Failed to get all languages: %v", err)
	}

	language, err := GetLanguage("uuid1")
	if err != nil {
		t.Fatalf("Failed to get language: %v", err)
	}

	err = language.Delete()
	if err != nil {
		t.Fatalf("Failed to delete language: %v", err)
	}

	newListLanguages, err := GetAllLanguages()
	if err != nil {
		t.Fatalf("Failed to get all languages: %v", err)
	}

	if len(newListLanguages) == len(languages) {
		t.Errorf("Expected length to be %d, got %d", len(languages)-1, len(newListLanguages))
	}
}
