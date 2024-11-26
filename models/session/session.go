package session_model

import (
	"fmt"
	"lets-go/database"
)

type Session struct {
	ID     string `json:"session_id"`
	UserID        string `json:"user_id"`
	RefreshToken  string `json:"refresh_token"`
}

func (s *Session) Create() error {
	query := `INSERT INTO session (session_id, user_id, refreshToken) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, s.ID, s.UserID, s.RefreshToken)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Get(sessionID string) (*Session, error) {
	query := `SELECT session_id, user_id, refreshToken FROM session WHERE session_id = ?`
	row := database.DB.QueryRow(query, sessionID)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var session Session
	if err := row.Scan(&session.ID, &session.UserID, &session.RefreshToken); err != nil {
		return nil, err
	}
	return &session, nil
}

// GetByUserID retrieves all sessions for a specific user.
func GetByUserID(userID string) ([]*Session, error) {
	query := `SELECT session_id, user_id, refreshToken FROM session WHERE user_id = ?`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sessions []*Session
	for rows.Next() {
		var session Session
		if err := rows.Scan(&session.ID, &session.UserID, &session.RefreshToken); err != nil {
			return nil, err
		}
		sessions = append(sessions, &session)
	}

	return sessions, nil
}

// Delete removes a session from the database by its session ID.
func (s *Session) Delete() error {
	query := `DELETE FROM session WHERE session_id = ?`
	_, err := database.DB.Exec(query, s.ID)
	return err
}

// Update refreshes the refreshToken of a session.
func (s *Session) Update() (*Session, error) {
	query := `UPDATE session SET refreshToken = ? WHERE session_id = ?`
	_, err := database.DB.Exec(query, s.RefreshToken, s.ID)
	if err != nil {
		return nil, err
	}
	return s, nil
}

// CheckDuplicate checks if a session exists for a specific refreshToken.
func (s *Session) CheckDuplicate() (bool, error) {
	query := `SELECT COUNT(*) FROM session WHERE refreshToken = ?`
	var count int
	if err := database.DB.QueryRow(query, s.RefreshToken).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
