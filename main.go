package main

import (
	"fmt"
	"lets-go/database"
	"lets-go/handlers/auth"
	"lets-go/libs/env"
	"lets-go/views"
	"log"
	"net/http"
	"lets-go/middlewares/logger"
	"lets-go/middlewares/auth"
)

func main() {
	env.Load()

	if err := database.InitializeDB("./database/database.db"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.DB.Close()

	mux :=  http.NewServeMux()

	mux.HandleFunc("/", views.IndexPage)
	mux.HandleFunc("/login", views.LoginPage)
	mux.HandleFunc("/register", views.RegisterPage)
	mux.Handle("POST /api/v1/auth/register", http.HandlerFunc(auth.Register))
	mux.Handle("POST /api/v1/auth/login", http.HandlerFunc(auth.Login))
	mux.Handle("GET /api/v1/ping", http.HandlerFunc(ping))
	mux.Handle("GET /api/v1/ping-protected", auth_middleware.NewTokenRefresher(http.HandlerFunc(ping)))
	
	wrapperMux := logger_middleware.NewLogger(mux)
	
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wrapperMux))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
