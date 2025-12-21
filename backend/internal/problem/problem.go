package problem

type Problem struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Signature   string `json:"signature"`
	Difficulty  string `json:"difficulty"`
}

type TestProblem struct {
	ID        int64  `json:"id"`
	ProblemID int64  `json:"problem_id"`
	Input     string `json:"input"`
	Output    string `json:"output"`
}
