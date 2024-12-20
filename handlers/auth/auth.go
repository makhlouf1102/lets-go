package auth

import (
	"encoding/json"
	"lets-go/libs/bcrypt"
	"lets-go/libs/env"
	"lets-go/libs/token"
	role_model "lets-go/models/role"
	user_model "lets-go/models/user"
	user_role_model "lets-go/models/user_role"
	"log"
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
	Status      string           `json:"status"`
	Message     string           `json:"message"`
	AccessToken string           `json:"accessToken"`
	Data        *user_model.User `json:"data"`
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

	role, err := role_model.GetRole("User")

	if err != nil {
		http.Error(w, "server error while getting the roles", http.StatusInternalServerError)
		return
	}

	user_role := &user_role_model.UserRole{
		ID:     uuid.New().String(),
		UserID: user.ID,
		RoleID: role.Name,
	}

	if duplicate, err := user_role.CheckDuplicate(); err != nil || duplicate {
		http.Error(w, "server error while setting up the roles", http.StatusInternalServerError)
		return
	}

	if err := user_role.Create(); err != nil {
		http.Error(w, "Server Error roles", http.StatusInternalServerError)
		return
	}

	response := RegisterResponseData{
		Status:  "success",
		Message: "User successfully registered",
		Data:    nil,
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
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	if !bcrypt.CheckPasswordHash(dataObj.Password, user.Password) {
		http.Error(w, "invalid login", http.StatusUnauthorized)
		return
	}

	user_roles, err := user_role_model.GetByUserID(user.ID)

	listRoles := make([]string, len(user_roles))

	for i, role := range user_roles {
		listRoles[i] = role.RoleID
	}

	if err != nil {
		log.Fatal("server error setting up roles")
		http.Error(w, "server error while getting the user role", http.StatusInternalServerError)
		return
	}

	accessToken, err := token.CreateAccessToken(user.ID, listRoles)

	if err != nil {
		log.Fatal("server error creating access token")
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	refreshToken, err := token.CreateRefreshToken(user.ID, listRoles)
	if err != nil {
		log.Fatal("server error creating refresh token")
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	cookieMaxAge := 3600 * 24 * 7

	refreshCookie := http.Cookie{
		Name:     env.Get("REFRESH_HTTP_COOKIE_NAME"),
		Value:    refreshToken,
		Path:     "/",
		MaxAge:   cookieMaxAge,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	isAuthCookie := http.Cookie{
		Name:     "has-refresh-token",
		Value:    "true",
		Path:     "/",
		MaxAge:   cookieMaxAge,
		HttpOnly: false,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}

	response := LoginResponseData{
		Status:      "success",
		Message:     "User successfully logged in",
		AccessToken: accessToken,
		Data:        user,
	}
	http.SetCookie(w, &refreshCookie)
	http.SetCookie(w, &isAuthCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Fatal("server error encoding response")
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}
