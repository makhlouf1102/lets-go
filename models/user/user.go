package user

import (
	"fmt"
	"lets-go/database"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"` // Hide from JSON output
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func (u *User) Create() error {
	query := `INSERT INTO user (user_id, username, email, password, first_name, last_name) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := database.DB.Exec(query, u.ID, u.Username, u.Email, u.Password, u.FirstName, u.LastName)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func Get(id string) (*User, error) {
	query := `SELECT user_id, username, email, password, first_name, last_name FROM user WHERE user_id = ?`
	row := database.DB.QueryRow(query, id)

	if err := row.Err(); err != nil {
		return nil, err
	}

	var user User
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.FirstName, &user.LastName); err != nil {
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
	query := `UPDATE user SET username = ?, email = ?, password = ?, first_name = ?, last_name = ? WHERE user_id = ?`
	_, err := database.DB.Exec(query, u.Username, u.Email, u.Password, u.FirstName, u.LastName, u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) CheckDuplicate() (bool, error) {
	query := `SELECT COUNT(*) FROM user WHERE username = ? OR email = ?`
	var count int
	if err := database.DB.QueryRow(query, u.Username, u.Email).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
