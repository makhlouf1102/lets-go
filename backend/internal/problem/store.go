package problem

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	GetProblem(ctx context.Context, problemID int64) (*Problem, error)
	CreateProblem(ctx context.Context, problem Problem) error
	ListTests(ctx context.Context, problemID int64) ([]TestProblem, error)
}

type ProblemStore struct {
	db *pgxpool.Pool
}

func NewProblemStore(db *pgxpool.Pool) Store {
	return &ProblemStore{db: db}
}

func (ps *ProblemStore) GetProblem(ctx context.Context, problemID int64) (*Problem, error) {
	row := ps.db.QueryRow(ctx, "SELECT * FROM problems WHERE id = $1", problemID)

	var p Problem
	if err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Template, &p.Difficulty); err != nil {
		return nil, err
	}

	return &p, nil
}

func (ps *ProblemStore) ListTests(ctx context.Context, problemID int64) ([]TestProblem, error) {
	rows, err := ps.db.Query(ctx, "SELECT * FROM tests WHERE problem_id = $1", problemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []TestProblem
	for rows.Next() {
		var t TestProblem
		if err := rows.Scan(&t.ID, &t.ProblemID, &t.Input, &t.Output); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}
	return tests, nil
}

func (ps *ProblemStore) CreateProblem(ctx context.Context, problem Problem) error {
	_, err := ps.db.Exec(ctx, "INSERT INTO problems (title, description, template, difficulty) VALUES ($1, $2, $3, $4)", problem.Title, problem.Description, problem.Template, problem.Difficulty)
	return err
}

