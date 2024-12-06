package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	viewsDir = "views/templates"
)

var tmpl = template.Must(template.ParseFiles(
	filepath.Join(viewsDir, "header.html"),
	filepath.Join(viewsDir, "footer.html"),
	filepath.Join(viewsDir, "index.html"),
	filepath.Join(viewsDir, "login.html"),
	filepath.Join(viewsDir, "register.html"),
	filepath.Join(viewsDir, "style.html"),
))

func IndexPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "index.html")
}

func LoginPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "login.html")
}

func RegisterPage(w http.ResponseWriter, r *http.Request) {
	RenderStaticPage(w, r, "register.html")
}


func RenderStaticPage(w http.ResponseWriter, r *http.Request, filename string) {
	if err := tmpl.ExecuteTemplate(w, filename, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
