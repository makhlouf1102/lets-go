package auth

import (
	"encoding/json"
	"errors"
	"fmt"
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
	"reflect"

	"github.com/google/uuid"
)

type Model interface {
	Create() error
	Delete() error
	CheckDuplicate() (bool, error)
}

func createModel(model Model) error {
	duplicate, err := model.CheckDuplicate()
	modelName := reflect.ValueOf(model).Type().String()
	var formated string
	if err != nil {
		formated = fmt.Sprintf("an error occured while checking for duplicates of %s", modelName)
		loglib.LogError(formated, err)
		return err
	}

	if duplicate {
		formated = fmt.Sprintf("the %s already exists", modelName)
		err = errors.New(formated)
		loglib.LogError(formated, err)
		return err
	}

	if err := model.Create(); err != nil {
		formated = fmt.Sprintf("an error occured while creating %s", modelName)
		loglib.LogError(formated, err)
		return err
	}

	return nil
}

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

	if err := createModel(user); err != nil {
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

	userRole := &user_role_model.UserRole{
		ID:     uuid.New().String(),
		UserID: user.ID,
		RoleID: role.Name,
	}

	if err := createModel(userRole); err != nil {
		loglib.LogError("Server Error Creating userRole", err)
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

func validateUser(dataObj *localTypes.LoginRequestData) (*user_model.User, error) {
	user, err := user_model.GetByEmail(dataObj.Email)
	if err != nil {
		loglib.LogError("the e-mail doesn't exist", err)
		return nil, err
	}

	if !bcrypt.CheckPasswordHash(dataObj.Password, user.Password) {
		loglib.LogError("wrong password", nil)
		return nil, errors.New("wrong password")
	}

	return user, nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj localTypes.LoginRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		http.Error(w, "invalid Json format", http.StatusBadRequest)
		return
	}

	user, err := validateUser(&dataObj)

	if err != nil {
		loglib.LogError("invalid login", err)
		http.Error(w, localconstants.UNAUTHORIZED, http.StatusUnauthorized)
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
