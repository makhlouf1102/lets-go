package role_model

import (
	"lets-go/database"
)

type Role struct {
	ID   string `json:"role_id"`
	Name string `json:"name"`
}

func GetAllRoles() ([]Role, error) {
	rows, err := database.DB.Query("SELECT * FROM role")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func GetRole(id string) (*Role, error) {
	row := database.DB.QueryRow("SELECT * FROM role WHERE role_id = ?", id)
	var role Role
	err := row.Scan(&role.ID, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Role) Create() error {
	query := "INSERT INTO role (role_id, name) VALUES (?, ?)"
	_, err := database.DB.Exec(query, r.ID, r.Name)
	return err
}

func (r *Role) Update() error {
	query := "UPDATE role SET name = ? WHERE role_id = ?"
	_, err := database.DB.Exec(query, r.Name, r.ID)
	return err
}

func (r *Role) Delete() error {
	query := "DELETE FROM role WHERE role_id = ?"
	_, err := database.DB.Exec(query, r.ID)
	return err
}
