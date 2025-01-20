package auth_middleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	commonerrors "lets-go/libs/commonErrors"
	"lets-go/libs/commonFunctions"
	"lets-go/libs/env"
	localconstants "lets-go/libs/localConstants"
	"lets-go/libs/token"
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
			commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "error parsing access token")
			return
		}

		if accessToken.Valid {

			accessClaims, err := token.ExtractClaims(accessToken, []byte(env.Get("TOKEN_ACCESS_SECRET")))
			if err != nil {
				commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "error extracting claim for access token")
				return
			}

			userId, ok := accessClaims["user_id"].(string)
			if !ok {
				commonerrors.HttpErrorWithMessage(w, err, http.StatusForbidden, "invalid token claim for value user_id")
				return
			}

			userRoles, ok := accessClaims["user_roles"].([]string)
			if !ok {
				commonerrors.HttpErrorWithMessage(w, err, http.StatusForbidden, "invalid token claim for value user_roles")
				return
			}

			data := token.TokenClaim{
				UserID:    userId,
				UserRoles: userRoles,
			}

			ctx := context.WithValue(r.Context(), localconstants.PROTECTED_DATA_KEY, data)

			t.handler.ServeHTTP(w, r.WithContext(ctx))
			return
		}
	}

	refreshCookie, err := r.Cookie(env.Get("REFRESH_HTTP_COOKIE_NAME"))
	if err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusUnauthorized, "failed to get refresh token cookie")
		return
	}
	if len(refreshCookie.Value) == 0 {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusUnauthorized, "failed to get refresh token cookie")
		return
	}

	refreshTokenStr := refreshCookie.Value

	refreshToken, err := token.Parse(refreshTokenStr, env.Get("TOKEN_REFRESH_SECRET"))
	if err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "error parsing refresh token")
		return
	}

	if !refreshToken.Valid {
		commonerrors.HttpError(w, http.StatusForbidden)
		return
	}

	refreshClaims, err := token.ExtractClaims(refreshToken, []byte(env.Get("TOKEN_REFRESH_SECRET")))
	if err != nil {
		commonerrors.HttpError(w, http.StatusInternalServerError)
		return
	}

	var refreshClaimsParsed token.TokenClaim
	var refreshClaimsBytes []byte

	refreshClaimsBytes, err = json.Marshal(refreshClaims)

	if err != nil {
		commonerrors.HttpError(w, http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(refreshClaimsBytes, &refreshClaimsParsed); err != nil {
		commonerrors.HttpError(w, http.StatusInternalServerError)
		return
	}

	if len(refreshClaimsParsed.UserID) == 0 || len(refreshClaimsParsed.UserRoles) == 0 {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusForbidden, "invalid refresh token claim for value user_id or user_roles are empty")
		return
	}

	expirationTime, err := refreshClaims.GetExpirationTime()
	if err != nil {
		commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "error getting expiration time")
		return
	}

	if time.Now().After(expirationTime.Time) {
		commonerrors.HttpError(w, http.StatusForbidden)
		return
	}

	if time.Now().Add(24 * time.Hour).After(expirationTime.Time) {
		newRefreshToken, err := token.CreateRefreshToken(refreshClaimsParsed.UserID, refreshClaimsParsed.UserRoles)
		if err != nil {
			commonerrors.HttpErrorWithMessage(w, err, http.StatusInternalServerError, "error creating new refresh token")
			return
		}

		refreshTokenName := env.Get("REFRESH_HTTP_COOKIE_NAME")

		newRefreshCookie := commonFunctions.CreateSecureCookie(refreshTokenName, newRefreshToken, refreshCookie.MaxAge, true)

		isAuthCookie := commonFunctions.CreateSecureCookie("has-refresh-token", "true", refreshCookie.MaxAge, false)

		http.SetCookie(w, &newRefreshCookie)
		http.SetCookie(w, &isAuthCookie)

	}

	newAccessToken, err := token.CreateAccessToken(refreshClaimsParsed.UserID, refreshClaimsParsed.UserRoles)
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
		commonerrors.EncodingError(w, err)
	}
}

func NewTokenRefresher(handlerToWrap http.Handler) *TokenRefresher {
	return &TokenRefresher{handler: handlerToWrap}
}
