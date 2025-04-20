package role_model

import (
	"lets-go/database"
)

type Role struct {
	Name string `json:"name"`
}

func GetAllRoles() ([]Role, error) {
	rows, err := database.DB.Query("SELECT name FROM role")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role
	for rows.Next() {
		var role Role
		err := rows.Scan(&role.Name)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func GetRole(name string) (*Role, error) {
	row := database.DB.QueryRow("SELECT name FROM role WHERE name = ?", name)
	var role Role
	err := row.Scan(&role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *Role) Create() error {
	query := "INSERT INTO role (name) VALUES (?)"
	_, err := database.DB.Exec(query, r.Name)
	return err
}

func (r *Role) Update(newName string) error {
	query := "UPDATE role SET name = ? WHERE name = ?"
	_, err := database.DB.Exec(query, newName, r.Name)
	r.Name = newName // Update local struct after DB update
	return err
}

func (r *Role) Delete() error {
	query := "DELETE FROM role WHERE name = ?"
	_, err := database.DB.Exec(query, r.Name)
	return err
}
