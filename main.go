package main

import (
	"encoding/json"
	"fmt"
	"lets-go/database"
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
	if err := database.InitializeDB("./database/database.db", database.DB); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer database.DB.Close()

	http.HandleFunc("/", views.IndexPage)
	http.HandleFunc("/login", views.LoginPage)
	http.HandleFunc("/register", views.RegisterPage)
	http.Handle("/api/v1/auth/login", MethodMiddleware("POST", http.HandlerFunc(login)))
	// http.Handle("/api/v1/auth/register", MethodMiddleware("POST", http.HandlerFunc(register)))
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

func login(w http.ResponseWriter, r *http.Request) {
	var loginData LoginData

	defer r.Body.Close()

	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
	}

	response := ResponseData{
		Status:  "success",
		Message: "success",
		Data: map[string]string{
			"email":    loginData.Email,
			"password": loginData.Password,
		},
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
