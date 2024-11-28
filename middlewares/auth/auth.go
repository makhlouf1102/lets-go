package auth_middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"lets-go/libs/env"
	"lets-go/libs/token"
)

var (
	accessSecret  = env.Get("TOKEN_ACCESS_SECRET")
	refreshSecret = env.Get("TOKEN_REFRESH_SECRET")
	cookieName    = env.Get("REFRESH_HTTP_COOKIE_NAME")
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
		
		accessToken, err := token.Parse(accessTokenStr, accessSecret)
		if err != nil {
			log.Println("error parsing access token:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		if accessToken.Valid {
			t.handler.ServeHTTP(w, r)
			return
		}
	}

	refreshCookie, err := r.Cookie(cookieName)
	if err != nil {
		log.Println("failed to get refresh token cookie:", err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	refreshTokenStr := refreshCookie.Value

	refreshToken, err := token.Parse(refreshTokenStr, refreshSecret)
	if err != nil {
		log.Println("error parsing refresh token:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	if !refreshToken.Valid {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	refreshClaims, err := token.ExtractClaims(refreshToken, []byte(refreshSecret))
	if err != nil {
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	userId, ok := refreshClaims["user_id"].(string)
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
		newRefreshToken, err := token.CreateRefreshToken(userId)
		if err != nil {
			log.Println("error creating new refresh token:", err)
			http.Error(w, "server error", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     cookieName,
			Value:    newRefreshToken,
			Path:     "/",
			MaxAge:   3600 * 24 * 7,
			HttpOnly: true,
			Secure:   true, // Ensure this is conditional for non-HTTPS environments
			SameSite: http.SameSiteLaxMode,
		})
	}

	newAccessToken, err := token.CreateAccessToken(userId)
	if err != nil {
		log.Println("error creating new access token:", err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

	response := ResponseToken{
		Status:      "success",
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
