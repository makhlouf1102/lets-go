package views

import (
	"html/template"
	problem_model "lets-go/models/problem"
	"log"
	"net/http"
	"path/filepath"
)

const (
	viewsDir = "views/templates"
)

var tmpl, err = template.ParseFiles(
	filepath.Join(viewsDir, "header.html"),
	filepath.Join(viewsDir, "footer.html"),
	filepath.Join(viewsDir, "index.html"),
	filepath.Join(viewsDir, "login.html"),
	filepath.Join(viewsDir, "register.html"),
	filepath.Join(viewsDir, "style.html"),
	filepath.Join(viewsDir, "problems.html"),
	filepath.Join(viewsDir, "solve-problem.html"),
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "index.html", nil)
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "login.html", nil)
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "register.html", nil)
}

func ProblemsPage(w http.ResponseWriter, r *http.Request) {
	problems, err := problem_model.GetAllProblems()
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	RenderStaticPage(w, r, "problems.html", problems)

}

func SolveProblemsPage(w http.ResponseWriter, r *http.Request) {
	problemID := r.PathValue("problemID")
	problem, err := problem_model.GetProblem(problemID)
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
	}

	RenderStaticPage(w, r, "solve-problem.html", problem)

}

func RenderStaticPage(w http.ResponseWriter, r *http.Request, filename string, data interface{}) {
	if err != nil {
		log.Fatalf("Error parsing templates: %v", err)
	}

	if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
