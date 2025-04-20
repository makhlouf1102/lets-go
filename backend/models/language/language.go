package language_model

import (
	"lets-go/database"
)

type Language struct {
	ID   string `json:"language_id"`
	Name string `json:"name"`
}

func GetAllLanguages() ([]Language, error) {
	rows, err := database.DB.Query("SELECT * FROM language")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []Language
	for rows.Next() {
		var language Language
		err := rows.Scan(&language.ID, &language.Name)
		if err != nil {
			return nil, err
		}
		languages = append(languages, language)
	}
	return languages, nil
}

func GetLanguage(id string) (*Language, error) {
	row := database.DB.QueryRow("SELECT * FROM language WHERE language_id = ?", id)
	var language Language
	err := row.Scan(&language.ID, &language.Name)
	if err != nil {
		return nil, err
	}
	return &language, nil
}

func (l *Language) Create() error {
	query := "INSERT INTO language (language_id, name) VALUES (?, ?)"
	_, err := database.DB.Exec(query, l.ID, l.Name)
	if err != nil {
		return err
	}
	return nil
}

func (l *Language) Update() error {
	query := "UPDATE language SET name = ? WHERE language_id = ?"
	_, err := database.DB.Exec(query, l.Name, l.ID)
	return err
}

func (l *Language) Delete() error {
	query := "DELETE FROM language WHERE language_id = ?"
	_, err := database.DB.Exec(query, l.ID)
	return err
}
