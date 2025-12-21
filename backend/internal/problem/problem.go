package problem

type Problem struct {
	ID          int64
	Title       string
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
