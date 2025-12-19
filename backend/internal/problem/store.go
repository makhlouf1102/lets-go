package problem

import (
	"context"
	"database/sql"
)

type Store interface {
	GetProblem(ctx context.Context, problemID int64) (*Problem, error)
	// ListProblems(ctx context.Context) ([]Problem, error)
	ListTests(ctx context.Context, problemID int64) ([]TestProblem, error)
}

type ProblemStore struct {
	db *sql.DB
}

func NewProblemStore(db *sql.DB) Store {
	return &ProblemStore{db: db}
}

func (ps *ProblemStore) GetProblem(ctx context.Context, problemID int64) (*Problem, error) {
	row := ps.db.QueryRowContext(ctx, "SELECT * FROM problems WHERE id = $1", problemID)

	var p Problem
	if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Template, &p.Difficulty); err != nil {
		return nil, err
	}

	return &p, nil
}

func (ps *ProblemStore) ListTests(ctx context.Context, problemID int64) ([]TestProblem, error) {
	rows, err := ps.db.QueryContext(ctx, "SELECT * FROM tests WHERE problem_id = $1", problemID)
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
