package profile

import (
	"database/sql"
	"errors"
	"lets-go/database"
	"time"
)

// Profile represents a user profile in the system
type Profile struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date"`
	Biography string    `json:"biography"`
}

func (p *Profile) Create() error {
	query := `INSERT INTO profile (user_id, first_name, last_name, date_of_birth, bio) VALUES (?, ?, ?, ?, ?)`
	result, err := database.DB.Exec(query, p.UserID, p.FirstName, p.LastName, p.BirthDate, p.Biography)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	p.ID = int(id)
	return nil
}

func Get(id int) (*Profile, error) {
	query := `SELECT profile_id, user_id, first_name, last_name, date_of_birth, bio FROM profile WHERE profile_id = ?`
	row := database.DB.QueryRow(query, id)

	var profile Profile
	if err := row.Scan(&profile.ID, &profile.UserID, &profile.FirstName, &profile.LastName, &profile.BirthDate, &profile.Biography); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("profile not found")
		}
		return nil, err
	}
	return &profile, nil
}

func (p *Profile) Delete() error {
	query := `DELETE FROM profile WHERE profile_id = ?`
	_, err := database.DB.Exec(query, p.ID)
	return err
}

func (p *Profile) Update() (*Profile, error) {
	query := `UPDATE profile SET first_name = ?, last_name = ?, date_of_birth = ?, bio = ? WHERE profile_id = ?`
	_, err := database.DB.Exec(query, p.FirstName, p.LastName, p.BirthDate, p.Biography, p.ID)
	if err != nil {
		return nil, err
	}
	return p, nil
}
