package problem

type Problem struct {
	ID          int64
	Name        string
	Description string
	Template    string
	Difficulty  string
}

type TestProblem struct {
	ID        int64
	ProblemID int64
	Input     string
	Output    string
}
