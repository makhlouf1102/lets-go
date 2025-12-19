package problem

import (
	"context"
	"database/sql"
)

type Store interface {
	GetProblem(ctx context.Context, problemID int64) (Problem, error)
	GetTests(ctx context.Context, problemID int64) ([]TestProblem, error)
}

type ProblemStore struct {
	db *sql.DB
}

func NewProblemStore(db *sql.DB) Store {
	return &ProblemStore{db: db}
}

func (ps *ProblemStore) GetProblem(ctx context.Context, problemID int64) (Problem, error) {
	return Problem{}, nil
}

func (ps *ProblemStore) GetTests(ctx context.Context, problemID int64) ([]TestProblem, error) {
	return []TestProblem{}, nil
}
