package problem_model

import (
	"lets-go/database"
)

type Problem struct {
	ID          string `json:"problem_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Difficulty  string `json:"difficulty"`
}

type ProblemList struct {
	Problems []Problem `json:"problems"`
}

func GetAllProblems() ([]ProblemList, error) {
	rows, err := database.DB.Query("SELECT * FROM problem")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var problemLists []ProblemList
	for rows.Next() {
		var problem Problem
		err := rows.Scan(&problem.ID, &problem.Title, &problem.Description, &problem.Difficulty)
		if err != nil {
			return nil, err
		}
		problemLists = append(problemLists, ProblemList{Problems: []Problem{problem}})
	}
	return problemLists, nil
}

func GetProblem(id string) (*Problem, error) {
	row := database.DB.QueryRow("SELECT * FROM problem WHERE problem_id = ?", id)
	var problem Problem
	err := row.Scan(&problem.ID, &problem.Title, &problem.Description, &problem.Difficulty)
	if err != nil {
		return nil, err
	}
	return &problem, nil
}

func (p *Problem) Create() error {
	query := "INSERT INTO problem (problem_id, title, description, difficulty) VALUES (?, ?, ?, ?)"
	_, err := database.DB.Exec(query, p.ID, p.Title, p.Description, p.Difficulty)
	if err != nil {
		return err
	}
	return nil
}

func (p *Problem) Update() error {
	query := "UPDATE problem SET title = ?, description = ?, difficulty = ? WHERE problem_id = ?"
	_, err := database.DB.Exec(query, p.Title, p.Description, p.Difficulty, p.ID)
	return err
}

func (p *Problem) Delete() error {
	query := "DELETE FROM problem WHERE problem_id = ?"
	_, err := database.DB.Exec(query, p.ID)
	return err
}
