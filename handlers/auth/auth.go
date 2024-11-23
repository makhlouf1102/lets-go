package auth

import (
	"encoding/json"
	"lets-go/libs/bcrypt"
	"lets-go/libs/token"
	user_model "lets-go/models/user"
	"net/http"

	"github.com/google/uuid"
)

type RegisterRequestData struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type RegisterResponseData struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Data    *user_model.User `json:"data"`
}

type LoginRequestData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseData struct {
	Status       string           `json:"status"`
	Message      string           `json:"message"`
	AccessToken  string           `json:"accessToken"`
	Data         *user_model.User `json:"data"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj RegisterRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.HashPassword(dataObj.Password)
	if err != nil {
		http.Error(w, "server error hashing password", http.StatusInternalServerError)
		return
	}

	user := &user_model.User{
		ID:        uuid.New().String(),
		Username:  dataObj.Username,
		Email:     dataObj.Email,
		Password:  hashedPassword,
		FirstName: dataObj.FirstName,
		LastName:  dataObj.LastName,
	}

	if duplicate, err := user.CheckDuplicate(); err != nil || duplicate {
		http.Error(w, "username or email already exists", http.StatusConflict)
		return
	}

	if err := user.Create(); err != nil {
		http.Error(w, "Server Error Creating User", http.StatusInternalServerError)
		return
	}

	response := RegisterResponseData{
		Status:  "success",
		Message: "User successfully registered",
		Data:    user,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "server error encoding response", http.StatusInternalServerError)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj LoginRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	user, err := user_model.GetByEmail(dataObj.Email)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	if !bcrypt.CheckPasswordHash(dataObj.Password, user.Password) {
		http.Error(w, "invalid password", http.StatusUnauthorized)
		return
	}

	accessToken, err := token.CreateAccessToken(user.ID)
	
	if err != nil {
		http.Error(w, "server error creating access token", http.StatusInternalServerError)
		return
	}

	refreshToken, err := token.CreateRefreshToken(user.ID)
	if err != nil {
		http.Error(w, "server error creating refresh token", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie {
		Name: "refreshToken",
		Value: refreshToken,
		Path:     "/",
        MaxAge:   3600 * 24 * 7,
        HttpOnly: true,
        Secure:   true,
        SameSite: http.SameSiteLaxMode,
	}

	response := LoginResponseData{
		Status:  "success",
		Message: "User successfully logged in",
		AccessToken: accessToken,
		Data:    user,
	}
	http.SetCookie(w, &cookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "server error encoding response", http.StatusInternalServerError)
		return
	}
}
