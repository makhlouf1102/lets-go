package commonFunctions

import (
	"net/http"
)

func CreateSecureCookie(name, value string, maxAge int, httpOnly bool) http.Cookie {
	return http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   maxAge,
		HttpOnly: httpOnly,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	}
}
