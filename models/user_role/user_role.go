package user_role_model

import (
	"database/sql"
	"fmt"
	"lets-go/database"
)

type UserRole struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

func (ur *UserRole) Create() error {
	query := `INSERT INTO user_role (user_role_id, user_id, role_id) VALUES (?, ?, ?)`
	_, err := database.DB.Exec(query, ur.ID, ur.UserID, ur.RoleID)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func Get(id string) (*UserRole, error) {
	query := `SELECT user_role_id, user_id, role_id FROM user_role WHERE user_role_id = ?`
	row := database.DB.QueryRow(query, id)

	var userRole UserRole
	if err := row.Scan(&userRole.ID, &userRole.UserID, &userRole.RoleID); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user role with ID %s not found", id)
		}
		return nil, err
	}
	return &userRole, nil
}

func GetByUserID(userID string) ([]UserRole, error) {
	query := `SELECT user_role_id, user_id, role_id FROM user_role WHERE user_id = ?`
	rows, err := database.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []UserRole
	for rows.Next() {
		var ur UserRole
		if err := rows.Scan(&ur.ID, &ur.UserID, &ur.RoleID); err != nil {
			return nil, err
		}
		userRoles = append(userRoles, ur)
	}

	return userRoles, nil
}

func (ur *UserRole) Delete() error {
	query := `DELETE FROM user_role WHERE user_role_id = ?`
	_, err := database.DB.Exec(query, ur.ID)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRole) Update() (*UserRole, error) {
	query := `UPDATE user_role SET user_id = ?, role_id = ? WHERE user_role_id = ?`
	_, err := database.DB.Exec(query, ur.UserID, ur.RoleID, ur.ID)
	if err != nil {
		return nil, err
	}
	return ur, nil
}

func (ur *UserRole) CheckDuplicate() (bool, error) {
	query := `SELECT COUNT(*) FROM user_role WHERE user_id = ? AND role_id = ?`
	var count int
	if err := database.DB.QueryRow(query, ur.UserID, ur.RoleID).Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
