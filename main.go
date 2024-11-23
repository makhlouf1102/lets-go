package main

import (
	"fmt"
	"lets-go/database"
	"lets-go/handlers/auth"
	"lets-go/libs/env"
	"lets-go/views"
	"log"
	"net/http"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterData struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ResponseData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func main() {
	env.Load()

	if err := database.InitializeDB("./database/database.db"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.DB.Close()

	http.HandleFunc("/", views.IndexPage)
	http.HandleFunc("/login", views.LoginPage)
	http.HandleFunc("/register", views.RegisterPage)
	http.Handle("POST /api/v1/auth/register", http.HandlerFunc(auth.Register))
	http.Handle("POST /api/v1/auth/login", http.HandlerFunc(auth.Login))
	http.Handle("GET /api/v1/ping", http.HandlerFunc(ping))
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Request Middleware
func MethodMiddleware(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Bad request should be "+method, http.StatusMethodNotAllowed)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
}
