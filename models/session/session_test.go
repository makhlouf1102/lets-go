package session_model

import (
	"os"
	"testing"

	"lets-go/database"

	_ "github.com/mattn/go-sqlite3"
)

func TestMain(m *testing.M) {
	// Set up the test database
	database.SetUpDBForTests("../../database/scripts/db_build.sql", "./session_test_data_script.sql")

	defer database.DB.Close()

	// Run tests
	code := m.Run()
	os.Exit(code)
}

func TestSessionCreate(t *testing.T) {
	session := &Session{
		ID:           "testsession1",
		UserID:       "testuser1",
		RefreshToken: "refreshtoken1",
	}

	err := session.Create()
	if err != nil {
		t.Fatalf("Failed to create session: %v", err)
	}

	// Verify the session was inserted
	var retrievedSession Session
	query := `SELECT session_id, user_id, refreshToken FROM session WHERE session_id = ?`
	row := database.DB.QueryRow(query, session.ID)
	err = row.Scan(&retrievedSession.ID, &retrievedSession.UserID, &retrievedSession.RefreshToken)
	if err != nil {
		t.Fatalf("Failed to retrieve created session: %v", err)
	}

	if session.ID != retrievedSession.ID {
		t.Errorf("Expected session ID '%s', got '%s'", session.ID, retrievedSession.ID)
	}
	if session.UserID != retrievedSession.UserID {
		t.Errorf("Expected UserID '%s', got '%s'", session.UserID, retrievedSession.UserID)
	}
	if session.RefreshToken != retrievedSession.RefreshToken {
		t.Errorf("Expected RefreshToken '%s', got '%s'", session.RefreshToken, retrievedSession.RefreshToken)
	}
}

func TestSessionGet(t *testing.T) {
	session, err := Get("testsession1")
	if err != nil {
		t.Fatalf("Failed to get session: %v", err)
	}

	// Verify session data
	if session.ID != "testsession1" || session.UserID != "testuser1" || session.RefreshToken != "refreshtoken1" {
		t.Errorf("Retrieved session data doesn't match expected values")
	}

	// Test getting non-existent session
	_, err = Get("nonexistentsession")
	if err == nil {
		t.Error("Expected error when getting non-existent session, got nil")
	}
}

func TestSessionUpdate(t *testing.T) {
	session := &Session{
		ID:           "testsession1",
		RefreshToken: "newrefreshtoken1",
	}

	updatedSession, err := session.Update()
	if err != nil {
		t.Fatalf("Failed to update session: %v", err)
	}

	if updatedSession.RefreshToken != "newrefreshtoken1" {
		t.Errorf("Expected updated RefreshToken 'newrefreshtoken1', got '%s'", updatedSession.RefreshToken)
	}
}

func TestSessionDelete(t *testing.T) {
	session := &Session{ID: "testsession1"}

	err := session.Delete()
	if err != nil {
		t.Fatalf("Failed to delete session: %v", err)
	}

	_, err = Get("testsession1")
	if err == nil {
		t.Error("Expected error when getting deleted session, got nil")
	}
}

func TestSessionCheckDuplicate(t *testing.T) {
	session := &Session{RefreshToken: "refreshtoken2"}

	// Duplicate exists
	isDuplicate, err := session.CheckDuplicate()
	if err != nil {
		t.Fatalf("Failed to check duplicate: %v", err)
	}
	if !isDuplicate {
		t.Errorf("Expected duplicate check to return true, got false")
	}

	// No duplicate
	session.RefreshToken = "uniquerefreshtoken"
	isDuplicate, err = session.CheckDuplicate()
	if err != nil {
		t.Fatalf("Failed to check duplicate: %v", err)
	}
	if isDuplicate {
		t.Errorf("Expected duplicate check to return false, got true")
	}
}
