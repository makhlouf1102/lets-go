package authHandler

import (
	"encoding/json"
	"lets-go/libs/bcrypt"
	commonerrors "lets-go/libs/commonErrors"
	localconstants "lets-go/libs/localConstants"
	loglib "lets-go/libs/logLib"
	role_model "lets-go/models/role"
	userModel "lets-go/models/user"
	user_role_model "lets-go/models/user_role"
	localTypes "lets-go/types"
	"net/http"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj localTypes.RegisterRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		commonerrors.DencodingError(w, err)
		return
	}

	hashedPassword, err := bcrypt.HashPassword(dataObj.Password)

	if err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "server error hashing password")
		return
	}

	user := &userModel.User{
		ID:        uuid.New().String(),
		Username:  dataObj.Username,
		Email:     dataObj.Email,
		Password:  hashedPassword,
		FirstName: dataObj.FirstName,
		LastName:  dataObj.LastName,
	}

	if err := createModel(user); err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "Server Error Creating user")
		return
	}

	role, err := role_model.GetRole("User")

	if err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "server error while getting the roles")
		return
	}

	userRole := &user_role_model.UserRole{
		ID:     uuid.New().String(),
		UserID: user.ID,
		RoleID: role.Name,
	}

	if err := createModel(userRole); err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "server error creating a user role in DB")
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
		commonerrors.EncodingError(w, err)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var dataObj localTypes.LoginRequestData

	if err := json.NewDecoder(r.Body).Decode(&dataObj); err != nil {
		commonerrors.DencodingError(w, err)
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

	tokens, err := generateTokens(user.ID, listRoles)

	if err != nil {
		loglib.LogError("error while creating tokens", err)
		http.Error(w, localconstants.SERVER_ERROR, http.StatusInternalServerError)
		return
	}

	expiresIn := 3600 * 24 * 7

	cookies := generateCookies(*tokens, expiresIn)

	response := localTypes.LoginResponseData{
		Status:      "success",
		Message:     "User successfully logged in",
		AccessToken: tokens.accessToken,
		Data:        user,
	}
	http.SetCookie(w, &cookies.RefreshCookie)
	http.SetCookie(w, &cookies.IsAuthCookie)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		commonerrors.EncodingError(w, err)
		return
	}
}
