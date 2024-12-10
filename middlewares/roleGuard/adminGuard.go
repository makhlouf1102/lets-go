package logger_middleware

import (
	"net/http"
	"slices"
)

type AdminGuar struct {
	handler http.Handler
}

func (a *AdminGuar) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	protectedData, ok := r.Context().Value("protected_data").(map[string]interface{})
	if !ok {
		http.Error(w, "server error: invalid protected data type", http.StatusInternalServerError)
		return
	}

	userRoles, exists := protectedData["userRoles"]
	if !exists {
		http.Error(w, "server error: userRoles key missing", http.StatusInternalServerError)
		return
	}

	roles, ok := userRoles.([]string)
	if !ok {
		http.Error(w, "server error: userRoles has invalid type", http.StatusInternalServerError)
		return
	}

	if !slices.Contains(roles, "Admin") {
		http.Error(w, "server error: You are not allowed use this feature", http.StatusForbidden)
		return
	}
	a.handler.ServeHTTP(w, r)
}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *AdminGuar {
	return &AdminGuar{handlerToWrap}
}
