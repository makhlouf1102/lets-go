package user

import (
	"database/sql"
	"errors"
	"lets-go/database"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // Hide from JSON output
}

func (u *User) Create() error {
	// return an error if the user exists
	
	query := `INSERT INTO user (username, email, password) VALUES (?, ?, ?)`
	result, err := database.DB.Exec(query, u.Username, u.Email, u.Password)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = int(id)
	return nil
}

func Get(id int) (*User, error) {
	query := `SELECT user_id, username, email, password FROM user WHERE user_id = ?`
	row := database.DB.QueryRow(query, id)

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}

func (u *User) Delete() error {
	query := `DELETE FROM user WHERE user_id = ?`
	_, err := database.DB.Exec(query, u.ID)
	return err
}

func (u *User) Update() (*User, error) {
	query := `UPDATE user SET username = ?, email = ?, password = ? WHERE user_id = ?`
	_, err := database.DB.Exec(query, u.Username, u.Email, u.Password, u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
