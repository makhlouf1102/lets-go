package auth

import (
	"encoding/json"
	"net/http"
)

type RegisterRequestData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var registerRequestData RegisterRequestData

	err := json.NewDecoder(r.Body).Decode(&registerRequestData)

	if err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {}
