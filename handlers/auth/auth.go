package auth

import (
	"encoding/json"
	"lets-go/libs/bcrypt"
	"lets-go/libs/env"
	localconstants "lets-go/libs/localConstants"
	loglib "lets-go/libs/logLib"
	"lets-go/libs/token"
	role_model "lets-go/models/role"
	user_model "lets-go/models/user"
	user_role_model "lets-go/models/user_role"
	localTypes "lets-go/types"
	"net/http"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj localTypes.RegisterRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.HashPassword(dataObj.Password)

	if err != nil {
		loglib.LogError("server error hashing password", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
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
		loglib.LogError("username or email already exists", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	if err := user.Create(); err != nil {
		loglib.LogError("Server Error Creating user", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	role, err := role_model.GetRole("User")

	if err != nil {
		loglib.LogError("server error while getting the roles", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	user_role := &user_role_model.UserRole{
		ID:     uuid.New().String(),
		UserID: user.ID,
		RoleID: role.Name,
	}

	if duplicate, err := user_role.CheckDuplicate(); err != nil || duplicate {
		loglib.LogError("server error while setting up the roles", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	if err := user_role.Create(); err != nil {
		loglib.LogError("server error while creating a role", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	response := localTypes.RegisterResponseData{
		Status:  "success",
		Message: "User successfully registered",
		Data:    nil,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		loglib.LogError("error while encoding response", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj localTypes.LoginRequestData

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
		loglib.LogError("error while setting up the roles", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	accessToken, err := token.CreateAccessToken(user.ID, listRoles)

	if err != nil {
		loglib.LogError("error while creating access token", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	refreshToken, err := token.CreateRefreshToken(user.ID, listRoles)
	if err != nil {
		loglib.LogError("error while creating refresh token", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
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

	response := localTypes.LoginResponseData{
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
		loglib.LogError("error while encoding response", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}
}
