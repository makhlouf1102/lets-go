package auth

import (
	"encoding/json"
	"lets-go/user"
	"net/http"
	"os/user"
	"github.com/google/uuid"

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

	var dataObj RegisterRequestData


	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	user := &user.User{
		ID: uuid.New().String(),
		Username:  dataObj.Username,
		Email:     dataObj.Email,
		Password:  dataObj.Password,
		FirstName: dataObj.FirstName,
		LastName:  dataObj.LastName,
	}

	
	if err := user.Create(); err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	
	
}

func Login(w http.ResponseWriter, r *http.Request) {}
