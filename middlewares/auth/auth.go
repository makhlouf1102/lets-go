package auth_middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"lets-go/libs/env"
	"lets-go/libs/token"
)

var (
	refreshSecret = env.Get("TOKEN_REFRESH_SECRET")
)

type TokenRefresher struct {
	handler http.Handler
}

type ResponseToken struct {
	Status      string `json:"status"`
	Message     string `json:"message"`
	AccessToken string `json:"accessToken"`
}

func (t *TokenRefresher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	accessTokenStr := r.Header.Get("Authorization")
	if accessTokenStr != "" {

		accessToken, err := token.Parse(accessTokenStr, env.Get("TOKEN_ACCESS_SECRET"))
		if err != nil {
			log.Println("error parsing access token:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		if accessToken.Valid {

			accessClaims, err := token.ExtractClaims(accessToken, []byte(env.Get("TOKEN_ACCESS_SECRET")))
			if err != nil {
				http.Error(w, "server error", http.StatusInternalServerError)
				return
			}

			userId, ok := accessClaims["user_id"].(string)
			if !ok {
				http.Error(w, "invalid token claims", http.StatusForbidden)
				return
			}

			userRoles, ok := accessClaims["user_roles"].([]string)
			if !ok {
				http.Error(w, "invalid token claims", http.StatusForbidden)
				return
			}

			data := token.TokenClaim{
				UserID:    userId,
				UserRoles: userRoles,
			}

			ctx := context.WithValue(r.Context(), "protected_data", data)

			t.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}

	refreshCookie, err := r.Cookie(env.Get("REFRESH_HTTP_COOKIE_NAME"))
	if err != nil {
		log.Println("failed to get refresh token cookie:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	refreshTokenStr := refreshCookie.Value

	refreshToken, err := token.Parse(refreshTokenStr, env.Get("TOKEN_REFRESH_SECRET"))
	if err != nil {
		log.Println("error parsing refresh token:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if !refreshToken.Valid {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	refreshClaims, err := token.ExtractClaims(refreshToken, []byte(env.Get("TOKEN_REFRESH_SECRET")))
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	userId, ok := refreshClaims["user_id"].(string)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusForbidden)
		return
	}

	userRoles, ok := refreshClaims["user_roles"].([]string)
	if !ok {
		http.Error(w, "invalid token claims", http.StatusForbidden)
		return
	}

	expirationTime, err := refreshClaims.GetExpirationTime()
	if err != nil {
		log.Println("error getting expiration time:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if time.Now().After(expirationTime.Time) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	if time.Now().Add(24 * time.Hour).After(expirationTime.Time) {
		newRefreshToken, err := token.CreateRefreshToken(userId, userRoles)
		if err != nil {
			log.Println("error creating new refresh token:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		newRefreshCookie := http.Cookie{
			Name:     env.Get("REFRESH_HTTP_COOKIE_NAME"),
			Value:    newRefreshToken,
			Path:     "/",
			MaxAge:   refreshCookie.MaxAge,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}

		isAuthCookie := http.Cookie{
			Name:     "has-refresh-token",
			Value:    "true",
			Path:     "/",
			MaxAge:   refreshCookie.MaxAge,
			HttpOnly: false,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, &newRefreshCookie)
		http.SetCookie(w, &isAuthCookie)

	}

	newAccessToken, err := token.CreateAccessToken(userId, userRoles)
	if err != nil {
		log.Println("error creating new access token:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	response := ResponseToken{
		Status:      "New Token",
		Message:     "New token sent to you",
		AccessToken: newAccessToken,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Println("error encoding response:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
	}
}

func NewTokenRefresher(handlerToWrap http.Handler) *TokenRefresher {
	return &TokenRefresher{handler: handlerToWrap}
}
