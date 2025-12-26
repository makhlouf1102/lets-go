package problem

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	GetProblem(ctx context.Context, problemID int64) (*Problem, error)
	CreateProblem(ctx context.Context, problem Problem) error
	ListProblems(ctx context.Context) ([]Problem, error)
	ListTests(ctx context.Context, problemID int64) ([]InputOutput, error)
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
	if err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Signature, &p.Difficulty); err != nil {
		return nil, err
	}

	return &p, nil
}

func (ps *ProblemStore) ListTests(ctx context.Context, problemID int64) ([]InputOutput, error) {
	rows, err := ps.db.Query(ctx, "SELECT * FROM tests WHERE problem_id = $1", problemID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tests []InputOutput
	for rows.Next() {
		var t InputOutput
		if err := rows.Scan(nil, nil, &t.Input, &t.Output); err != nil {
			return nil, err
		}
		tests = append(tests, t)
	}
	return tests, nil
}

func (ps *ProblemStore) CreateProblem(ctx context.Context, problem Problem) error {
	_, err := ps.db.Exec(ctx, "INSERT INTO problems (title, description, signature, difficulty) VALUES ($1, $2, $3, $4)", problem.Title, problem.Description, problem.Signature, problem.Difficulty)
	return err
}

func (ps *ProblemStore) ListProblems(ctx context.Context) ([]Problem, error) {
	rows, err := ps.db.Query(ctx, "SELECT * FROM problems")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []Problem
	for rows.Next() {
		var p Problem
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Signature, &p.Difficulty); err != nil {
			return nil, err
		}
		problems = append(problems, p)
	}
	return problems, nil
}
