package main

import (
	"fmt"
	"lets-go/database"
	authHandler "lets-go/handlers/auth"
	"lets-go/handlers/problem"
	"lets-go/libs/dockerController"
	"lets-go/libs/env"
	auth_middleware "lets-go/middlewares/auth"
	logger_middleware "lets-go/middlewares/logger"
	"lets-go/middlewares/roleGuard"
	"lets-go/views"
	"log"
	"net/http"
)

func main() {
	env.Load()

	if err := dockerController.InitContainers(); err != nil {
		log.Fatalf("Failed to run the containers: %v", err)
	}

	if err := database.InitializeDB("./database/database-test.db"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.DB.Close()

	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))
	mux.HandleFunc("/", views.IndexPage)
	mux.HandleFunc("GET /login", views.LoginPage)
	mux.HandleFunc("GET /register", views.RegisterPage)
	mux.HandleFunc("GET /problems", views.ProblemsPage)
	mux.HandleFunc("GET /solve-problem/{problemID}", views.SolveProblemsPage)
	mux.Handle("POST /api/v1/auth/register", http.HandlerFunc(authHandler.Register))
	mux.Handle("POST /api/v1/auth/login", http.HandlerFunc(authHandler.Login))
	mux.Handle("GET /api/v1/ping", http.HandlerFunc(ping))
	mux.Handle("GET /api/v1/problem/{programmingLanguage}/{problemID}", http.HandlerFunc(problem.GetProblemCode))
	mux.Handle("GET /api/v1/ping-protected", auth_middleware.NewTokenRefresher(
		roleGuard.NewAdminGuard(
			http.HandlerFunc(ping),
		),
	))
	mux.Handle("POST /api/v1/problem/runcode", http.HandlerFunc(problem.RunCode))

	wrapperMux := logger_middleware.NewLogger(mux)

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wrapperMux))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
