package roleGuard

import (
	"fmt"
	commonerrors "lets-go/libs/commonErrors"
	localconstants "lets-go/libs/localConstants"
	loglib "lets-go/libs/logLib"
	"net/http"
	"slices"
)

type AdminGuard struct {
	handler http.Handler
}

func (a *AdminGuard) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Admin guard")
	protectedData, ok := r.Context().Value(localconstants.PROTECTED_DATA_KEY).(map[string]interface{})
	if !ok {
		commonerrors.HttpErrorWithMessage(w, nil, http.StatusInternalServerError, "invalid protected data type")
		return
	}

	userRoles, exists := protectedData["user_roles"]
	if !exists {
		commonerrors.HttpErrorWithMessage(w, nil, http.StatusInternalServerError, "userRoles key missing")
		return
	}

	roles, ok := userRoles.([]string)
	if !ok {
		commonerrors.HttpErrorWithMessage(w, nil, http.StatusInternalServerError, "userRoles has invalid type")
		return
	}

	if !slices.Contains(roles, "Admin") {
		commonerrors.HttpErrorWithMessage(w, nil, http.StatusForbidden, "You are not allowed use this feature")
		return
	}
	a.handler.ServeHTTP(w, r)
	loglib.LogError("it is not working", nil)
}

// NewLogger constructs a new Logger middleware handler
func NewAdminGuard(handlerToWrap http.Handler) *AdminGuard {
	return &AdminGuard{handlerToWrap}
}
