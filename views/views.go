package views

import (
	// "html/template"
	"net/http"
	"path/filepath"
)

const (
	viewsDir = "views/templates"
)

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
	http.ServeFile(w, r, filepath.Join(viewsDir, filename))
}
